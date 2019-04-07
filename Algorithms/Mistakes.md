# Mistakes

This file is to keep track of all mistakes made in my code/problem solving.

## Trees
>    def inorder(self, root):
        if root is not None:
            for val in self.inorder(root.left):
                yield val
            yield root.val
            for val in self.inorder(root.right):
                yield val


It is often a good idea to use a generator as a generator will yield a more elegant solution. We can use iteration instead of recursion. 
(Problem: Range sum of BST)

# Graph
Think first about the right representation. 
If dfs traversal is needed, NEVER use matrix representation unless its searching 1 row/col away

In matrix dfs, REMEMBER to check bounds

# LinkedList

Reversing singly linked list


# General
1. Off by 1
