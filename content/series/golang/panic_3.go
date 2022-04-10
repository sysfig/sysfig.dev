package main

import "fmt"

type Fruit struct {
	Name string
}

func addCream(f *Fruit) {
	f.Name = f.Name + " with cream"
}

func main() {
	var strawberry *Fruit
	addCream(strawberry)
	fmt.Println("Post-panic code do not continue executing")
}
