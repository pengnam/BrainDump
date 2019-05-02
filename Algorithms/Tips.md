# Tips
## General
Ensuring that two values are 1 away from each other
abs(m-n)

if list is used at any time, always ask what happens if list is empty

Steps:
1. Give a higher level idea of the solution
2. Give the expected time complexity
## Arrays and Strings

Algorithms to consider:
KMP

ASCII character set is 128 characters

If you have to modify something inplace, try doing it from the back
Lists also support slicing assignment.

A[1:3] = "Ab"


!! When handling strings, remember that there exist upper case and lower case, and sometimes they are case insensitive

Nifty trick:
return min(string, ''.join(compressed), key=len)

Check if 1 substring is a subset of the other substring

## LinkedList
Always check if you are removing from the start of the linkedlist (requires reassignment of head ptr)
Consider padding linkedlist with nodes infront (consider ctci sum lists)


## Backtracking
Be clear that its adding the piece at this position, checking if it works, and continue (in-place, so your helpers should not return the board but continue modifying the same board)

!! Backtracking should be done in-place

## Binary Search
Binary search appears in funny places. If there is a bounds on the domains to search, try binary searching the value and applying some form of verifying algorithm


## Graph 
If need to do BFS, use a deque()
Ask if there are duplicates (!) esp. in bst
