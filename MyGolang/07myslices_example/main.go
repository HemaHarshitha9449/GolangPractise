package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start with slices")
	var sliceList []int
	fmt.Println(sliceList)
	sliceList = append(sliceList, 2, 3, 4, 5)
	fmt.Println(sliceList)
	sliceList = sliceList[2:]
	fmt.Println(sliceList)
	newSlice := make([]string, 3)
	newSlice[0] = "Hello"
	newSlice[1] = "Name"
	newSlice[2] = "Naam"
	newSlice = append(newSlice, "You", "Are")
	fmt.Println(newSlice)
	sliceindex := []string{"Are", "You", "An", "Enemy"}
	var index int = 3
	sliceindex = append(sliceindex[:index], sliceindex[index+1:]...)
	fmt.Println(sliceindex)
}
