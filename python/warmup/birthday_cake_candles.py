def birthdayCakeCandles(n, ar):
#
# Write your code here.
#
    int(n)
    max_height = max(ar)
    matching = sum([ar[i]==max_height for i in range(n)])
    return matching


print(birthdayCakeCandles(4, [3, 2, 1, 3]))
