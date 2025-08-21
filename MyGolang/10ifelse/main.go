package main

import "fmt"

func main() {
	var num int = 3
	if num < 3 {
		fmt.Printf("the number is less than 3")
	} else if num > 3 {
		fmt.Printf("The number is greater than 3")
	} else {
		fmt.Printf("The number is equal to 3\n")
	}
	if num := 5; num < 10 {
		fmt.Println("The number is less than 10")
	} else {
		fmt.Println("The number is greater than 10")
	}
}
