from fastapi import FastAPI
from dotenv import dotenv_values
from pymongo import MongoClient


config = dotenv_values(".env")

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Welcome to the PyMongo tutorial!"}


@app.on_event("startup")
def startup_db_client():
    app.mongodb_client = MongoClient(config["MONGO_URI"])
    app.database = app.mongodb_client[config["DB_NAME"]]

@app.on_event("shutdown")
def shutdown_db_client():
    app.mongodb_client.close()
