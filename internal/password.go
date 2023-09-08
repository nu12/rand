package rand

func GeneratePassword(i int, l, u bool) string {
	out := ""

	if i > 0 {
		out += GeneratePassword(i-1, l, u)
		p := getRandomInt(3)
		switch p {
		case 0:
			out += GenerateNum(1)
		case 1:
			out += GenerateChar(1, l, u)
		case 2:
			out += GenerateSpecial(1)
		}
	}
	return out
}
