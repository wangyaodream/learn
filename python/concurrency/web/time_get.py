from aiohttp import web
from datetime import datetime
from aiohttp.web_request import Request
from aiohttp.web_response import Response


routes = web.RouteTableDef()

@routes.get('/time')
async def time(request: Request) -> Response:
    today = datetime.today()

    result = {
        "month": today.month,
        "day": today.day,
        "time": f"{today.hour}:{today.minute}:{today.second}"
    }

    return web.json_response(result)


@routes.get('/')
async def get_data(request: Request) -> Response:
    share_data = request.app['share_data']
    return web.Response(share_data)


app = web.Application()
app.add_routes(routes)
web.run_app(app, port=8888)