package main

import "fmt"

type Key struct {
	x int
	y int
}

func canCross(positions []int, cur int, jumpLen int, state map[int]bool, ans map[Key]bool) bool {
	nextState := cur + jumpLen
	if !state[nextState] {
		return false
	}

	len := len(positions)
	lastState := positions[len-1]
	if lastState == nextState {
		ans[Key{cur, jumpLen}] = true
		return true
	}
	if nextState > lastState {
		return false
	}

	for _, newLen := range []int{-1, 0, 1} {
		if jumpLen+newLen <= 0 {
			continue
		}
		if canCross(positions, nextState, jumpLen+newLen, state, ans) {
			ans[Key{cur, jumpLen}] = true
			return true
		}
	}
	ans[Key{cur, jumpLen}] = false
	return false
}

func frogJump() {
	positions := []int{0, 1, 2, 3, 4, 8, 9, 11}
	m1 := map[int]bool{}
	for _, v := range positions {
		m1[v] = true
	}
	m2 := map[Key]bool{}
	fmt.Println(canCross(positions, 0, 1, m1, m2))
}
