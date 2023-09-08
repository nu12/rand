package rand

func GenerateChar(i int, l, u bool) string {
	out := ""

	if i > 0 {
		out += GenerateChar(i-1, l, u)
		pool := providePool(l, u)
		n := getRandomInt(len(pool))
		out += string(pool[n])
	}

	return out
}

func providePool(l bool, u bool) string {
	const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	localPool := pool

	if l {
		localPool = pool[0:26]
	}

	if u {
		localPool = pool[26:52]
	}
	return localPool
}
