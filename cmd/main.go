package main

import (
	"fmt"
	"sammod_1/internal"
)

func main() {
	result, xResult := internal.LehmerAlgorithm(2, 2, 2, 2)

	for i := range result {
		fmt.Printf("%.2f || %.2f\n", result[i], xResult[i])
	}

	mx, dx, sx := internal.EstimationCalculation(xResult)
	fmt.Printf("mx = %.5f, dx = %.5f, sx = %.5f", mx, dx, sx)
}
