package scope 

import "fmt"

var x string = "hello world"

func main() {
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}

