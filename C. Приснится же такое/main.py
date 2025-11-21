from collections import deque


class Node:
    def __init__(self, value, parent=None, left=None, right=None) -> None:
        self.value = value
        self.left = left
        self.right = right
        self.parent = parent


def swap_with_left_child(p: Node, v: Node, node_map: dict[int, Node]):
    if v.right:
        v.right.parent = p
    if p.right:
        p.right.parent = v

    p.value, v.value = v.value, p.value
    p.right, v.right = v.right, p.right
    node_map[v.value] = v
    node_map[p.value] = p


def swap_with_right_child(p: Node, v: Node, node_map: dict[int, Node]):
    if v.left:
        v.left.parent = p
    if p.left:
        p.left.parent = v

    p.value, v.value = v.value, p.value
    p.left, v.left = v.left, p.left
    node_map[v.value] = v
    node_map[p.value] = p


def swap(p: Node, v: Node, node_map: dict[int, Node]):
    if p.left == v:
        swap_with_left_child(p, v, node_map)
    else:
        swap_with_right_child(p, v, node_map)


def build_tree(n: int) -> tuple[Node, dict[int, Node]]:
    root = Node(1)
    queue = deque()
    queue.append(root)
    current_node_id = 2
    node_map = {1: root}

    while queue and current_node_id <= n:
        parent = queue.popleft()

        if current_node_id <= n:
            parent.left = Node(current_node_id, parent)
            queue.append(parent.left)
            node_map[current_node_id] = parent.left
            current_node_id += 1
        if current_node_id <= n:
            parent.right = Node(current_node_id, parent)
            queue.append(parent.right)
            node_map[current_node_id] = parent.right
            current_node_id += 1

    return root, node_map


def lvr(root: Node):
    if not root:
        return

    stack = []
    curr = root

    while curr or stack:
        while curr:
            stack.append(curr)
            curr = curr.left

        curr = stack.pop()
        print(curr.value, end=" ")
        curr = curr.right


def read_words(f, buffer_size=4096):
    buffer = ''

    while True:
        chunk = f.read(buffer_size)

        if not chunk:
            if buffer:
                yield from buffer.split()
            break

        buffer += chunk
        parts = buffer.split()

        if not buffer or buffer[-1].isspace():
            yield from parts
            buffer = ''
        else:
            yield from parts[:-1]
            buffer = parts[-1] if parts else ''


def main(filename: str) -> Node:
    with open(filename) as f:
        N = int(f.readline().split()[0])
        root, node_map = build_tree(N)

        for word in read_words(f):
            node_id = int(word)
            node = node_map[node_id]

            if node == root:
                continue

            swap(node.parent, node, node_map)

    return root


if __name__ == "__main__":
    root = main("input.txt")
    lvr(root)
