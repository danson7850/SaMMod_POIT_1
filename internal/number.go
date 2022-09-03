package internal

import (
	"math"
)

// k - const that we need to calculate histogram
const k = 20

// LehmerAlgorithm - function that generates fake random numbers
// Parameters:
// a, m - random int numbers
// r - start number
// n - count of iterations (new generated numbers)
// result - array of int numbers
func LehmerAlgorithm(a, m, n int, r float64) (result []float64) {
	for i := 0; i < n; i++ {
		r1 := r
		r = math.Mod(float64(a)*r1, float64(m))
		result = append(result, r/float64(m))
	}
	return result
}

// EstimationCalculation - function which provides calculation of math.expectation, dispersion
// and rms(sqrt from dispersion)
func EstimationCalculation(numbers []float64) (mx, dx, sx float64) {

	mx, dx = 0, 0

	for _, i := range numbers {
		mx += i
	}
	mx /= float64(len(numbers))

	for _, i := range numbers {
		dx += math.Pow(i-mx, float64(2))
	}

	dx /= float64(len(numbers)) - 1

	sx = math.Sqrt(dx)

	return mx, dx, sx
}

func HistogramCalculation() {

}

func UniformityChecker(numbers []float64) float64 {
	count := 0

	for i := 0; i < len(numbers)/2; i++ {
		if math.Pow(numbers[2*i], 2)+math.Pow(numbers[2*i+1], 2) < 1 {
			count++
		}
	}

	return 2 * float64(count) / float64(len(numbers))
}

func AperiodicCalculation() {

}
