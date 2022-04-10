package main

import "fmt"

type Fruit struct {
	Name string
}

func addCream(f *Fruit) {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Panic occurred and recovered:", error)
		}
	}()
	defer fmt.Println("addCream() finished executing")
	f.Name = f.Name + " with cream"
	fmt.Println("Post-panic code in panicking function does not continue to execute")
}

func main() {
	defer fmt.Println("main() finished executing")
	var strawberry *Fruit
	addCream(strawberry)
	fmt.Println("Post-panic code in calling function continues to execute")
}
