k = 4
n = 4
def permutation(input_list, total, pre, iteration):

        if len(input_list) == total - k + 1:
                print("HERE")
                return [[x] for x in input_list]
        n = len(input_list)
        result = []
        for i in range(n):
                word = input_list[i]
                wrong = False
                for j in range(iteration):
                    if word[j] != pre[j][iteration]:
                        wrong = True
                        break
                if wrong:
                    continue

                remaining_words = input_list[:i] + input_list[i+1:]
                orders = permutation(remaining_words, total, pre+[word], iteration+1)

                for order in orders:
                        result.append([word] + order)
        print("----------------")
        print(input_list)
        print(result)
        return result
def permutation2(input_list):
    if len(input_list) == 1:
        return [input_list]
    result = []
    n = len(input_list)
    for i in range(n):
        orders = permutation2(input_list[:i] + input_list[i+1:])
        for order in orders:
            result.append(order+[input_list[i]])
    return result
#print(permutation(["AREA", "BALL", "DEAR", "LADY", "LEAD", "YARD"], 6, [], 0))
#print(permutation2(["A","B","C"]))
def combination(input_list, k):
    if k:
        for i in range(k-1, len(input_list)):
            for c in combination(input_list[:i], k-1):
                nc = c + (input_list[i],)
                yield nc
    else:
        yield tuple()


print([x for x in combination(["1","2","3"], 2)])

