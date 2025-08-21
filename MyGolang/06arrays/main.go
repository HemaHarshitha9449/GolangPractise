package main

import "fmt"

func main() {
	var arrayList [7]string
	arrayList[0] = "Hema"
	arrayList[1] = "Hema"
	arrayList[2] = "Hema"
	arrayList[5] = "Hema"
	fmt.Println(arrayList)
	fmt.Println(len(arrayList))
	var List2 = [5]int{3, 4, 5, 6, 6}
	fmt.Println(List2)
	fmt.Println(len(List2))
}
