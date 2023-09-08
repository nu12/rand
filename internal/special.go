package rand

func GenerateSpecial(i int) string {
	const pool = "!\"#$%&'()*+,-./:;<>?@[]{}^_|~"
	out := ""

	if i > 0 {
		out += GenerateSpecial(i - 1)
		n := getRandomInt(len(pool))
		out += string(pool[n])
	}

	return out
}
