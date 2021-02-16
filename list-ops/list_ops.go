package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// IntList is List of int
type IntList []int

// Foldr (given a function, a list, and an initial
// accumulator, fold (reduce) each item into the accumulator
// from the right using function(item, accumulator));
func (list IntList) Foldr(fn binFunc, inital int) int {
	for i := len(list) - 1; i >= 0; i-- {
		inital = fn(list[i], inital)
	}
	return inital
}

// Foldl (given a function, a list, and initial accumulator,
// fold (reduce) each item into the accumulator from the left
// using function(accumulator, item));
func (list IntList) Foldl(fn binFunc, inital int) int {
	for _, item := range list {
		inital = fn(inital, item)
	}
	return inital
}

// Filter (given a predicate and a list, return the list of
// all items for which predicate(item) is True);
func (list IntList) Filter(fn predFunc) IntList {
	result := []int{}
	for _, item := range list {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// Length return number of items
func (list IntList) Length() int {
	return len(list)
}

// Map (given a function and a list, return the list of the
// results of applying function(item) on all items);
func (list IntList) Map(fn unaryFunc) IntList {
	result := []int{}
	for _, item := range list {
		result = append(result, fn(item))
	}
	return result
}

// Reverse (given a list, return a list with all the original
// items, but in reversed order);
func (list IntList) Reverse() IntList {
	result := []int{}
	for _, item := range list {
		result = append([]int{item}, result...)
	}
	return result
}

// Append (given two lists, add all items in the second list
// to the end of the first list)
func (list IntList) Append(items IntList) IntList {
	return append(list, items...)
}

// Concat (given a series of lists, combine all items in all
// lists into one flattened list);
func (list IntList) Concat(lists []IntList) IntList {
	result := list
	for _, list := range lists {
		result = append(result, list...)
	}
	return result
}
