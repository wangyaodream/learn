from fastapi import FastAPI, Request, Query, Body, status
from fastapi.encoders import jsonable_encoder
from pydantic import BaseModel
from starlette.middleware.base import BaseHTTPMiddleware


app = FastAPI()

class CustomMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        print("@@@ 请求参数 @@@", request.query_params)
        print("@@@ 请求方法 @@@", request.method)
        response = await call_next(request)
        return response



@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    print(request.query_params)
    print(request.method)

    response = await call_next(request)


    response.headers["X-Process-Token"] = str("test_token_polo")
    response.status_code = status.HTTP_202_ACCEPTED

    return response

@app.middleware("http")
async def log_request(request: Request, call_next):
    print("@@@ 请求参数 @@@", request.query_params)
    print("@@@ 请求方法 @@@", request.method)

    response = await call_next(request)

    return response

@app.get("/")
async def root():
    return {"message": "Hello World"}

class User(BaseModel):
    name: str = None
    age: int = None


@app.post("items/")
async def read_item(item_id: str = Query(...), user: User = Body(...)):
    res = {'item_id': item_id}
    if user:
        res.update(jsonable_encoder(user))

    print("@@@ 执行操作 @@@", res)

    return res


app.add_middleware(CustomMiddleware)