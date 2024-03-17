package util

import "math/rand"

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
