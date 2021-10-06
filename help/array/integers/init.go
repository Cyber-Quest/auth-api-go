package integers

func Find(a []int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
