package rand

import (
	"regexp"
	"testing"
)

var validLower = regexp.MustCompile(`[a-z]+`)
var validUpper = regexp.MustCompile(`[A-Z]+`)

func TestChar(t *testing.T) {

	i := 1
	for i < 10 {
		n := GenerateChar(i, false, false)
		if len(n) != i {
			t.Errorf("Error creating characters. Expected length of %d, got %s", i, n)
		}

		i++
	}
}

func TestLower(t *testing.T) {

	n := GenerateChar(99, true, false)

	match := validLower.Match([]byte(n))
	if !match {
		t.Errorf("Error creating characters. Expected all lower case, got %s", n)
	}
}

func TestUpper(t *testing.T) {

	n := GenerateChar(99, false, true)

	match := validUpper.Match([]byte(n))
	if !match {
		t.Errorf("Error creating characters. Expected all upper case, got %s", n)
	}
}

func TestDefaultPool(t *testing.T) {
	full := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	p := ProvideDefaultPool(false, false)
	if p != full {
		t.Errorf("Error creating pool of characters. Expected %s, got %s", full, p)
	}

	p = ProvideDefaultPool(true, false)
	if p != lower {
		t.Errorf("Error creating pool of characters. Expected %s, got %s", lower, p)
	}

	p = ProvideDefaultPool(false, true)
	if p != upper {
		t.Errorf("Error creating pool of characters. Expected %s, got %s", upper, p)
	}

	// Overwrites to upper
	p = ProvideDefaultPool(true, true)
	if p != upper {
		t.Errorf("Error creating pool of characters. Expected %s, got %s", upper, p)
	}
}

func TestUUIDPool(t *testing.T) {
	full := "abcdef"

	p := ProvideUUIDPool()
	if p != full {
		t.Errorf("Error creating pool of characters. Expected %s, got %s", full, p)
	}
}
