// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)

//!+

package main

import (
	"fmt"
	// "io"
	// "io/ioutil"
	// "net/http"
	// "os"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc { //for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1) //末尾值+右移一位表中值
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 for loop implement
func PopCount2(x uint64) int {
	var res int
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>byte((i*8)))])
	}
	return res
}

// PopCount3 while
func PopCount3(x uint64) int {
	var res int
	for x > 0 {
		res += int(x & 1)
		x >>= 1
	}
	return res
}

// PopCount4 while
func PopCount4(x uint64) int {
	var res int
	for x > 0 {
		res++
		x = x & (x - 1)
	}
	return res
}

func main() {
	var count = 100000000
	start := time.Now()
	for index := 0; index < count; index++ {
		// fmt.Println("pop1:",PopCount(1123))
		PopCount(1123)
	}
	fmt.Println("pop1:", PopCount(1123))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	for index := 0; index < count; index++ {
		PopCount2(1123)
	}
	fmt.Println("pop2:", PopCount2(1123))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	for index := 0; index < count; index++ {
		PopCount3(1123)
	}
	fmt.Println("pop3:", PopCount3(1123))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	for index := 0; index < count; index++ {
		PopCount4(1123)
	}
	fmt.Println("pop4:", PopCount4(1123))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//!-
