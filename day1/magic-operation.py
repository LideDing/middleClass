"""
给一个包含n个整数元素的集合a，一个包含m个整数元素的集合b。
定义magic操作为，从一个集合中取出一个元素，放到另一个集合里，
且操作过后每个集合的平均值都大于操作前。

注意以下两点：
1）不可以把一个集合的元素取空，这样就没有平均值了
2）值为x的元素从a移到b，但集合a中已经有值为x的元素，则a的平均值不变（因为集合元素不会重复），b的平均值
可能会改变（因为x被取出了）
问最多可以进行多少次magic操作？
"""
from typing import List


def magic_operation(arr1: List[int], arr2: List[int]) -> int:
    arr1_average = average(arr1)
    arr2_average = average(arr2)
    # 选出平均值较大的一个
    if arr1_average > arr2_average:
        more_arr, less_arr = arr1, arr2
        more_sum, less_sum = sum(arr1), sum(arr2)
    else:
        more_arr, less_arr = arr2, arr1
        more_sum, less_sum = sum(arr2), sum(arr1)
    more_arr.sort()  # 对较大的集合排序
    set_less = set(less_arr)  # 将较小的集合转换为set
    more_size, less_size = len(more_arr), len(less_arr)  # 记录两个集合的长度
    count = 0  # 记录操作次数
    for i in range(more_size):
        cur = more_arr[i]  # 当前元素
        if average(more_arr) > cur > average(less_arr) and cur not in set_less:
            # 如果当前元素小于较大集合的平均值，大于较小集合的平均值，且不在较小集合中
            more_sum -= cur
            more_size -= 1
            less_sum += cur
            less_size += 1
            set_less.add(cur)
            count += 1
    return count


def average(arr: List[int]) -> float:
    return sum(arr) / len(arr)


if __name__ == '__main__':
    arr1 = [1, 2, 3, 4, 5]
    arr2 = [6, 7, 8, 9, 10]
    print(magic_operation(arr1, arr2))
