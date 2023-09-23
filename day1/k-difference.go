package day1
/*
给定一个数组arr，求差值为k的去重数字对
*/


func kDifference(arr []int,k int) [][]int{
	mp:=make(map[int]int)
	for _,v:=range arr{
		mp[v]=1
	}
	var res [][]int
	for i:=range mp{
		if _,ok:=mp[i+k];ok{
			res=append(res,[]int{i,i+k})
		}
	}
	return res
	
}
