import pi
import time


iter_round = 100_000_000

def single_proc() -> float:
    return pi.calculate_pi(0, iter_round)


def main():
    start_time = time.time()
    print(single_proc())
    print(f"Single_process:{time.time() - start_time} second(s)")


if __name__ == "__main__":
    main()

