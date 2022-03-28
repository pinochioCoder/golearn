package main

import (
	"errors"
	"fmt"
	"os"
)

type Car interface {
	Milage() error
	ExpectedLife()
}

//Custom BMW error
type BMWFuelError struct{}

func (b BMWFuelError) Error() string {
	return "BMW Fuel based error"
}

type BMW struct {
	FuelCapacity      int
	YearOfManufacture int
}

func (b *BMW) Milage() (err error) {
	if b.FuelCapacity == 0 {
		// err = errors.New("fuel cap zero")
		err = BMWFuelError{}

	}
	fmt.Println(b.FuelCapacity * 5)
	return
}

func (b *BMW) ExpectedLife() {
	fmt.Println(b.YearOfManufacture + 25)
}

type AUDI struct {
	FuelCapacity      int
	YearOfManufacture int
}

func (b *AUDI) Milage() error {
	var err error
	if b.FuelCapacity == 0 {
		err = errors.New("fuel cap zero")
	}
	fmt.Println(b.FuelCapacity * 5)
	return err
}

func (b *AUDI) ExpectedLife() {
	fmt.Println(b.YearOfManufacture + 25)
}

func main() {
	var simpleCar Car

	bmCar := &BMW{FuelCapacity: 0, YearOfManufacture: 2020}
	auCar := &AUDI{FuelCapacity: 40, YearOfManufacture: 2000}

	simpleCar = bmCar

	err := simpleCar.Milage()
	if err != nil {
		fmt.Println("Error with BMW: ", err)
		os.Exit(1)
	}
	simpleCar.ExpectedLife()

	simpleCar = auCar

	simpleCar.Milage()
	simpleCar.ExpectedLife()

	var dum interface{} // -> type and value

	dum = 5
	fmt.Println(dum)

	dum = "hello"
	fmt.Println(dum)

	dum = simpleCar
	fmt.Println(dum)

	add(5, 5)

	add("hello ", "world")

}

func add(a, b interface{}) {
	switch a.(type) {
	case string:
		fmt.Println(a.(string) + b.(string))
	case int:
		fmt.Println(a.(int) + b.(int))
	}
}

type Bank interface {
	AddFunds(AccNo int) error
	GetAccountInfo(AccNo int) interface{}
	WithdrawFunds(AccNo int) error
	OpenAccount(AccountData interface{}) (interface{}, error)
}
