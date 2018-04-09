def miniMaxSum(array):
    #
    # Write your code here.
    #
    sums = []
    for i in range(len(array)):
        shifted = shiftArr(i + 1, array)
        sums.append(sum(shifted[:4]))

    sums.sort()
    print('%s %s' % (sums[0], sums[-1],))


def shiftArr(n, array):
    return array[n:] + array[:n]


miniMaxSum([1, 2, 3, 4, 5])
