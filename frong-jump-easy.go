package main

import (
	"fmt"
	"math"
)

func frogJumpEasy(N int, heights []int) {
	energy := make([]int, N)
	for i := 1; i < N; i++ {
		energy[i] = math.MaxInt
	}
	// this is the tabulation method
	for i := 0; i < N; i++ {
		if i+1 < N {
			energy[i+1] = int(math.Min(float64(energy[i+1]), float64(energy[i])+math.Abs(float64(heights[i]-heights[i+1]))))
		}
		if i+2 < N {
			energy[i+2] = int(math.Min(float64(energy[i+2]), float64(energy[i])+math.Abs(float64(heights[i]-heights[i+2]))))
		}
	}
	fmt.Println(energy[N-1])
}
