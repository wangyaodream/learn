import os
import asyncpg
from aiohttp import web
from aiohttp.web_app import Application
from aiohttp.web_request import Request
from aiohttp.web_response import Response
from asyncpg import Record
from asyncpg.pool import Pool
from typing import List, Dict


routes = web.RouteTableDef()
DB_KEY = 'database'


async def create_database_pool(app: Application):
    """
    创建数据库连接池的协程
    """
    print('Creating database pool.')
    pool: Pool = await asyncpg.create_pool(host=os.getenv("PG_HOST"),
                                           port=5432,
                                           user=os.getenv("PG_USER"),
                                           password=os.getenv("PG_PASSWORD"),
                                           database='products',
                                           max_size=6,
                                           min_size=6)
    app[DB_KEY] = pool


async def destory_database_pool(app: Application):
    """
    销毁数据库连接
    """
    print('Destorying database pool.')
    pool: Pool = app[DB_KEY]
    await pool.close()


@routes.get('/brands')
async def brands(request: Request) -> Response:
    connection: Pool = request.app[DB_KEY]
    brand_query = "SELECT brand_id, brand_name FROM brand"
    results: List[Record] = await connection.fetch(brand_query)
    result_as_dict: List[Dict] = [dict(brand) for brand in results]
    return web.json_response(result_as_dict)


@routes.get('/products/{id}')
async def get_product(request: Request) -> Response:
    try:
        # 从url中获取id数据
        str_id = request.match_info['id']
        product_id = int(str_id)

        query = \
        """
        SELECT product_id, product_name, brand_id
        FROM product
        WHERE product_id = $1;
        """

        connection: Pool = request.app[DB_KEY]
        # 对单个产品进行查询
        result: Record = await connection.fetchrow(query, product_id)

        if result is not None:
            return web.json_response(dict(result))
        else:
            return web.HTTPNotFound()
    except ValueError:
        return web.HTTPBadRequest()



app = web.Application()
app.on_startup.append(create_database_pool)
app.on_cleanup.append(destory_database_pool)
app.add_routes(routes)
web.run_app(app, port=8000)
