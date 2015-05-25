//+build !test

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb"
)

func main() {
	limit := flag.Uint64("l", 100000, "Upper limit for prime searching")
	write := flag.Bool("s", false, "Save the prime numbers to primes.txt")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "A quick test implementing the Sieve of Eratostenes in Go")
		fmt.Fprintln(os.Stderr, "It searches for prime numbers up to the specified limit")
		fmt.Fprintf(os.Stderr, "Usage: %s [-l limit] [-s]\n\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Parse()

	primes := allocate(*limit)

	fmt.Printf("Finding all primes <= %v\n", *limit)
	start := time.Now()

	sieve(*limit, primes)

	fmt.Printf("Done in %v.\n", time.Since(start))

	if *write {
		file, err := os.Create("primes.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fmt.Printf("\nSaving them to primes.txt\n")
		start := time.Now()

		if *limit >= 2 {
			fmt.Fprintln(file, 2)
		}

		output(file, primes)

		fmt.Printf("Done in %v.\n", time.Since(start))
	}
}

func output(w io.Writer, primes []uint64) {
	bar := pb.StartNew(len(primes))
	bar.SetMaxWidth(80)
	bar.ShowCounters = false

	for k, p := range primes {
		bar.Increment()

		for l := uint8(0); l < 64; l++ {
			if (p>>l)&1 == 0 {
				fmt.Fprintln(w, (k*64+int(l))*2+3)
			}
		}
	}

	bar.Finish()
}
