package series

// All function return all the series
func All(n int, s string) []string {
	series := make([]string, 0)

	for len(s) >= n {
		_s, _ := First(n, s)
		series = append(series, _s)
		s = s[1:]
	}

	return series
}

// UnsafeFirst function return series
func UnsafeFirst(n int, s string) string {

	return s[:n]
}

// First function
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
