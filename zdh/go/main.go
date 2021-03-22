package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 倒序输出数组函数
func daoxu(arr [10]int) {
	temp := 0
	for i := 0; i < len(arr)/2; i++ {
		temp = arr[10-i-1]
		arr[10-i-1] = arr[i]
		arr[i] = temp
	}
	fmt.Println("倒序打印的数组为:", arr)
}

// 求平均值
func average(arr [10]int) {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}
	fmt.Println("平均值为:", sum/float64(len(arr)))
}

// 求最大值
func max(arr [10]int) {
	max := 0
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	fmt.Printf("最大值是:%d,原数组的最大值的下标为:%d\n", max, index)
}

// 查找是否含有55
func find(arr [10]int, findval int) {
	var num = -1
	var index = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == findval {
			num = arr[i]
			index = i
		}
	}
	if num != -1 {
		fmt.Printf("找到55了，下标为:%d", index)
	} else {
		fmt.Println("没有找到55")
	}
}

func main() {
	/*随机生成1个整数(1-100的范围)保存到数组，并倒序打印以及求平均值，
	求最大值，以及最大值的下标，并查找里面是否含有55*/
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}
	fmt.Println("生成的原数组为:", arr)

	// 倒序输出数组
	daoxu(arr)
	//求平均值
	average(arr)
	//求最大值及下标
	max(arr)
	find(arr, 55)
}