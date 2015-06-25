package variable

import "fmt"

func main() {
    var x string = "hello world"
    fmt.Println(x)

    // alternative way
    var y string
    y = "hello, go"
    fmt.Println(y)

    var z string
    z = "first"
    fmt.Println(z)
    z = "second"
    fmt.Println(z)
    z = z + " third"
    fmt.Println(z)

    // equals
    var h string = "hello"
    var w string = "world"
    fmt.Println(h == w)

    var hello = "hello"
    var hi = "hello"
    fmt.Println(hello == hi)

    name := "my name"
    fmt.Println(name)

    var age = 25
    fmt.Println(age)
}
