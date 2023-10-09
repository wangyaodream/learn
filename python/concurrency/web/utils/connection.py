import os
import asyncpg
from aiohttp.web_app import Application
from asyncpg.pool import Pool

DB_KEY = 'database'

async def create_database_pool(app: Application):
    print("Creating database pool.")
    pool: Pool = await asyncpg.create_pool(host=os.getenv("PG_HOST"),
                                           user=os.getenv("PG_USER"),
                                           password=os.getenv("PG_PASSWORD"),
                                           port=5432,
                                           database="products",
                                           max_size=6,
                                           min_size=6)
    app[DB_KEY] = pool


async def destory_database_pool(app: Application):
    print('Destorying database pool.')
    pool: Pool = app[DB_KEY]
    await pool.close()




