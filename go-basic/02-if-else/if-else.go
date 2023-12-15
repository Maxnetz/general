package main

import (
	"errors"
	"fmt"
)

func main() {
	score := 70

	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

func genErr() error {
	return errors.New("some error")
}

func checkError() {
	if err := genErr(); err != nil {

	}
}
