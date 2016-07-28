package utils

import (
	"math/rand"
	"time"
)

var randBase = rand.New(rand.NewSource(time.Now().UnixNano()))

func EncodeInt32(i int32) []byte {
	code := []byte{
		byte(i),
		byte(i >> 8),
		byte(i >> 16),
		byte(i >> 24),
	}
	return code
}

func DecodeInt32(code []byte) int32 {
	var i int32 = 0
	for k, v := range code {
		if k > 3 {
			return i
		}
		i = i | (int32(v) << uint(k*8))
	}
	return i
}

func Rand() int {
	return randBase.Int()
}
