from threading import Lock, Thread
import time


# 创建一个死锁
lock_a = Lock()
lock_b = Lock()


def a():
    with lock_a:
        print('Acquired lock_a, trying to acquire lock_b')
        time.sleep(1)
        with lock_b:
            print('Acquired lock_b')
        


def b():
    with lock_b:
        print('Acquired lock_b, trying to acquire lock_a')
        with lock_a:
            print('Acquired lock_a')


thread_1 = Thread(target=a)
thread_2 = Thread(target=b)
thread_1.start()
thread_2.start()
thread_1.join()
thread_2.join()