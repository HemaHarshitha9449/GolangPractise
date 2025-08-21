package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println(ptr)
	number := 24
	ptr = &number
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr - 3
	fmt.Println(*ptr)
}
