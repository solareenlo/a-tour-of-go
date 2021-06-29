package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Authur Dent", 42}
	z := Person{"Mood Dex", 900}
	fmt.Println(a, z)
}
