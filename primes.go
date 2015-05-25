package main

func allocate(limit uint64) []uint64 {
	bits := (limit - 1) / 2

	if bits < 0 || limit <= 2 {
		bits = 0
	}

	mod := bits % 64
	ints := (bits - mod) / 64

	if mod != 0 {
		ints++
	}

	primes := make([]uint64, ints)

	if mod != 0 && ints > 0 {
		// if the last bitmask isn't complete, set all remaining bits to 1
		primes[ints-1] |= (0xFFFFFFFFFFFFFFFF << mod)
	}

	return primes
}

func sieve(limit uint64, primes []uint64) {
	if limit < 3 {
		return
	}

	for k, i, run := uint64(0), uint64(3), true; run; k++ {
		for l := uint8(0); l < 64; l, i = l+1, i+2 {
			if (primes[k]>>l)&1 == 1 {
				// number was already marked as composite
				continue
			}

			sqr := i * i
			if sqr > limit {
				run = false
				break
			}

			for d := i * 2; sqr <= limit; sqr += d {
				// mark all odd multiples from i*i to limit as composites
				p := uint64((sqr - 3) / 2)
				primes[p/64] |= (1 << (p % 64))
			}
		}
	}
}
