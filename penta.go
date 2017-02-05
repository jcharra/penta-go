package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Pentago")

	interactive := flag.Bool("i", false, "interactive")
	port := flag.Int("p", 9977, "Start pentago server listening on this port")

	flag.Parse()

	if *interactive {
		fmt.Println("Starting interactive play ...")
	} else {
		fmt.Println("Start server at port", *port)
	}
}
