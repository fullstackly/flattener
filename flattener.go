package flattener

func Flatten(twoD [][]int) []int {
	oneD := make([]int, 0, len(twoD)*2)

	var top, left int
	bottom, right := len(twoD)-1, len(twoD)-1

	for top <= bottom && left <= right {
		//right
		for i := left; i <= right; i++ {
			oneD = append(oneD, twoD[top][i])
		}
		top++

		//down
		for i := top; i <= bottom; i++ {
			oneD = append(oneD, twoD[i][right])
		}
		right--

		//left
		for i := right; i >= left; i-- {
			oneD = append(oneD, twoD[bottom][i])
		}
		bottom--

		//up
		for i := bottom; i >= top; i-- {
			oneD = append(oneD, twoD[i][left])
		}
		left++
	}
	return oneD
}
