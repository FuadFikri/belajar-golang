package main

import "fmt"

func main() {
	x := 15
	a := &x
	fmt.Println(a)  //0xc000054058
	fmt.Println(*a) //15
	*a = 5
	fmt.Println(x) //5

	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)

}
