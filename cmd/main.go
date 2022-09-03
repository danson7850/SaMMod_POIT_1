package main

import "sammod_1/internal"

func main() {
	result := internal.LehmerAlgorithm(3, 5, 3, 1)
	for _, i := range result {
		println(i)
	}
}
