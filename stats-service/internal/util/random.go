package util

import (
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomDescription() string {
	return RandomString(15)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
