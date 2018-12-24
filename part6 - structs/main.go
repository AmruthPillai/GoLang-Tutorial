package main

import "fmt"

type car struct {
	gasPedal       uint16 // min 0, max 65k
	brakePedal     uint16
	streeringWheel int16 // min -32k, max +32k
	topSpeedKmh    float64
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
}
