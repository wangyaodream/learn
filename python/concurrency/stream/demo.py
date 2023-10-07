import asyncio
from ..util import delay

async def main():
    while True:
        delay_time = input('Enter a time to sleep:')
        asyncio.create_task(delay(int(delay_time)))


if __name__ == "__main__":
    asyncio.run(main())

