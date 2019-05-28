package main

import (
	"fmt"
	"math"
	"math/rand"
)

func akar() {
	fmt.Println("akar dari 9 adalah", math.Sqrt(9))
}

func acak() {
	fmt.Println("acak  antara 1-99 ", rand.Intn(99))
}

func main() {
	akar()
	acak()
}
