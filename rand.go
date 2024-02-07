package main

import rand "github.com/nu12/rand/internal"

// Alpha returns a string of random characters and numbers of a given length.
//
// length defines the length of the output string.
//
// lowerOnly and upperOnly define whether only lower case or upper case characters are present in the output.
func Alpha(lenght int, lowerOnly, upperOnly bool) string {
	return rand.GenerateAlpha(lenght, lowerOnly, upperOnly)
}

// Char returns a string of random characters of a given length.
//
// length defines the length of the output string.
//
// lowerOnly and upperOnly define whether only lower case or upper case characters are present in the output.
func Char(lenght int, lowerOnly, upperOnly bool) string {
	return rand.GenerateChar(lenght, lowerOnly, upperOnly)
}

// Num returns a string of random numbers of a given length.
//
// length defines the length of the output string.
func Num(lenght int) string {
	return rand.GenerateNum(lenght)
}

// Password returns a string of random characters, numbers and special signs of a given length.
//
// length defines the length of the output string.
//
// lowerOnly and upperOnly define whether only lower case or upper case characters are present in the output.
func Password(lenght int, lowerOnly, upperOnly bool) string {
	return rand.GeneratePassword(lenght, lowerOnly, upperOnly)
}

// Special returns a string of random special characters of a given length.
//
// length defines the length of the output string.
func Special(lenght int) string {
	return rand.GenerateSpecial(lenght)
}

// UUID returns an UUID
func UUID() string {
	return rand.GenerateUUID()
}
