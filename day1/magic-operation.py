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
