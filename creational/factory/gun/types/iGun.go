package main

type IGun interface {
	setName(name string)
	setPower(power int)
	GetName() string
	GetPower() int
}