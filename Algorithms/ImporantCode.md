
def iterator(node):
    if node is None:
        return
    yield from iterator(node.left)
    yield node.val
    yield from iterator(node.right)
