import asyncio

from util import delay


async def main():
    delay_task = asyncio.create_task(delay(10))
    try:
        result = await asyncio.wait_for(delay_task, timeout=3)
        print(result)
    except asyncio.exceptions.TimeoutError:
        print('Got a timeout')
        print(f"Was the task cancelled? {delay_task.cancelled}")


async def main2():
    task = asyncio.create_task(delay(10))

    # try:
    #     result = await asyncio.wait_for(asyncio.shield(task), 5)
    #     print(result)
    # except TimeoutError:
    #     print('Task took longer tnan five seconds, it will finish soon!')
    #     result = await task
    #     print(result)
    result = await asyncio.wait_for(asyncio.shield(task), 5)
    print(result)



asyncio.run(main2())
