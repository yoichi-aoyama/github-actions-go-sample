package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please specify a command.")
		return
	}
	fmt.Println(args[0])
	fmt.Println(args[1])
}
