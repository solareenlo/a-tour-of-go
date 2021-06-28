package main

import "fmt"

func main() {
    i := 42
    f := 3.142
    g := 0.8 + 0.5i
    fmt.Printf("v is of type %T\n", i)
    fmt.Printf("f is of type %T\n", f)
    fmt.Printf("g is of type %T\n", g)
}
