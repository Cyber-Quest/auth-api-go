package bools

func Find(a []bool, x bool) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func Contains(a []bool, x bool) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Filter(arr []bool, cond func(bool) bool) []bool {
	result := []bool{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
