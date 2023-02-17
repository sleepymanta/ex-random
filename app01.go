package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
)

func main() {
	// Random Byte
	sliceRand := make([]byte, 8)
	rand.Read(sliceRand)
	fmt.Printf("Byte Random: % x\n", sliceRand)
	// Byte Random: 72 04 3a 54 ca 55 e3 69

	// Random BigInt
	seedInt64, _ := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))

	// Math Random Value 1 - 100 and 0 - 99
	// mathRandom := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	mathRandom := mrand.New(mrand.NewSource(seedInt64.Int64()))
	fmt.Printf("Math Random: %d %d\n", mathRandom.Intn(101-1)+1, mathRandom.Intn(100))
	// Math Random: 45 91

	// DEPRECATED: Go 1.20
	// mrand.Seed(seedInt64.Int64())
	// fmt.Printf("Math Random: %d\n", mrand.Intn(101-1)+1)
}
