package main

import(
	"github.com/g3n/engine/math32"
)

func one(x,y float32) float32{
	return x*x-2*y
}

func two(x,y float32) float32{
	return math32.Sin(x)
}

func three(x,y float32) float32{
	return 0
}

func four(x,y float32) float32{
	return math32.Sin(2.5*math32.Cos(x))
}