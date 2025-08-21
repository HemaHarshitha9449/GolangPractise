package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Now lets check on switches")
	number := rand.Intn(2) + 1
	switch number {
	case 1:
		fmt.Println("The number is 1")
		fallthrough
	case 2:
		fmt.Println("The number is 2")
	default:
		fmt.Println("The number is default")
	}
}
