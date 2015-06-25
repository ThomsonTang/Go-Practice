package main

import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }

    for i := 0; i < 5; i++ {
        fmt.Printf("test %s\n", i)
    }
}
