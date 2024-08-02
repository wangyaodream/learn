import time


def timer(func):
    def wrapper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        end_time = time.time()
        execution_time = end_time - start_time
        print(f"Execution time: {execution_time} second(s)")
        return result
    return wrapper


def debugger(func):
    def wrapper(*args, **kwargs):
        print(f"Calling {func.__name__} with args: {args} kwargs: {kwargs}")
        result = func(*args, **kwargs)
        print(f"{func.__name__} returned: {result}")
        return result
    return wrapper


@debugger
def add_number(x, y):
    return x + y


@timer
def train_model():
    # Do somthing
    time.sleep(3)
    print("Done!")


if __name__ == "__main__":
    # train_model()
    add_number(10, y=29)
