package rand

func GenerateAlpha(i int, l, u bool) string {
	return GenerateAlphaFromPool(i, ProvideDefaultPool(l, u))
}

func GenerateAlphaFromPool(i int, pool string) string {
	out := ""

	if i > 0 {
		out += GenerateAlphaFromPool(i-1, pool)
		p := getRandomInt(2)
		switch p {
		case 0:
			out += GenerateNum(1)
		case 1:
			out += GenerateCharFromPool(1, pool)
		}
	}
	return out
}
