package main

import (
	"flag"
	"log"
)

func main() {
	// fmt.Printf("%#v\n", flag.CommandLine)
	// flag.Parse()

	var name string
	// fmt.Printf("%#v\n", flag.CommandLine)
	flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go语言变成之旅", "帮助信息")
	// fmt.Printf("%#v\n", flag.CommandLine)
	flag.Parse()
	// fmt.Printf("%#v\n", flag.CommandLine)
	log.Printf("name: %s", name)
}
