import asyncio


async def worker(semaphore, task_id):
    async with semaphore:
        print(f"Task {task_id} started")
        await asyncio.sleep(2)
        print(f"Task {task_id} finished")


async def main():
    semaphore = asyncio.Semaphore(2)
    tasks = [worker(semaphore, i) for i in range(5)]
    await asyncio.gather(*tasks)


asyncio.run(main())
