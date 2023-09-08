package rand

import (
	"fmt"
)

func GenerateNum(i int) string {
	num := ""
	if i > 0 {
		num += GenerateNum(i - 1)
		num += fmt.Sprintf("%d", getRandomInt(10))
	}
	return num
}
