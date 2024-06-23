import pymongo

client = pymongo.MongoClient("mongodb://localhost:27017")

db = client["school"]

collection = db["student"]


result = collection.find_one({"student_id": "1002"})

print(result)
