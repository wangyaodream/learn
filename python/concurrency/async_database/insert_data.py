import asyncio
import asyncpg
from asyncpg import Record
from typing import List


async def main():
    connection = await asyncpg.connect(host="blog.trojancow.top",
                                       port=5432,
                                       user='postgres',
                                       database='products',
                                       password='Dream462213925')
    # execute方法是一个协程，所以需要await
    await connection.execute("INSERT INTO brand VALUES (DEFAULT, 'Levis')")
    await connection.execute("INSERT INTO brand VALUES (DEFAULT, 'Seven')")

    brand_query = "SELECT brand_id, brand_name FROM brand"
    results: List[Record] = await connection.fetch(brand_query)
    
    for brand in results:
        print(f'id: {brand["brand_id"]}, name: {brand["brand_name"]}')

    await connection.close()


asyncio.run(main())
