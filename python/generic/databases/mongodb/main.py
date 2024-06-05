import pymongo


# client = pymongo.MongoClient("mongodb://localhost:27017")
client = pymongo.MongoClient(host="localhost", port=27017)


# 指定数据库
db = client.school  # client["school"]

# 指定集合
collection = db["students"]

my_data = {
        "name": "wangyao",
        "age": 18,
        "fullTime": False,
        "courses": ["Biology", "Calculus", "psychics"],
        "gender": "male"
}

result = collection.insert_one(my_data)

print(result)

# 获取对象的_id值
print(result.inserted_id)
