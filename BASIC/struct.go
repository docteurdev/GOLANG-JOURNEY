package main

import "fmt"

func main() {
	person := newPerson(Person{name: "oumar adje", email: "oumar@gmail.co", phone: "0142269019"})

	fmt.Println(person)
	fmt.Println("================================")

	bmw := newPersonCar(Car{brand: "BMW", model: "X5", year: 2020})
	opel := newPersonCar(Car{brand: "Opel", model: "Astra", year: 2018})

	fmt.Println(opel.speed())
	fmt.Println(bmw.speed())
}

type Person struct {
	name  string
	email string
	phone string
}

func newPerson(p Person) *Person {
	pers := &Person{
		name:  p.name,
		email: p.email,
		phone: p.phone,
	}

	fmt.Println(pers)
	return pers
}

type Car struct {
	brand string
	model string
	year  int
}

func newPersonCar(c Car) *Car {

	car := &Car{
		brand: c.brand,
		model: c.model,
		year:  c.year,
	}
	return car
}

func (c *Car) speed() int {
	if c.year > 2019 {
		return 200
	}
	return 150
}
