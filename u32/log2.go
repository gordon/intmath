package u32

// Log2 returns log base 2 of n. It's the same as index of the highest
// bit set in n. n == 0 returns 0
func Log2(n uint32) (r uint32) {
	if n >= 1<<16 {
		r += 16
		n >>= 16
	}
	if n >= 1<<8 {
		r += 8
		n >>= 8
	}
	if n >= 1<<4 {
		r += 4
		n >>= 4
	}
	if n >= 1<<2 {
		r += 2
		n >>= 2
	}
	r += n >> 1
	return
}
