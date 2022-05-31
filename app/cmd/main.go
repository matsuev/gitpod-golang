package main

import "fmt"

type AAA struct {
	I int
	S string
}

func main() {
	A := AAA{
		I: 124,
		S: "ghsjfkgsdgf",
	}
	fmt.Printf("%#v\n", A)
}
