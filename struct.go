package main

import "fmt"

type car struct {
	gas_pedal      uint16
	steering_wheel int16
}

func main() {
	a_car := car{gas_pedal: 1233,
		steering_wheel: 10}

	fmt.Println(a_car.steering_wheel)
}
