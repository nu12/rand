package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNum(i int) string {
	num := ""
	if i > 0 {
		num += GenerateNum(i - 1)
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		n := r.Intn(10)
		num += fmt.Sprintf("%d", n)
	}
	return num
}
