package main

import (
	"fmt"
)

func Mode(nums []float64) float64 {

	if len(nums) == 0 {
		return 0
	}

	maxCount := 0

	mode := nums[0]

	freq := make(map[float64]int)
	
	for _, n := range nums {
		freq[n]++

		if freq[n] > maxCount || (freq[n] == maxCount && n < mode) {
			maxCount = freq[n]
			mode = n
		}
	}
	return mode
}

func main() {

	nums := []float64{1, 2, 2, 3, 3}
	fmt.Println("Mode of [1, 2, 2, 3, 3]:", Mode(nums))

}
