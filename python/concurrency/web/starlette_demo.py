import os
import asyncpg
from asyncpg import Record
from asyncpg.pool import Pool
from starlette.applications import Starlette
from starlette.requests import Request
from starlette.responses import Response, JSONResponse
from starlette.routing import Route
from typing import List, Dict


async def create_database_pool():
    pool: Pool = await asyncpg.create_pool(host="127.0.0.1",
                                     user="postgres",
                                     password="Dream462213925",
                                     port=5432,
                                     database="products",
                                     min_size=6,
                                     max_size=6)
    app.state.DB = pool


async def destroy_database_pool():
    pool = app.state.DB
    await pool.close()


async def brands(request: Request) -> Response:
    connection: Pool = request.app.state.DB
    brand_query = "SELECT brand_id, brand_name FROM brand"
    results: List[Record] = await connection.fetch(brand_query)
    result_as_dict: List[Dict] = [dict(row) for row in results]
    return JSONResponse(result_as_dict)


async def test(request: Request) -> Response:
    return JSONResponse({"Hello": "World"})


app = Starlette(routes=[Route('/brands', brands), Route("/test", test)],
                on_startup=[create_database_pool],
                on_shutdown=[destroy_database_pool])

