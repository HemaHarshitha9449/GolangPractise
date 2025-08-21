package main

import "fmt"

func main() {
	fmt.Println("Now lets check on the maps")
	mlist := make(map[int]string)
	mlist[1] = "Hema"
	mlist[2] = "You"
	mlist[3] = "will"
	mlist[4] = "be"
	mlist[5] = "okay"
	fmt.Println(mlist)
	delete(mlist, 3)
	fmt.Println(mlist)
	for key, value := range mlist {
		fmt.Printf("The key is %v and value is %v\n", key, value)
	}
}
