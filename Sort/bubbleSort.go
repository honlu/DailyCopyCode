package main

import "fmt"

/*
冒泡排序：重复地走访要排序地数列，一次比较两个元素，如果它们的顺序错误就把它们交换。
走访数列的工作是重复地进行直到没有再需要交换。

算法步骤：
	1、比较相邻地元素。如果第一个比第二个大，就交换它们两个。
	2、对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对，这样结束一次循环后最后的元素就是最大的数。
	3、除了最后一个元素，再次针对所有的元素重复以上的步骤。
	4、重复1-3，直到排序完成
稳定

时间复杂度：
	最佳：O(n)
	最差：O(n^2)
	平均：O(n^2)


什么时候最快？
	当要排序的序列已经是正序。

什么时候最好？
	当输入的数据是反序时
*/

func bubbleSort(arr []int) []int {
	length := len(arr)
	// 第一层循环代表：要‘循环’的次数
	for i := 0; i < length; i++ {
		// 第二层循环：一次‘循环比较’中要比较的次数
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] { // 升序
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var arr []int
	arr = []int{1, 5, 3, 6}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}