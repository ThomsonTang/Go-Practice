package main

import "fmt"
import "github.com/ThomsonTang/Go-Practice/gobook/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Println(xs)
	avg := math.Average(xs)
	x := "test space"
	fmt.Println(avg, x)
	y := "go lang perfect"
	y = "chage the value"
	fmt.Println(y)
}
