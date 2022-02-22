package main

import (
	"fmt"
)

func main() {
	mySlice1 := new([]int)
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)
	fmt.Printf("mySlice1 %+v\n", mySlice1)
	fmt.Printf("mySlice2 %+v, length: %d, capacity: %d\n", mySlice2, len(mySlice2), cap(mySlice2))
	fmt.Printf("mySlice3 %+v, length: %d, capacity: %d\n", mySlice3, len(mySlice3), cap(mySlice3))
	fmt.Printf("mySlice4 %+v, length: %d, capacity: %d\n", mySlice4, len(mySlice4), cap(mySlice4))
}
