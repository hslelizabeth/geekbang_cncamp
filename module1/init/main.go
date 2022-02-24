package main

import (
	"fmt"

	_ "geekbang_cncamp/module1/init/a"
	_ "geekbang_cncamp/module1/init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("main function")
}
