import time
from multiprocessing import Pool


def say_hello(name: str, delay) -> str:
    time.sleep(delay)
    return f'Hello {name}'


if __name__ == "__main__":
    start_time = time.time()
    with Pool() as process_pool:
        hi_jeff = process_pool.apply_async(say_hello, args=('jeff', 2))
        hi_john = process_pool.apply_async(say_hello, args=('john', 3))
        print(hi_jeff.get())
        print(hi_john.get())
    print(f"Use time {time.time() - start_time}")
