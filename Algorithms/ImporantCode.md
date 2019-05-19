# Behaviour Questions
Why do you want to work at Google?
 - Google is a successful company. 
 - Learn what makes these companies successful (from a technical perspective) and how they work. 
     -If you don’t see early on what structure does look like - it will be hard to set goals later on.
 - Technically challenging, and I think it would be nice to understand how engineering teams work in a company as large as google. 

# Tell me about yourself
I like coding and thinking logically. I actually came from a research orientated background.
I like both the idea of pushing the boundaries of what I know, and I also like making things. 

I started to realise I like forming actual real ideas into code, more than I like the research ideas. 

In my spare time, I read books, exercise. 


# Iterator for tree
def iterator(node):
    if node is None:
        return
    yield from iterator(node.left)
    yield node.val
    yield from iterator(node.right)


# Prefix sum
P = [0]
for x in A:
    P.append(P[-1] + x)


# Binary search (Duplicates)
	low = 0
	high = len(nums)-1
	index = -1
	while low <= high:
			mid = (high + low) // 2
			if nums[mid] == target and (mid == 0 or nums[mid-1] != target):
					index = mid
					break
			elif target > nums[mid]:
					low = mid + 1
			else:
					high = mid - 1

# Merge sort
# Quick sort
