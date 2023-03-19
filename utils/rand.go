package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandInt(n int) string {
	s := fmt.Sprintf("%12v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
	return s
}
