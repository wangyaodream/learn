from bson.objectid import ObjectId
import pymongo


client = pymongo.MongoClient("mongodb://localhost:27017")

db = client["school"]

c = db["students"]

result = c.find_one({"name": "Alice"})


# 通过_id来进行查询
result = c.find_one({'_id': ObjectId('666063676a9416cfc7a26a13')})

# 通过正则查询
results = c.find({'name': {'$regex': '^B.*'}})

# 查询属性是否存在
results = c.find({'courses': {'$exists': True}})  # courses字段存在的


for i in results:
    print(i)
