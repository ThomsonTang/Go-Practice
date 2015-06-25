package main

import "fmt"

func main() {
    const x string = "Hello World"
    fmt.Println(x)

    const y string = "cannot change"
    //y = "some other string" //cannot assign to y
}
