package main

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func Test_Allocate(t *testing.T) {
	var expected = map[uint64]int{
		0:          0,
		1:          0,
		2:          0,
		3:          1,
		130:        1,
		131:        2,
		4294967297: 33554432,
		4294967298: 33554432,
		4294967299: 33554433,
	}

	for limit, ints := range expected {
		primes := allocate(limit)
		if len(primes) != ints {
			t.Errorf("allocate() didn't allocate the correct number of ints: limit=%v ints=%v expected=%v", limit, len(primes), ints)
			return
		}
	}

	t.Log("allocate() tests completed")
}

func Test_Sieve(t *testing.T) {
	var expected = map[uint64]string{
		1:        "26ab0db90d72e28ad0ba1e22ee510510",
		100:      "45d886e08500b82881519d5cf5cbe1d6",
		1000:     "2d2382f376350089fd94503d7da478db",
		10000:    "c37d039ddada44d3976ed948c3d0ef21",
		100000:   "ba89921a4ba02bb51ee28d66dbfc3451",
		1000000:  "c13929ee9d2aea8f83aa076236079e94",
		10000000: "60e34d268bad671a5f299e1ecc988ff6",
	}

	for limit, md5sum := range expected {
		primes := allocate(limit)
		sieve(limit, primes)

		m := md5.New()
		fmt.Fprintln(m, 2)
		saveTo(m, primes)
		hash := fmt.Sprintf("%x", m.Sum(nil))

		if hash != md5sum {
			t.Errorf("md5 doesn't match: limit=%v md5=%v expected=%v", limit, hash, md5sum)
		}
	}

	t.Log("sieve() tests completed")
}

func Benchmark_Primes(b *testing.B) {
	b.StopTimer()
	primes := allocate(uint64(b.N))
	b.StartTimer()
	sieve(uint64(b.N), primes)
}
