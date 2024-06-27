import functools
import logging
from typing import Any, Callable
from math import sqrt
from time import perf_counter


logger = logging.getLogger("my-app")

def is_prime(number: int) -> bool:
    if number < 2:
        return False

    for element in range(2, int(sqrt(number)) + 1):
        if number % element == 0:
            return False
    return True


# def bechmark(upper_bound: int) -> int:
#     start_time = perf_counter()
#     value = count_primer_numbers(upper_bound)
#     end_time = perf_counter()
#     run_time = end_time - start_time
#     logging.info(
#             f"Execution of count_primer_numbers took {run_time:.2f} seconds."
#     )
#     return value

def bechmark(func: Callable[..., Any]) -> Callable[..., Any]:
    @functools.wraps(func)
    def wrapper(*args: Any, **kwargs: Any) -> Any:
        start_time = perf_counter()
        value = func(*args, **kwargs)
        end_time = perf_counter()
        run_time = end_time - start_time
        logging.info(
                f"Execution of {func.__name__} took {run_time:.2f} seconds."
        )
        return value
    return wrapper

# 无参数
# def with_logging(func: Callable[..., Any]) -> Callable[..., Any]:
#     @functools.wraps(func)
#     def wrapper(*args: Any, **kwargs: Any) -> Any:
#         logging.info(f"Calling {func.__name__}")
#         value = func(*args, **kwargs)
#         logging.info(f"Finished calling {func.__name__}")
#         return value
#     return wrapper
def with_logging(logger: logging.Logger):
    def decorator(func: Callable[..., Any]) -> Callable[..., Any]:
        @functools.wraps(func)
        def wrapper(*args: Any, **kwargs: Any) -> Any:
            logger.info(f"Calling {func.__name__}")
            value = func(*args, **kwargs)
            logger.info(f"Finished calling {func.__name__}")
            return value
        return wrapper
    return decorator

with_default_logging = functools.partial(with_logging, logger=logger)

@with_default_logging()
@bechmark
def count_primer_numbers(upper_bound: int) -> int:
    count = 0
    for number in range(upper_bound):
        if is_prime(number):
            count += 1
    return count


def main() -> None:
    logging.basicConfig(level=logging.INFO)
    # wrapper = bechmark(count_primer_numbers)
    # value = wrapper(100000)
    value = count_primer_numbers(100000)
    logging.info(f"Number of primes: {value}")


if __name__ == "__main__":
    main()

