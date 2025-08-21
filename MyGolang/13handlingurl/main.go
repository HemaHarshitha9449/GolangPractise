package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.example.com/questions/3456/my-document"

func main() {
	fmt.Println("Lets check on handling URL")
	fmt.Println(myurl)
	res, _ := url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Query())
	fmt.Println(res.RawQuery)
	qparams := res.Query()
	fmt.Printf("The type of query params are:%T\n", qparams)
	for _, val := range qparams {
		fmt.Println(val)
	}
	variable := &url.URL{
		Scheme:  "https",
		Host:    "localhost",
		Path:    "/jhfsd",
		RawPath: "user=harshitha",
	}
	fmt.Println(variable.String())
}
