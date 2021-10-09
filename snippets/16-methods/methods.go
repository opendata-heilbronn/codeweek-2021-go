package main

import "fmt"

type Car struct {
	Registration   string
	MileageKm      int64
	LitersPer100Km float64
}

func (c *Car) Drive(kilometers int64) {
	c.MileageKm += kilometers
}

func (c Car) GasUsed() float64 {
	return float64(c.MileageKm) / 100 * c.LitersPer100Km
}

func main() {
	myCorsa := Car{Registration: "HN-PH 404", LitersPer100Km: 7.2}
	myCorsa.Drive(135)
	fmt.Printf("mileage: %dkm, liters of gas: %f", myCorsa.MileageKm, myCorsa.GasUsed())
}
