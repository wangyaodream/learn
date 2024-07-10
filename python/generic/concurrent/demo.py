import asyncio
import random


async def process(num: int):
    await asyncio.sleep(3)
    return f"process-{num}"


async def main():
    for _ in range(4):
        r = await process(random.randint(1, 999))
        print(r)

    # 作为对比，使用gather后只需要3秒
    result = await asyncio.gather(*[process(random.randint(1, 999)) for _ in range(4)])
    print(result)


if __name__ == "__main__":
    asyncio.run(main())

