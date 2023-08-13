import sys
sys.path.append("..")
import asyncio
import aiohttp
from aiohttp import ClientSession
from aiohttp_req import fetch_status
from util import async_timed


@async_timed()
async def main():
    async with aiohttp.ClientSession() as session:
        urls = ['https://www.baidu.com' for _ in range(100)]
        requests = [fetch_status(session, url) for url in urls]
        status_code = await asyncio.gather(*requests)
        print(status_code)


asyncio.run(main())
