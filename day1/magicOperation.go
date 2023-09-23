package day1
/*
给一个包含n个整数元素的集合a，一个包含m个整数元素的集合b。
定义magic操作为，从一个集合中取出一个元素，放到另一个集合里，
且操作过后每个集合的平均值都大于操作前。

注意以下两点：
1）不可以把一个集合的元素取空，这样就没有平均值了
2）值为x的元素从a移到b，但集合a中已经有值为x的元素，则a的平均值不变（因为集合元素不会重复），b的平均值
可能会改变（因为x被取出了）
问最多可以进行多少次magic操作？
*/
import(
	"sort"
)

func magicOperation(a,b []int)int{
	var(
		bigList []int
		smallList []int
		bigSum float64
		smallSum float64
	)
	
	//计算平均值
	averageA := average(a)
	averageB := average(b)
	//如果a的平均值等于b的平均值，那么就不用操作了
	if averageA==averageB{
		return 0
	}else{ //否则就要操作，从大的平均值的集合中取出一个元素，放到小的平均值的集合中

		if averageA>averageB{
			bigList=a
			smallList=b
			bigSum=sum(a)
			smallSum=sum(b)
		}else{
			bigList=b
			smallList=a
			bigSum=sum(b)
			smallSum=sum(a)
		}
		sort.Ints(bigList)
		smallSet:=make(map[int]int)
		for _,v:=range smallList{
			smallSet[v]=1
		}
		bigSize:=len(bigList) //平均值大的集合还剩几个数
		smallSize:=len(smallList) //平均值小的集合还剩几个数
		count:=0 //操作次数
		for i:=0;i<len(bigList);i++{ //小->大
			cur:=float64(bigList[i])
			if cur<average(bigList)&&cur>average(smallList)&&!exist(smallList,bigList[i]){
				bigSum-=cur
				smallSum+=cur
				bigSize--
				smallSize++				
				smallList=append(smallList,bigList[i])
				count++
			}
		}
		return count
		
	}
		
}

func sum(a []int)float64{
	sum:=0.0
	for _,v:=range a{
		sum+=float64(v)
	}
	return sum
}
func average(a []int)float64{
	return sum(a)/float64(len(a))
}


func exist(a []int,target int)bool{
	for _,v:=range a{
		if v==target{
			return true
		}
	}
	return false
}