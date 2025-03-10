package rand

import (
	"testing"
)

func TestPassword(t *testing.T) {

	i := 1
	for i < 10 {
		n := GeneratePassword(i, false, false)
		if len(n) != i {
			t.Errorf("Error creating string. Expected length of %d, got %s", i, n)
		}

		i++
	}
}
