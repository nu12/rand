package rand

func GenerateChar(i int, l, u bool) string {
	return GenerateCharFromPool(i, ProvideDefaultPool(l, u))
}

func GenerateCharFromPool(i int, pool string) string {
	out := ""

	if i > 0 {
		out += GenerateCharFromPool(i-1, pool)
		n := getRandomInt(len(pool))
		out += string(pool[n])
	}

	return out
}

func ProvideDefaultPool(l bool, u bool) string {
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

func ProvideUUIDPool() string {
	return "abcdef"
}
