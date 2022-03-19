package main

import "fmt"

//teach user defined types

type Car struct {
	Name    string
	Make    int
	Model   string
	YoM     int //year of manufacture
	FuelEff FuelEfficiency
}

type FuelEfficiency struct {
	Milage       int
	FuelCapacity int
}

func (c Car) GetName() {
	fmt.Println(c.Name)
}

func structs() {
	var bmw = Car{Name: "BMW", Make: 2020, Model: "320i", YoM: 2020}
	fmt.Printf("%+v \n", bmw)

	bmw.YoM = 2019
	fmt.Printf("%+v \n", bmw)

	bmw.FuelEff = FuelEfficiency{Milage: 20, FuelCapacity: 50}

	fmt.Printf("%+v \n", bmw)

	bmw.GetName()

}
