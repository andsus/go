package change

import "errors"

// Change return array of coin used
func Change(coins []int, amount int) (result []int, err error) {
	if amount < 0 {
		return nil, errors.New("negative amount")
	}

	t := make([][]int, amount+1)
	t[0] = []int{}

	for a := range t {
		for _, c := range coins {
			if a-c < 0 { // ignore, amount < coin
				continue
			}

			p := t[a-c] // amount - change
			if p == nil {
				continue
			}

			if t[a] == nil || len(p)+1 < len(t[a]) {
				t[a] = append([]int{c}, p...)
			}

		}
	}

	if t[amount] != nil {
		return t[amount], nil
	}

	return nil, errors.New("no change")
}
