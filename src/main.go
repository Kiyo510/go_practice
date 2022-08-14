package main

import "fmt"

type Vehicle interface {
	Accelerate()
	Brake()
}

type Car struct {
	color string
}

func (c Car) Accelerate() {
	fmt.Println("車が加速する")
}

func (c Car) Brake() {
	fmt.Println("車がブレーキをかける")
}

type Bike struct {
	color string
}

func (c Bike) Accelerate() {
	fmt.Println("バイクが加速する")
}

func (c Bike) Brake() {
	fmt.Println("バイクがブレーキをかける")
}

func drive(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Brake()
}

func main() {
	car := Car{color: "red"}
	drive(car)

	bike := Bike{color: "blue"}
	drive(bike)
}
