package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Trying on user input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	input, _ := reader.ReadString('\n')
	fmt.Printf("The user input is:%v\n", input)
	fmt.Printf("Type of the input is: %T\n", input)

}
