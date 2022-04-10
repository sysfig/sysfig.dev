package main

func bar() {
	panic("baz")
}

func foo() {
	bar()
}

func main() {
	foo()
}
