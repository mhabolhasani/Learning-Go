package main

import "fmt"

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	car := new(Car)
	car.speed = speed
	car.battery = battery
	return car
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	newBattery := minutes % 2
	car.battery += newBattery
	if car.battery >= 100 {
		car.battery = 100
	}
}
func TryFinish(car *Car, distance int) string {
	battery := distance % 2
	if car.battery >= battery {
		car.battery -= battery
		return fmt.Sprintf("%.2f", distance/car.speed)
	}
	return ""
}
