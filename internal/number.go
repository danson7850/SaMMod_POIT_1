package internal

// k - const that we need to calculate histogram
const k = 20

// LehmerAlgorithm - function that generates fake random numbers
// Parameters:
// a, m - random int numbers
// r - start number
// n - count of iterations (new generated numbers)
// result - array of int numbers
func LehmerAlgorithm(a, m, n, r int) (result []int) {

	for i := 0; i < n; i++ {
		r1 := r
		r = (a * r1) % m
		result = append(result, r)
	}
	return result
}

func HistogramCalculation(numbers []int) {

}

func EstimationCalculation() {

}

func UniformityChecker() {

}

func AperiodicCalculation() {

}
