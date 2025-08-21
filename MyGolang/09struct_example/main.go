package main

import "fmt"

func main() {
	fmt.Println("Now lets check on structs")
	employee := Person{"Hema", 123, 23, "Banglore"}
	fmt.Println(employee)
	fmt.Printf("The values are %+v\n", employee)
	fmt.Printf("The name of employee is %v and age is %v", employee.Name, employee.Age)
}

type Person struct {
	Name    string
	id      int
	Age     int
	Address string
}
