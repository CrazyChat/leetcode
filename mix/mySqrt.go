package mix

func MySqrt(x int) int {
	l, r := 0, x
	result := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}
