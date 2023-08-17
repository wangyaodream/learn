import sys
sys.path.append("..")
from util import async_timed
import asyncio
import asyncpg


product_query = \
        """select p.product_id, p.product_name,p.brand_id,s.sku_id,pc.product_color_name,ps.product_size_name 
from product as p
join sku as s on s.product_id  = p.product_id 
join product_color as pc on pc.product_color_id = s.product_color_id
join product_size as ps on ps.product_size_id  = s.product_size_id
where p.product_id = 100;
"""

async def query_product(pool):
    async with pool.acquire() as connection:
        return await connection.fetchrow(product_query)


@async_timed()
async def query_product_synchronously(pool, queries):
    return [await query_product(pool) for _ in range(queries)]


@async_timed()
async def query_product_concurrently(pool, queries):
    queries = [query_product(pool) for _ in range(queries)]
    return await asyncio.gather(*queries)

async def main():
    async with asyncpg.create_pool(host="blog.trojancow.top",
                                       port=5432,
                                       user="postgres",
                                       database="products",
                                       password="Dream462213925",
                                       min_size=6,
                                   max_size=6) as pool:
        await query_product_synchronously(pool, 1000)
        await query_product_concurrently(pool, 1000)


asyncio.run(main())


