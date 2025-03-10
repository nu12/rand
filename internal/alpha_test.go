package rand

import (
	"regexp"
	"testing"
)

var validAlpha = regexp.MustCompile(`[a-zA-Z\d]+`)
var validAlphaLower = regexp.MustCompile(`[a-z\d]+`)
var validAlphaUpper = regexp.MustCompile(`[A-Z\d]+`)

func TestAlpha(t *testing.T) {

	i := 1
	for i < 10 {
		n := GenerateAlpha(i, false, false)
		if len(n) != i {
			t.Errorf("Error creating string. Expected length of %d, got %s", i, n)
		}
		match := validAlpha.Match([]byte(n))
		if !match {
			t.Errorf("Error creating string. Expected alphanumeric characters, got %s", n)
		}

		i++
	}
}

func TestAlphaLower(t *testing.T) {

	n := GenerateAlpha(99, true, false)

	match := validAlphaLower.Match([]byte(n))
	if !match {
		t.Errorf("Error creating characters. Expected all lower case, got %s", n)
	}
}

func TestAlphaUpper(t *testing.T) {

	n := GenerateAlpha(99, false, true)

	match := validAlphaUpper.Match([]byte(n))
	if !match {
		t.Errorf("Error creating characters. Expected all upper case, got %s", n)
	}
}
