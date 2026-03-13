package popcount

// TODO: Return to this task after the chapter 11

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var sum uint64
	var mask uint64 = 1
	for i := 0; i < 64; i++ {
		sum += (x & mask) >> i
		mask <<= 1
	}

	return int(sum)
}
