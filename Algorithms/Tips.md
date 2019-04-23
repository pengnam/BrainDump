# Tips
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


## Backtracking
Be clear that its adding the piece at this position, checking if it works, and continue (in-place, so your helpers should not return the board but continue modifying the same board)

!! Backtracking should be done in-place
