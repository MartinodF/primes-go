Primes.go
====

### Introduction

A quick test implementing the Sieve of Eratostenes in Go.

It searches for prime numbers up to the specified limit, optionally saving them to file.

It is one of many small programs I wrote to get a feel for [#golang](http://golang.org/). If you're interested in generating prime numbers, you should probably look into [primegen](http://cr.yp.to/primegen.html) or some other, more suitable alternative.

### Compiling

    $ go build primes.go

### Usage

primes [-l limit] [-s]

  * -l limit *(default 100000)* Upper limit for prime searching
  * -s *(default false)* Save the prime numbers to primes.txt

### Speed

These are some non-accurate benchmarks I ran on my machine.

**Specs**

* AMD Phenom II X6 1055T @3.4GHz
* 8GB of cheap Kingston DDR3 @486MHz
* Windows 8 RTM 64-bit

**Results**

Ram usage and CPU time for finding primes up to:

           1'000'000    <1MB 2ms
          10'000'000    <1MB 26ms
         100'000'000   7.4MB 401ms
       1'000'000'000  61.2MB 9.01s
      10'000'000'000 599.8MB 1m51s

### License

Copyright 2012 Martino di Filippo

Licensed under the MIT License
