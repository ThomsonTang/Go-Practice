package slice

import "fmt"

func main() {
	// in this case y has been created with a length of 0
	var y []int
	fmt.Println(y)

	// create a slice, although the slice can never be longer than the array,
	// they can be smaller.
	x := make([]int, 5)
	x[0] = 10
	x[1] = 11
	fmt.Println(x)

	// make function allows a 3rd parameter
	// 10 represents the capacity of the underlying array which the slice points to.
	f := make([]float64, 5, 10)
	f[1] = 24
	f[4] = 25
	fmt.Println(f)

	// another way to create slices is to use the [low : high] expression
	arr := []float64{1, 2, 3, 4, 5}
	a := arr[0:5] // [1, 2, 3, 4, 5]
	a = arr[0:4]
	fmt.Println(a)

	name := []string{"pop", "tim", "tony", "manu"}
	n := name[0:4]
	// same as name[0:len(name)]
	n = name[0:]
	n = name[:4]
	n = name[:]
	fmt.Println(n)

}
