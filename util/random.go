package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(ownerLength int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < ownerLength; i++ {
		sb.WriteByte(alphabets[rand.Intn(k)])
	}

	return sb.String()
}

func RandomOwner() string {
	return randomString(10)
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomBalance() int64 {
	return randomInt(100, 10000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "NAIRA", "EUR"}
	k := len(currencies)
	return currencies[rand.Intn(k)]
}
