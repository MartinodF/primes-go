package main

import (
  "flag"
  "fmt"
  "os"
  "path/filepath"
  "time"
)

func allocate(limit int64) []uint64 {
  bits := uint64((limit / 2) - 1)
  mod  := bits % 64
  ints := (bits - mod) / 64 + 1

  primes := make([]uint64, ints)

  if mod != 0 {
    // if the last bitmask isn't complete, set all remaining bits to 1
    primes[ints - 1] |= (0xFFFFFFFFFFFFFFFF << mod)
  }

  return primes
}

func sieve(limit uint64, primes []uint64) {
  fmt.Printf("Finding all primes <= %v\n", limit)

  start := time.Now()

  for k, i, run := uint64(0), uint64(3), true; run; k++ {
    for l := uint8(0); l < 64; l, i = l + 1, i + 2 {
      if (primes[k] >> l) & 1 == 1 {
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
        primes[p / 64] |= (1 << (p % 64))
      }
    }
  }

  fmt.Printf("Done in %v.\n", time.Since(start))
}

func save(primes []uint64) {
  fmt.Printf("\nSaving to primes.txt...\n");
  ints := uint(len(primes))
  start := time.Now()

  file, err := os.Create("primes.txt")
  if err != nil { panic(err) }
  defer file.Close()

  fmt.Fprintln(file, 2)
  for k := uint(0); k < ints; k++ {
    for l := uint(0); l < 64; l++ {
      if (primes[k] >> l) & 1 == 0 {
        fmt.Fprintln(file, (k * 64 + l) * 2 + 3)
      }
    }
  }

  fmt.Printf("Done in %v.\n", time.Since(start))
}

func main() {
  limit := flag.Int64("l", 100000, "Upper limit for prime searching")
  write := flag.Bool("s", false, "Save the prime numbers to primes.txt")

  flag.Usage = func() {
    fmt.Fprintln(os.Stderr, "A quick test implementing the Sieve of Eratostenes in Go")
    fmt.Fprintln(os.Stderr, "It searches for prime numbers up to the specified limit\n")
    fmt.Fprintf(os.Stderr, "Usage: %s [-l limit] [-s]\n\n", filepath.Base(os.Args[0]))
    flag.PrintDefaults()
    fmt.Fprint(os.Stderr, "\n")
  }

  flag.Parse()

  primes := allocate(*limit)
  sieve(uint64(*limit), primes)

  if *write {
    save(primes)
  }
}
