package main

import(
	"math"
)

import(
	"github.com/g3n/engine/math32"
)

func one(x,y float32) float32{
	return x*x-2*y
}

func one_true(x float32) float32{
	return float32(math.Exp(float64(-2*x)))+(x*x)/2-x/2+1/4
}

func two(x,y float32) float32{
	return math32.Sin(x)
}

func two_true(x float32) float32{
	return -math32.Cos(x)
}

func three(x,y float32) float32{
	return float32(math.Log(float64(x+1)))
}

func three_true(x float32) float32{
	return -x+x*float32(math.Log(float64(x+1)))+float32(math.Log(float64(x+1)))
}

func four(x,y float32) float32{
	return 1/(x+1)
}

func four_true(x float32) float32{
	return float32(math.Log(float64(x+1)))
}