package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	//city := flag.String("city", "GoWorld", "specify you place")
	var times int
	flag.IntVar(&times, "times", 1000, "times of hello world")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Printf("input parameter is: %s %d\n", *name, times)
	//fullString := fmt.Sprintf("Hello %s from Go @ %s\n", *name, *city)
	//fmt.Println(fullString)
	fmt.Printf("Hello %s from Go %d times\n", *name, times)
}
