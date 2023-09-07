package rand

import (
	"regexp"
	"testing"
)

var validSpecial = regexp.MustCompile(`[!"#$%&'()*+,-./:;<>?@\[\]{}^_|~]+`)

func TestSpecial(t *testing.T) {

	i := 1
	for i < 10 {
		n := GenerateSpecial(i)
		if len(n) != i {
			t.Errorf("Error creating string. Expected length of %d, got %s", i, n)
		}
		match := validSpecial.Match([]byte(n))
		if !match {
			t.Errorf("Error creating string. Expected all special characters, got %s", n)
		}

		i++
	}

}
