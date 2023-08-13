import sys
sys.path.append("..")
import aiohttp
import asyncio
from aiohttp import ClientSession
from util import async_timed
from aiohttp_req import fetch_status


@async_timed()
async def main():
    async with aiohttp.ClientSession() as session:
        fetchers = [asyncio.create_task(fetch_status(session, "https://www.baidu.com")),
                    asyncio.create_task(fetch_status(session, "https://www.baidu.com"))]
        done, pending = await asyncio.wait(fetchers)

        print(f"Done task count: {len(done)}")
        print(f'Pending task count: {len(pending)}')


        for done_task in done:
            result = await done_task
            print(result)


asyncio.run(main())
