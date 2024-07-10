import asyncio
import random
from time import perf_counter


async def process(num: int):
    await asyncio.sleep(1)
    return f"process-{num}"


async def next_process(total: int):
    for _ in range(total):
        res = await process(random.randint(1, 999))
        yield res


async def main():
    time_before = perf_counter()
    # for _ in range(4):
    #     r = await process(random.randint(1, 999))
    #     print(r)

    async for i in next_process(4):
        print(i)
    print(f"Total time (synchronous): {perf_counter() - time_before}")

    
    time_before = perf_counter()
    # 作为对比，使用gather后只需要3秒
    result = await asyncio.gather(*[process(random.randint(1, 999)) for _ in range(4)])
    print(result)
    print(f"Total time (asynchronous): {perf_counter() - time_before}")


if __name__ == "__main__":
    asyncio.run(main())

