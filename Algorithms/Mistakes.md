# Mistakes

This file is to keep track of all mistakes made in my code/problem solving.

## Strings
Remember that some strings have spaces in them. Upper case, lower case.

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

1. Use the parameter to store the previous value, and assign
2. Use it recursively.

When implementing skipping, must use the .next instead of the node itself.

# General
1. Off by 1
2. If there is a valid/invalid state, ask what happens if starting value is valid, invalid
3. Boolean optimization:
```
if not reg and not string:
  return True
if not reg
  return False


if not reg:
  return not String
```
