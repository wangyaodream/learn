from typing import Any


class Node:
    def __init__(self, data):
        self._data = data
        self.left: Node|None = None
        self.right: Node|None = None


class BST:
    def __init__(self):
        self._root: Node|None = None
        self._size = 0

    def get_size(self):
        return self._size

    def is_empty(self):
        return self._size == 0

    def add(self, data):
        if self._root is None:
            self.root = Node(data)
        else:
            self._add(self._root, data)

    def _add(self, node: Node, data):
        if data < node._data:
            if node.left is None:
                node.left = Node(data)

            else:
                self._add(node.left, data)
        else:
            if node.right is None:
                node.right = Node(data)
            else:
                self._add(node.right, data)
        
    def search(self, data: Any):
        return self._search(self._root, data)

    def _search(self, node: Node|None, data: Any):
        if node is None or node._data == data:
            return node
        if data < node._data:
            return self._search(node.left, data)
        return self._search
