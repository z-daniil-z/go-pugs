package generator

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
)

func Uuid() string {
	return ""
}

func RandomSHA256() string {
	r := rand.Uint64()
	str := strconv.FormatInt(int64(r), 10)
	tmp := sha256.Sum256([]byte(str))
	return string(tmp[:])
}
