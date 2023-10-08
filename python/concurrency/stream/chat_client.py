import asyncio
import os
import logging
import tty
from asyncio import StreamReader, StreamWriter
from collections import deque
from util.common import create_stdin_reader, read_line, MessageStore
from util.sequence import *


async def send_message(message: str, writer: StreamWriter):
    writer.write(message.encode())
    await writer.drain()


async def listen_for_messages(reader: StreamReader, message_store: MessageStore):
    while (message := await reader.readline()) != b'':
        await message_store.append(message.decode())
    await message_store.append('Server closed connection.')


async def read_and_send(stdin_reader: StreamReader, writer: StreamWriter):
    while True:
        message = await read_line(stdin_reader)
        await send_message(message, writer)


async def main():
    async def redraw_output(items: deque):
        save_cursor_position()
        move_to_top_of_screen()
        for item in items:
            clear_line()
            sys.stdout.write(item)
        restore_cursor_position()
    
    tty.setcbreak(0)
    os.system("clear")
    rows = move_to_bottom_of_screen()

    messages = MessageStore(redraw_output, rows - 1)

    stdin_reader = await create_stdin_reader()
    sys.stdout.write('Enter username:')
    username = await read_line(stdin_reader)

    reader, writer = await asyncio.open_connection('127.0.0.1', 8888)

    writer.write(f'CONNECT {username}\n'.encode())
    await writer.drain()

    message_listener = asyncio.create_task(listen_for_messages(reader, messages))
    input_listener = asyncio.create_task(read_and_send(stdin_reader, writer))

    try:
        await asyncio.wait([message_listener, input_listener], return_when=asyncio.FIRST_COMPLETED)
    except Exception as e:
        logging.exception("Error in main", exc_info=e)
        writer.close()
        await writer.wait_closed()


asyncio.run(main())