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


def memoize(func):
    cache = {}
    def wrapper(*args):
        if args in cache:
            return cache[args]
        else:
            result = func(*args)
            cache[args] = result
            return result
    return wrapper


def retry(max_attepts, delay=1):
    def decorator(func):
        def wrapper(*args, **kwargs):
            attempts = 0
            while attempts < max_attepts:
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    attempts += 1
                    print(f"Attempt {attempts} failed: {e}")
                    time.sleep(delay)
            print(f"Function failed after {max_attepts} attempts")
        return wrapper
    return decorator


@retry(max_attepts=3, delay=2)
def fetch_data(url):
    print("Fetching the data...")
    raise TimeoutError("Server is not responding")


@memoize
def fibonacci(n):
    if n <= 1:
        return n
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)


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
    # add_number(10, y=29)
    # r = fibonacci(19)
    # print(r)
    fetch_data("https://www.google.com")

