package main

import (
  "flag"
  "fmt"
  "io"
  "os"
  "path/filepath"
  "time"
)

func allocate(limit uint64) []uint64 {
  bits := (limit - 1) / 2

  if bits < 0 || limit <= 2 {
    bits = 0
  }

  mod  := bits % 64
  ints := (bits - mod) / 64

  if mod != 0 {
    ints++
  }

  primes := make([]uint64, ints)

  if mod != 0 && ints > 0 {
    // if the last bitmask isn't complete, set all remaining bits to 1
    primes[ints - 1] |= (0xFFFFFFFFFFFFFFFF << mod)
  }

  return primes
}

func sieve(limit uint64, primes []uint64) {
  if limit < 3 {
    return
  }

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
}

func saveTo(w io.Writer, primes []uint64) {
  for k, p := range primes {
    for l := uint8(0); l < 64; l++ {
      if (p >> l) & 1 == 0 {
        fmt.Fprintln(w, (k * 64 + int(l)) * 2 + 3)
      }
    }
  }
}

func main() {
  limit := flag.Uint64("l", 100000, "Upper limit for prime searching")
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

  fmt.Printf("Finding all primes <= %v\n", *limit)
  start := time.Now()

  sieve(*limit, primes)

  fmt.Printf("Done in %v.\n", time.Since(start))

  if *write {
    file, err := os.Create("primes.txt")
    if err != nil { panic(err) }
    defer file.Close()

    fmt.Printf("\nSaving to primes.txt...\n");
    start := time.Now()

    if *limit >= 2 {
      fmt.Fprintln(file, 2)
    }

    saveTo(file, primes)

    fmt.Printf("Done in %v.\n", time.Since(start))
  }
}
