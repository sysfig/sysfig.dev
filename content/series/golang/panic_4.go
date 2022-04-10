package main

import "fmt"

type Fruit struct {
	Name string
}

func addCream(f *Fruit) {
	defer fmt.Println("addCream() finished executing")
	f.Name = f.Name + " with cream"
}

func main() {
	defer fmt.Println("main() finished executing")
	var strawberry *Fruit
	addCream(strawberry)
	fmt.Println("Post-panic code do not continue executing")
}
