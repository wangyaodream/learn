class NodeIter:
    def __init__(self, node):
        self.curr_node = node

    def __next__(self):
        if self.curr_node is None:
            raise StopIteration
        node, self.curr_node = self.curr_node, self.curr_node.next
        return node


class Node:
    def __init__(self, name):
        self.name = name
        self.next = None

    def __iter__(self):
        return NodeIter(self)


node1 = Node("node1")
node2 = Node("node2")
node3 = Node("node3")

node1.next = node2
node2.next = node3


