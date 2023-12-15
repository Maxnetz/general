package main

import "fmt"

func main() {
	// var name type = value
	// 8 bit = 1 byte
	// 32 bit = 4 bytes
	var a int32 = 1
	fmt.Println(a)

	var msg string = "string"
	fmt.Println(msg)

	// dynamic type
	b := 2
	// convert type
	c := float64(b)
	fmt.Println(b)
	fmt.Println(c)

	// interface == type any
	var d any
	d = "interface string"
	e , ok:= d.(string)
	if !ok {
		panic("err")
	}
	fmt.Println(d)
	fmt.Println(e)
}