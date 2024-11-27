from contextlib import asynccontextmanager
from datetime import datetime
import os
import sys

from bson import ObjectId
from fastapi import FastAPI, status
from motor.motor_asyncio import AsyncIOMotorClient
from pydantic import Base64Str, BaseModel
import uvicorn

from dal import ToDoDAL, ListSummary, ToDoList

COLLECTION_NAME = "todo_lists"
MONGO_URI = os.environ["MONGO_URI"]
DEBUG = os.environ.get("DEBUG", "").strip().lower() in {"1",  "true", "on", "yes"}


@asynccontextmanager
async def lifespan(app: FastAPI):
    # startup
    client = AsyncIOMotorClient(MONGO_URI)
    db = client.get_default_database()

    pong = await db.command("ping")
    if int(pong["ok"]) != 1:
        raise Exception("Could not connect to MongoDB")

    todo_list = db.get_collection(COLLECTION_NAME)
    app.todo_dal = ToDoDAL(todo_list)

    yield

    client.close()

app = FastAPI(lifespan=lifespan, debug=DEBUG)

@app.get("/api/lists")
async def get_all_lists() -> list[ListSummary]:
    return [i async for i in app.todo_dal.list_todo_lists()]



class NewList(BaseModel):
    name: str


class NewListResponse(BaseModel):
    id: str
    name: str 


@app.post("/api/lists", status_code=status.HTTP_201_CREATED)
async def create_todo_list(new_list: NewList) -> NewListResponse:
    return NewListResponse(
            id=await app.todo_dal.create_todo_list(new_list.name),
            name=new_list.name,
    )

@app.get("/api/lists/{list_id}")
async def get_list(list_id: str) -> ToDoList:
    return await app.todo_dal.get_todo_list(ObjectId(list_id))


@app.delete("/api/lists/{list_id}")
async def delete_list(list_id: str):
    await app.todo_dal.delete_todo_list(ObjectId(list_id))

class NewItem(BaseModel):
    label: str


class NewItemResponse(BaseModel):
    id: str
    label: str


@app.post("/api/lists/{list_id}", status_code=status.HTTP_201_CREATED)
async def create_item(list_id: str, new_item: NewItem) -> ToDoList:
    return await app.todo_dal.create_item(list_id, new_item.label)

@app.delete("/api/lists/{list_id}/items/{item_id}")
async def delete_item(list_id: str, item_id: str) -> ToDoList:
    return await app.todo_dal.delete_item(list_id, item_id)


class ToDoItemUpdate(BaseModel):
    item_id: str
    checked_state: bool

@app.patch("/api/lists/{list_id}/checked_state")
async def set_checked_state(list_id: str, update: ToDoItemUpdate) -> ToDoList:
    return await app.todo_dal.set_checked_state(
            list_id, update.item_id, update.checked_state
    )


class DummyResponse(BaseModel):
    id: str
    when: datetime

@app.get("/api/dummy")
async def get_dummy() -> DummyResponse:
    return DummyResponse(id=str(ObjectId()), when=datetime.now())


def main(argv=sys.argv[1:]):
    try:
        uvicorn.run(app, host="0.0.0.0", port=3001, reload=DEBUG)
    except KeyboardInterrupt:
        pass


if __name__ == "__main__":
    main()

