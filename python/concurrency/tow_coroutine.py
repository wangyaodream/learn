import asyncio
from util import delay


async def hello_every_second():
    for _ in range(2):
        await asyncio.sleep(1)
        print("I'm running other code whule I'm waiting")


async def main():
    first_delay = asyncio.create_task(delay(3))
    second_delay = asyncio.create_task(delay(3))
    await hello_every_second()
    await first_delay
    await second_delay


asyncio.run(main())
    
