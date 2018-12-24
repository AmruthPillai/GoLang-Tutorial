package main

import "fmt"

const uSixteenBitMax float64 = 65535
const kmphMultiple float64 = 1.609344

type car struct {
	gasPedal       uint16 // min 0, max 65k
	brakePedal     uint16
	streeringWheel int16 // min -32k, max +32k
	topSpeedKmh    float64
}

func (c car) kmph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixteenBitMax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixteenBitMax / kmphMultiple)
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func newerTopSpeed(c car, newSpeed float64) car {
	c.topSpeedKmh = newSpeed
	return c
}

func main() {
	aCar := car{
		gasPedal:       22341,
		brakePedal:     0,
		streeringWheel: 12561,
		topSpeedKmh:    225.0,
	}

	// aCar := car{22341, 0, 12561, 225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmph())
	fmt.Println(aCar.mph())

	aCar.newTopSpeed(500)
	fmt.Println(aCar.kmph())
	fmt.Println(aCar.mph())

	aCar = newerTopSpeed(aCar, 1000)
	fmt.Println(aCar.kmph())
	fmt.Println(aCar.mph())
}
