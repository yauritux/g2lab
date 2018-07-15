package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16
	brake_pedal uint16
	streering_wheel int16
	top_speed_kmh float64
}

func (c car) kmh() float64 { // This is method (value receivers) since it's bounded to car's struct
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 { // this is also method (value receivers)
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

func main() {
	a_car := car{
		gas_pedal: 65000,
		brake_pedal: 0,
		streering_wheel: 12561,
		top_speed_kmh: 225.0,
	}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}