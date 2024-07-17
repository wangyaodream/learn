
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class LinkedList:
    def __init__(self):
        self.head = None

    def append(self, data):
        new_node = Node(data)
        if not self.head:
            # 如果是空表示没有任何node
            self.head = new_node
        last_node = self.head
        while last_node.next:
            last_node = last_node.next
        if last_node.next:
            last_node.next = new_node


    def prepend(self, data):
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node


    def insert_after_node(self, prev_node, data):
        if prev_node is None:
            print("Previous node must not be null")
            return
        new_node = Node(data)
        new_node.next = prev_node.next
        prev_node.next = new_node

