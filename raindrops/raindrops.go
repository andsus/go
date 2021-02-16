package raindrops

import (
	"strconv"
)

// Convert int to sound of raindrops in string
func Convert(n int) string {
	res := ""
	factorMap := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	for _, v := range []int{3, 5, 7} {
		if n%v == 0 {
			res += factorMap[v]
		}
	}
	if len(res) == 0 {
		return strconv.Itoa(n)
	}
	return res
}
