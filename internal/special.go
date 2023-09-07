package rand

import (
	"math/rand"
	"time"
)

func GenerateSpecial(i int) string {
	const pool = "!\"#$%&'()*+,-./:;<>?@[]{}^_|~"
	out := ""

	if i > 0 {
		out += GenerateSpecial(i - 1)

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		out += string(pool[r.Intn(len(pool))])
	}

	return out
}
