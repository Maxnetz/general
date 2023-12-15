package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	// a := new(int)
	// b := 10

	// // *a = b
	// a = &b

	// fmt.Printf("%p\n", a)
	// fmt.Printf("%p\n", &b)
	// fmt.Printf("%p\n", &a)
	// fmt.Printf("%p\n", *a)

	x := 5
	y := 10

	fmt.Println(x, y) // 5, 10
	swap(&x, &y)
	fmt.Println(x, y) // 10, 5
}
