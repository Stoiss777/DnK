package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var sum int
	for x != 0 {
		x &= x - 1
		sum++
	}
	return sum
}
