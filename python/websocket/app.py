import asyncio

from websockets.server import serve


async def handler(websocket):
    while True:
        message = await websocket.recv()
        print(message)


async def main():
    async with serve(handler, "", 8001):
        await asyncio.Future()


if __name__ == "__main__":
    asyncio.run(main())
