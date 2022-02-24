package main

import "fmt"

func main() {
	/*
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
	*/

	/*
		ch := make(chan int, 10)
		go func() {
			fmt.Println("from child thread")
			ch <- 1
		}()
		a := <-ch
		fmt.Println(a)
	*/

	/*
		ch := make(chan int, 10)
		go func() {
			for i := 0; i < 10; i++ {
				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(10)
				fmt.Println("putting: ", n)
				ch <- n
			}
			close(ch)
		}()
		fmt.Println("hello from main")
		for v := range ch {
			fmt.Println("receiving: ", v)
		}
	*/

	ch := make(chan int, 5)
	ch <- 3
	ch <- 4
	// defer close(ch)
	if v, notClosed := <-ch; notClosed {
		fmt.Println(v)
	}

}
