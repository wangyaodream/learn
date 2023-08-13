import sys
sys.path.append("..")
import asyncio
import aiohttp
from aiohttp_req import fetch_status


async def main():
    async with aiohttp.ClientSession() as session:
        api_a = asyncio.create_task(fetch_status(session, "https://www.baidu.com"))
        api_b = asyncio.create_task(fetch_status(session, "https://www.baidu.com", delay=5))

        done, pending = await asyncio.wait([api_a, api_b], timeout=1)

        for task in pending:
            if task is api_b:
                print('API B is too slow, canceling')
                task.cancel()

asyncio.run(main())
