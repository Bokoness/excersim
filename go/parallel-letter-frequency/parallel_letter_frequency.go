//Package letter counts the frequency of letters in text
package letter

type FreqMap map[rune]int

// ConcurrentFrequency uses concurreny to calculate
//the frequency of each letter each string in list map
func ConcurrentFrequency(list []string) FreqMap {
	f := FreqMap{}
	c := make(chan FreqMap)
	for _, s := range list {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}
	for range list {
		for i, v := range <-c {
			f[i] += v
		}
	}
	return f
}

// Frequency calculates the frequency of each letter in s
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
