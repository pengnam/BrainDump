# Dynamic Programming

Note that there are two forms of dynamic programming. The first is a bottom-up programming, while the second is top-down programming (often known as memoization).

Try writing a recursive version, and then memoize it. Alternatively, write a
top-down version of it

## Common Problems
1. Longest Common Subsequence (Not contiguous)

If the two last characters match, try 1 + LCS(m-1, n-1)
If they do not match, try max(LCS(m-1,n), LCS(m, n-1))
If m == 0 or n == 0, return 0

O(mn)

2. Longest Palindromic Subsequence
For a given start and ending vertex i,j

if A[i] == A[j], try 1 + LPS(i-1, j-1), and use value if LPS returns True.
also, try LPS(i-1, j) and LPS(i, j-1).

O(n^2)
There is also an O(1) space algorithm.

Note:
- There also exist an O(n) solution to palindromic subsequence, Manchester algorithm.
- There is an interesting solution that uses LCS:
https://www.techiedelight.com/implement-diff-utility/

3. Longest Increasing Subsequence (Not continuous)
Each pass, find the all elements less than it, find the memoized value for each of them, take the maximum.

O(n^2)

4. 0-1 Knapsack

Time: O(nw)

Memo table is n (number of items) * W (maximum weight)


I believe there should be an O(W) table solution.

5. Subset Sum Problem

Memo table is O(n * sum)

Same as knapsack

6. Partition Problem

Same as subset sum problem

7. K partition

Same as subset sum, but now form a k dimensional memo

8. Rod cutting problem

For rod of length i,
Try:
price[j] + rodCut(i-j)
