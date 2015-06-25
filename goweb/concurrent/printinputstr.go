package main

import "fmt"

func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	f(0) //正常调用函数f()，其与主goroutine，即main函数同步运行

	go f(1) //在一个goroutine中调用函数f()，这个goroutine将与主goroutine并发执行

	// 可以在一个goroutine中调用执行匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//Scanln()使得函数f()执行时展示打印的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	fmt.Println("done")
}
