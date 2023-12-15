package main

import (
	"fmt"

	"github.com/Maxnetz/03-function/module"
)

func main() {
	fmt.Println(sum("Max", 10, 20))

	// local variable of main()
	x := 10
	changer(x) // x still = 10

	// import module
	fmt.Println(module.Sum("Max2", 10, 20))
}

func sum(name string, a, b int) (int, string) {
	return a + b, "hello" + name
}

func changer(x int) {
	// local variable of changer
	x = 20
}
