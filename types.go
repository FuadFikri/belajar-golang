package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}
func multiple(a, b string) (string, string) {
	return a, b
}
func main() {
	// var num1, num2 float64 = 5.6, 9.5
	num1, num2 := 5.6, 9.5

	fmt.Println(add(num1, num2))
	fmt.Println(multiple("hey", "hi"))

	var a int = 82
	var b float64 = float64(a)

	x := a
}
