package floats

func Find(a []float64, x float64) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func Contains(a []float64, x float64) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Filter(arr []float64, cond func(float64) bool) []float64 {
	result := []float64{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
