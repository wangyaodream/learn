import asyncio
import logging
import asyncpg



async def main():
    connection = await asyncpg.connect(host="blog.trojancow.top",
                                       port=5432,
                                       user='postgres',
                                       database='products',
                                       password='Dream462213925')
    try:
        async with connection.transaction():
            insert_brand = 'INSERT INTO brand VALUES (9999, "big_brand")'
            await  connection.execute(insert_brand)
            await  connection.execute(insert_brand)
    except Exception:
        logging.exception('Error running tttttttt')
    finally:
        query = "SELECT brand_name from brand where brand_name like 'big_%'"
        brands = await connection.fetch(query)

        print(f'Query result was: {brands}')

        await connection.close()

asyncio.run(main())
