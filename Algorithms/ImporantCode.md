#Iterator for tree
def iterator(node):
    if node is None:
        return
    yield from iterator(node.left)
    yield node.val
    yield from iterator(node.right)


#Prefix sum
P = [0]
for x in A:
    P.append(P[-1] + x)


#Binary search (Duplicates)
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
