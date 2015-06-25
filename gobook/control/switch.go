package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        switch i {
        case 0: fmt.Println("zero")
        case 1: fmt.Println("first")
        case 2: fmt.Println("two")
        case 3: fmt.Println("three")
    default: fmt.Println("unkown error")
        }
    }
}
