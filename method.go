package main

import "fmt"

const top_speed = 100

type car struct {
	gas_pedal      uint16
	steering_wheel int16
}

// fungsi kmh() dideklarasikan sebagai method milik struct car
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * top_speed
}

func main() {
	a_car := car{gas_pedal: 1233,
		steering_wheel: 10}

	fmt.Println(a_car.steering_wheel)
	fmt.Println(a_car.kmh())
}
