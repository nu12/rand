package rand

import (
	"regexp"
	"testing"
)

var validNum = regexp.MustCompile(`\d+`)

func TestNum(t *testing.T) {

	i := 1
	for i < 10 {
		n := GenerateNum(i)
		if len(n) != i {
			t.Errorf("Error creating number. Expected length of %d, got %s", i, n)
		}

		match := validNum.Match([]byte(n))
		if !match {
			t.Errorf("Error creating number. Got %s", n)
		}

		i++
	}
}
