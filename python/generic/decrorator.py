import time

class Timer:
    def __init__(self, func):
        self.func = func

    def __call__(self, *args, **kwargs):
        start = time.time()
        ret = self.func(*args, **kwargs)
        print(f"Time: {time.time() - start}")
        return ret

@Timer
def add(a, b):
    return a + b

# 上面的定义方式等价于 add = Timer(add)
# add从一个函数变成了一个Timer类实例化后的对象
print(type(add))

print(add(2, 3))


class TimerPre:
    def __init__(self, prefix):
        self.prefix = prefix

    def __call__(self, func):
        def wrapper(*args, **kwargs):
            start = time.time()
            ret = func(*args, **kwargs)
            print(f"{self.prefix}: {time.time() - start}")
            return ret
        return wrapper

@TimerPre(prefix="curr")
def add2(a, b):
    return a + b

print(add2(10, 20))
