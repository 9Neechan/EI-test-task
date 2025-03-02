package util

import (
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

// Initializes the random number generator with the current time
func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

// Defines the alphabet for generating random strings
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Generates a random integer within the specified range
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random string of a specified length using the defined alphabet
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random name of 6 characters
func RandomName() string {
	return RandomString(6)
}

// Generates a random description of 15 characters
func RandomDescription() string {
	return RandomString(15)
}

// Generates a random amount of money between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
