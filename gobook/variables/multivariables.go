package multivar

import "fmt"

var (
    a = 4
    b = 5
    c = "test"
)

func main() {
    fmt.Println("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2
    fmt.Println(output)
}
