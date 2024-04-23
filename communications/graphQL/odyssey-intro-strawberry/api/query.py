import strawberry


# hello resolver
def get_hello():
    return "hello world"

@strawberry.type
class Query:
    hello: str = strawberry.field(resolver=get_hello)



