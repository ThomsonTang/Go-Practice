package main

import "fmt"

//可以有多个返回值
func swap(x, y string)(string, string) {
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
