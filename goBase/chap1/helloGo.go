package main

import (
	"fmt"
)

func main() {
	name := "123"
	var position = &name
	n, m := 1,2
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(name)
	fmt.Println(position)
}

