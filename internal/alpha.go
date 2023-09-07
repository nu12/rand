package rand

import (
	"math/rand"
	"time"
)

func GenerateAlpha(i int, l, u bool) string {
	out := ""

	if i > 0 {
		out += GenerateAlpha(i-1, l, u)

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		p := r.Intn(2)
		switch p {
		case 0:
			out += GenerateNum(1)
		case 1:
			out += GenerateChar(1, l, u)
		}
	}

	return out
}
