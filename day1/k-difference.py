# 给定一个数组arr，求差值为k的去重数字对
# 例如：arr = [3, 1, 5, 4, 2], k = 2
from typing import List


def k_difference(arr: List[int], k: int) -> List[List[int]]:
    res = []
    hash_map = {}
    for i in arr:
        hash_map[i] = 1
    for i in hash_map:
        if i + k in hash_map:
            res.append([i, i + k])
    return res


print(k_difference([3, 1, 5, 4, 2], 2))
