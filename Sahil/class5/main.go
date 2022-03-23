package main

import "fmt"

type Company struct {
	Name        string
	Departments []Department
	Adr         Address
}

type Department struct {
	Name      string
	WorkForce []Employee
}

type Employee struct {
	EmpId       int
	Name        string
	Designation string
	Salary      int
	Address
}

func (e *Employee) ChangeDesignation(newDesignation string) {
	e.Designation = newDesignation
}

func (e *Employee) IncrementSalary(newSalary int) {
	e.Salary = newSalary
}

type Address struct {
	HouseNo    string
	StreetName string
	City       string
	State      string
	Country    string
}

//GetFormattedAddress returns formatted address
func (a Address) GetFormattedAddress() string {
	return a.HouseNo + ", " + a.StreetName + ", " + a.City + ", " + a.State + ", " + a.Country
}

func main() {

	alphabet := Company{
		Name: "Alphabet",
		Adr: Address{
			HouseNo:    "1",
			StreetName: "random street",
			City:       "San jose",
			State:      "California",
			Country:    "USA",
		},
	}
	fmt.Println(alphabet.Adr.GetFormattedAddress())
	itDept := Department{
		Name: "IT",
	}

	john := Employee{
		EmpId:       1,
		Name:        "John",
		Designation: "Engineer",
		Salary:      50000,
		Address: Address{
			HouseNo:    "12",
			StreetName: "mangro road",
			City:       "san jose",
			State:      "California",
			Country:    "USA",
		},
	}
	fmt.Println("John Address ", john.GetFormattedAddress())

	itDept.WorkForce = append(itDept.WorkForce, john)
	alphabet.Departments = append(alphabet.Departments, itDept)

	fmt.Println("Entire Company")
	fmt.Printf("%+v \n", alphabet)

	fmt.Printf("%+v \n\n", john)
	john.IncrementSalary(70000)
	fmt.Printf("%+v \n", john)

	var e Employee
	fmt.Printf("%+v", e)
}
