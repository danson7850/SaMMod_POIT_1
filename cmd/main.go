package main

import (
	"fmt"
	"math"
	"sammod_1/internal/service"
)

const (
	a = 1473
	m = 2094991
	n = 10000000
	r = 18
)

func main() {
	result := service.LehmerAlgorithm(a, m, n, r)

	mx, dx, sx := service.EstimationCalculation(result)
	fmt.Printf("mx = %.5f, dx = %.5f, sx = %.5f\n", mx, dx, sx)

	un := service.UniformityChecker(result)
	fmt.Printf("π/4 check = %.7f ---> %f\n", un, math.Pi/4)

	per, aper := service.AperiodicCalculation(result, n, m)
	fmt.Println(per, aper)
}
