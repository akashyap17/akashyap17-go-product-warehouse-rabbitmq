package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	name := flag.String("name", "World", "a name to say hello")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}
