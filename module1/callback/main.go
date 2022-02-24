package main

func main() {
	func() {
		println("hello world")
	}()
	DoOperation(2, increase)
	DoOperation(2, decrease)
}

func increase(a, b int) {
	println("decrease result is:", a+b)
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}
