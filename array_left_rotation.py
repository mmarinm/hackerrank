def array_left_rotation(a, n, k):
    shiftedArray = a[:k]
    rest = a[k:]
    return rest + shiftedArray


print(array_left_rotation([1,2,3,4,5], 5, 4))