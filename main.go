package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "Ok"

	fmt.Print(reflect.TypeOf(message))
	fmt.Println("Hello world!")
	fmt.Println(message)

	for i := 0; i < 5; i++ {
		println("FuCk!")
	}
}
