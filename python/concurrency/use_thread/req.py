import sys
sys.path.append("..")
import asyncio
import requests
import functools
from concurrent.futures import ThreadPoolExecutor


from util import async_timed


def get_status_code(url: str) -> int:
    response = requests.get(url)
    #
     
    return response.status_code


# start = time.time()

# with ThreadPoolExecutor() as pool:
#     urls = ['https://www.example.com' for _ in range(1000)]
#     results = pool.map(get_status_code, urls)
#     for result in results:
#         print(result)

# end = time.time()

# print(f'finished requests in {end - start:.4f} second(s)')
@async_timed()
async def main():
    loop = asyncio.get_running_loop()
    with ThreadPoolExecutor() as pool:
        urls = ['https://www.example.com' for _ in range(1000)]
        tasks = [loop.run_in_executor(pool, functools.partial(get_status_code, url)) for url in urls]
        results = await asyncio.gather(*tasks)
        print(results)

asyncio.run(main())

    