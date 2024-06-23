import json

from fastapi import FastAPI
from dotenv import dotenv_values
from pymongo import MongoClient


config = dotenv_values(".env")

app = FastAPI()

client = MongoClient(config["MONGO_URI"])
db = client["school"]
collection = db["student"]


@app.get("/")
def index():
    return {"message": "ok"}


@app.get("/student/{student_id}")
def detail(student_id: str):
    info = collection.find_one({'student_id': student_id})
    if info is None:
        return {"message": "empty"}
    else:
        info["_id"] = str(info["_id"])
        return info


@app.post('/student/')
def create_student(info: dict):
    result = collection.insert_one(info)
    return {'_id': str(result.inserted_id)}
