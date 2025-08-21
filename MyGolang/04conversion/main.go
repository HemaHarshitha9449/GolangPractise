package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your rating betwen 1 to 4:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newinput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The new value of rating is:", newinput+1)
	}
}
