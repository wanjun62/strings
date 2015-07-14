package strings

import (
	"math/rand"
	"strings"
)

// GenerateRandomString generates a random string with specified length
func GenerateRandomString(r rand.Rand, length int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
	str := make([]string, length)
	for i := 0; i < length; i++ {
		index := r.Intn(len(chars))
		str[i] = chars[index : index+1]
	}
	return strings.Join(str, "")
}
