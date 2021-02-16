package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts word occurances using concurrent processes
func ConcurrentFrequency(s []string) FreqMap {
	// Map the frequency function over all the words
	channel := make(chan FreqMap, 10) // buffer channel to 10 slots
	for _, words := range s {
		go func(w string) { // goroutine
			channel <- Frequency(w)
		}(words)
	}

	// Reduce the results to a single map
	frequency := FreqMap{}
	for range s {
		for key, value := range <-channel { // read from channel
			frequency[key] += value
		}
	}
	return frequency
}
