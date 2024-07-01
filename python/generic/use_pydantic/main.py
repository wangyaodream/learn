from pydantic import BaseModel


class Address(BaseModel):
    street: str
    city: str
    zipcode: str


class User(BaseModel):
    name: str
    age: int
    address: Address



user = User(
    name="wangyao",
    age=18,
    address={
        "street": "Center road",
        "city": "Chongqing",
        "zipcode": "CQ"
    }
)

print(user)