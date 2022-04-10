package main

func sum(x, y interface{}) int {
	return x.(int) + y.(int)
}

func main() {
	var x interface{} = "4"
	var y interface{} = 5
	sum(x, y)
}
