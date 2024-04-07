from typing import Optional, Callable, Literal


# 形式等价与typing中的Union，Union[int, None]
def f(x: Optional[int]) -> int:
    if x is None:
        return 0
    return x


users: list[int] = []
# 语法检测会报错，但是解释器并不会报错
users.append("1")

# 此时Callable会规定func函数时一个有两个int参数和一个int返回值的函数
def my_dec(func: Callable[[int, int], int]):
    def wrapper(a: int, b: int) -> int:
        print(f'args = {a}, {b}')
        ret = func(a, b)
        print(f"result = {ret}")
        return ret

    return wrapper

class Person:
    def __init__(self, 
                 name: str, 
                 gender: Literal["male", "female"]
                 ):
        self.name = name
        self.gender = gender

a = Person("Bob", "woman")
b = Person("Bob", "male")

ReturnType = tuple[int, Optional[str]]

def f(a) -> ReturnType:
    if a > 0:
        print(a)
        return 0, None
    else:
        return 1, "input is <= 0"


