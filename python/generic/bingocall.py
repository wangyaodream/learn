import random

class BingoCage:
    def __init__(self, items):
        self._items = items
        random.shuffle(self._items)

    def pick(self):
        try:
            return self._items.pop()
        except IndexError:
            raise LookupError('pick from empty BingoCage')

    def __call__(self):
        """
        实现__call__方法，让BingoCage实例化的object可以直接使用()的方式来调用
        """
        return self.pick()


if __name__ == "__main__":
    bingo = BingoCage([1,2,3])
    res = bingo()
    print(res)
