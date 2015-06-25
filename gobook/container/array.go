package main

import "fmt"

func main() {
	var x [5]int
	x[1] = 10
	x[3] = 100
	fmt.Println(x)
	fmt.Println(x[2])

	/*
	   var y [5]float64
	   y[0] = 10
	   y[1] = 100
	   y[2] = 1000
	   y[3] = 10000
	   y[4] = 100000
	*/
	y := [5]float64{
		10,
		11,
		12,
		13,
		// 14,
	}
	fmt.Println(y)

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println("average: ", total/float64(len(y)))

	// range关键字类似于其他语言中的foreach
	var anotherTotal float64 = 0
	for _, value := range y {
		anotherTotal += value
	}
	fmt.Println("average: ", anotherTotal/float64(len(y)))

}
