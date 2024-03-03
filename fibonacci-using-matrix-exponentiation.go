package main

var calls int

func matrixMul(m1 [2][2]int, m2 [2][2]int) [2][2]int {
	var ans [2][2]int
	ans[0][0] = m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0]
	ans[0][1] = m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]
	ans[1][0] = m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0]
	ans[1][1] = m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]
	return ans
}

func powerMatrix(power int) [2][2]int {
	calls++
	base := [2][2]int{{1, 1}, {1, 0}}
	if power == 0 {
		return [2][2]int{{1, 0}, {0, 1}}
	}
	if power == 1 {
		return base
	}

	value := powerMatrix(power / 2)
	ans := matrixMul(value, value)
	if power%2 == 1 {
		ans = matrixMul(ans, base)
	}
	return ans
}
