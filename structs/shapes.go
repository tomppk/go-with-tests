package main

import "math"

// In Go interface resolution is implicit. If the type you pass in matches what
// the interface is asking for, it will compile.
// Both Rectangle and Circle types have method Area() that returns float64 so
// they satisfy the Shape interface automatically.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

// (r Rectangle) declares the "receiver" of this method. So the receiver is the
// type Rectangle struct
// This method is bound to the receiver and it can be called on a variable of that
// type. Syntax:
// func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }