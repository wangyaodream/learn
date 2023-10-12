import asyncpg
from aiohttp.web_app import Application
from asyncpg.pool import Pool


DB_KEY = 'database'

async def create_database_pool(app: Application,
                               host: str,
                               port: int,
                               user: str,
                               password: str,
                               database: str):
    print("creating database pool.")
    pool: Pool = await asyncpg.create_pool(host=host,
                                     port=port,
                                     user=user,
                                     password=password,
                                     database=database,
                                     max_size=6,
                                     min_size=6)
    app[DB_KEY] = pool


async def destory_database_pool(app: Application):
    pool = app[DB_KEY]
    await pool.close()

