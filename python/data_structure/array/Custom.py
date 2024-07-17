

class Array:
    def __init__(self, capacity=10):
        self._data = []
        self._size = 0
        self._capacity = capacity

    def get_size(self):
        return self._size

    def get_capacity(self):
        return self._capacity

    def is_empty(self):
        return self._size == 0

    def add_last(self, e):
        self.add(self._size, e)

    def add_first(self, e):
        self.add(0, e)

    def add(self, index, e):
        if len(self._data) >= self._capacity:
            raise Exception("Array is already fulled.")
        if index < 0 and index > self._size:
            raise Exception("index error")
        for i in reversed(range(index, self._size + 1)):
            self._data[i + 1] = self._data[i]
        self._data[index] = e
        self._size += 1

