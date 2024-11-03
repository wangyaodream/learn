import queue

class Node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

def bfs(root):
    q = queue.Queue()
    q.put(root)
    while not q.empty():
        node = q.get()
        print(node.left)
        if node.left:
            q.put(node.left)
        if node.right:
            q.put(node.right)


root = Node(1)
root.left = Node(2)
root.right = Node(3)

root.left.left = Node(4)
root.right.right = Node(5)
bfs(root)
