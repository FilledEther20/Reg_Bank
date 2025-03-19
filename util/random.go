// To create random testcases aligning with our requirements
package util

import (
	"math/rand"
	"time"
)

const (
	alphabets = "abcdefghijklmnopqrstuvwxyz"
)

var (
	currency  = []string{"USD", "EUR", "INR"}
	localRand = rand.New(rand.NewSource(time.Now().UnixNano())) // ✅ Local generator
)

// For generating a random number in a range
func RandomInt(min, max int64) int64 {
	return min + localRand.Int63n(max-min+1)
}

// For generating a random string of size n
func RandomString(n int) string {
	k := len(alphabets)
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = alphabets[localRand.Intn(k)]
	}
	return string(s)
}

// For generating a random owner
func RandomOwner(n int) string {
	return RandomString(n)
}

// For generating a random amount
func RandomBalance() int64 {
	return RandomInt(0, 500000)
}

// For generating random currency
func RandomCurrency() string {
	return currency[localRand.Intn(len(currency))]
}
