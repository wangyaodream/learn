import random
from functools import wraps
from time import time


class Record():
    """ 定义类的方式来定义装饰器 """

    def __init__(self, output):
        self.output = output

    def __call__(self, func):
        
        @wraps(func)
        def wrapper(*args, **kwargs):
            start = time()
            result = func(*args, **kwargs)
            self.output(func.__name__, time() - start)
            return result
        return wrapper

def record_time(func):
    """ 常规装饰器 """

    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time()
        result = func(*args, **kwargs)
        print(f'{func.__name__}: {time() - start}')
        return result
    return wrapper


def record(output):
    """ 可参数化的装饰器 """

    def decorate(func):

        @wraps(func)
        def wrapper(*args, **kwargs):
            start = time()
            result = func(*args, **kwargs)
            output(func.__name__, time() - start)
            return result
        return wrapper
    return decorate

def bar(name, time):
    print(f'{name}: {time} bzw!')


@record(bar)
def foo(s):
    return [x * 2 for x in s]

sample = [random.randint(1, 999) for _ in range(1000000)]

foo(sample)



