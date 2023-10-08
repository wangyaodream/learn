import os
import asyncpg
import asyncio


async def main():
    connection = await asyncpg.connect(host=os.getenv("PG_HOST"),
                                       port=5432,
                                       user=os.getenv("PG_USER"),
                                       database='postgres',
                                       password=os.getenv("PG_PASSWORD"))
    version = connection.get_server_version()
    print(f'Connected! Postgres version {version}')
    await connection.close()


asyncio.run(main())
