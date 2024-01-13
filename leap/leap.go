// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear check if the year is leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (!(year%100 == 0) || year%400 == 0)
}
