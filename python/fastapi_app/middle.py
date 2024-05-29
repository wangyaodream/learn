from fastapi import FastAPI, Request, Query, Body, status
from fastapi.encoders import jsonable_encoder
from pydantic import BaseModel


app = FastAPI()


@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    print(request.query_params)
    print(request.method)

    response = await call_next(request)


    response.headers["X-Process-Token"] = str("test_token_polo")
    response.status_code = status.HTTP_202_ACCEPTED

    return response


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
