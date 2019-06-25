package main

import (
	"fmt"
)

type Gopher struct {
	Gopher string `json:"gopher"`
}

func main() {
	const gopher = "GOPHER"
	gogopher := GOPHER()
	gogopher.Hello()
	gogopher.Gopher = gopher
	fmt.Println(gogopher)
}

func (g *Gopher) Hello() {
	fmt.Println("Hello, World!")
}

func GOPHER() (gopher *Gopher) {
	gopher = &Gopher{Gopher: "gopher"}
	return
}
