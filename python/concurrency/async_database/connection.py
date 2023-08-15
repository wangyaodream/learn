import asyncpg
import asyncio


async def main():
    connection = await asyncpg.connect(host="blog.trojancow.top",
                                       port=5432,
                                       user='postgres',
                                       database='postgres',
                                       password='Dream462213925')
    version = connection.get_server_version()
    print(f'Connected! Postgres version {version}')
    await connection.close()


asyncio.run(main())
