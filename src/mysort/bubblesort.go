package mysort

import (
	"fmt"
)

func BubbleSort(data []int) []int {
	fmt.Println("Data is : ", data)

	//假设从小往大排
	//从第0个遍历到第n-1个，第n个（最后一个）不用比较。
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	// for _, value := range data{
	fmt.Println("Result is : ", data)

	// }
	//result := []int{0, 9, 9}
	return data
	//遍历原数组
	//比较当前数据和后一个数据，如果当前数据大则交换位置
	//遍历到n-1个结束
}
