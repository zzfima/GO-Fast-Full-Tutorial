package main

import "fmt"

func main() {
	fmt.Println(sumInt([]int{4, 5, 6}))
	fmt.Println(sumFloat([]float32{4.1, 5.1, 6.1}))

	fmt.Println(sumGeneric([]int{4, 5, 6}))
	fmt.Println(sumGeneric([]float32{4.1, 5.1, 6.1}))
}

func sumInt(nums []int) int {
	var sum int = 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sumFloat(nums []float32) float32 {
	var sum float32 = 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sumGeneric[T int | float32 | float64](nums []T) T {
	var sum T = 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
