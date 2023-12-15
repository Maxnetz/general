package main

import "fmt"

func arrayWatcher(data []int) {
	data[0] = 0
}

func arrayAppend(data *[]int) {
	*data = append(*data, 5)
}

func main() {
	// Array

	// a := [3]int{1, 2, 3}
	// fmt.Println(a[1])

	// Slice
	a := []int{1, 2, 3}
	a = append(a, 4)

	fmt.Println(a)

	arrayWatcher(a)

	fmt.Println(a)

	arrayAppend(&a)

	fmt.Println(a)

	// empty slice
	b := make([]int, 0)

	b = append(b, 4)

	fmt.Println(b)

	arrayAppend(&b)

	// Looping Slice & Array
	for i := range b {
		fmt.Println(i)
	}

	for _, v := range b {
		fmt.Println(v)
	}
}
