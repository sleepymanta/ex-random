package main

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"time"
)

func main() {
	// Random Byte
	sliceRand := make([]byte, 8)
	rand.Read(sliceRand)
	fmt.Printf("Byte Random: % x\n", sliceRand)
	// Byte Random: 72 04 3a 54 ca 55 e3 69

	// Random Value 1 - 100
	mrand.Seed(time.Now().UnixNano())
	fmt.Printf("Math Random: %d\n", mrand.Intn(101-1)+1)
	// Math Random: 45
}
