import asyncio
from asyncio import StreamReader
import sys
from .sequence import move_back_one_char, clear_line
from collections import deque
from typing import Callable, Deque, Awaitable


async def create_stdin_reader() -> StreamReader:
    stream_reader = asyncio.StreamReader()
    procotol = asyncio.StreamReaderProtocol(stream_reader)
    loop = asyncio.get_running_loop()
    await loop.connect_read_pipe(lambda: procotol, sys.stdin)
    return stream_reader


async def read_line(stdin_reader: StreamReader) -> str:
    def erase_last_char():
        move_back_one_char()
        sys.stdout.write(' ')
        move_back_one_char()


    delete_char = b'\x7f'
    input_buffer = deque()
    while (input_char := await stdin_reader.read(1)) != b'\n':
        if input_char == delete_char:
            if len(input_buffer) > 0:
                input_buffer.pop()
                erase_last_char()
                sys.stdout.flush()
        else:
            input_buffer.append(input_char)
            sys.stdout.write(input_char.decode())
            sys.stdout.flush()
    clear_line()
    return b''.join(input_buffer).decode()

class MessageStore:
    def __init__(self, callback: Callable[[Deque], Awaitable[None]], max_size: int):
        self._deque = deque(maxlen=max_size)
        self._callback = callback
    
    async def append(self, item):
        self._deque.append(item)
        await self._callback(self._deque)