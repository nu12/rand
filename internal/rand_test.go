package rand

import (
	"testing"
)

func TestGetRandomInt(t *testing.T) {
	i := 1
	for i < 100 {
		n := getRandomInt(i)
		if n >= i {
			t.Errorf("Error generating random number. It should not be greater or equal to %d, got %d", i, n)
		}
		i++
	}
}
