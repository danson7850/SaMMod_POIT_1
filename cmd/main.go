package main

import (
	"fmt"
	"math"
	"sammod_1/internal"
)

func main() {
	result := internal.LehmerAlgorithm(1473, 2094991, 10000000, 18)

	mx, dx, sx := internal.EstimationCalculation(result)
	fmt.Printf("mx = %.5f, dx = %.5f, sx = %.5f\n", mx, dx, sx)

	un := internal.UniformityChecker(result)
	fmt.Printf("π/4 check = %.7f ---> %f", un, math.Pi/4)
}
