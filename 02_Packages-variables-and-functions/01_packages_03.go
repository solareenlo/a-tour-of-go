package main

import (
    crand "crypto/rand"
    "fmt"
    "math"
    "math/big"
    "math/rand"
)

func main() {
    seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
    rand.Seed(seed.Int64())
    fmt.Println("My favarit number is", rand.Intn(10))
}
