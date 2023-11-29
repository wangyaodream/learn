import random
import asyncio
from websockets.server import serve


NAMES = ["wangyao", "zhangman", "dream"]

async def echo(websocket):
    async for message in websocket:
        await websocket.send(message + " " + random.choice(NAMES))


async def main():
    async with serve(echo, "localhost", 8765):
        await asyncio.Future()


asyncio.run(main())
