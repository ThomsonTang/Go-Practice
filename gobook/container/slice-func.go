package main

import "fmt"

func main() {
	// append 在给定的slice上追加元素
	slice1 := []int{1, 2, 3, 4}
	slice2 := append(slice1, 5, 6)
	fmt.Println(slice1, slice2)

	// copy 可以用来复制slice，返回值为一个整数，表示copy的元素个数
	slice3 := make([]int, 2)
	result := copy(slice3, slice1)
	fmt.Println(slice1, slice3, result)
}
