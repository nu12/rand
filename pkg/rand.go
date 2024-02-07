package rand

import internal "github.com/nu12/rand/internal"

// Alpha returns a string of random characters and numbers of a given length.
//
// length defines the length of the output string.
//
// lowerOnly and upperOnly define whether only lower case or upper case characters are present in the output.
func Alpha(lenght int, lowerOnly, upperOnly bool) string {
	return internal.GenerateAlpha(lenght, lowerOnly, upperOnly)
}

// Char returns a string of random characters of a given length.
//
// length defines the length of the output string.
//
// lowerOnly and upperOnly define whether only lower case or upper case characters are present in the output.
func Char(lenght int, lowerOnly, upperOnly bool) string {
	return internal.GenerateChar(lenght, lowerOnly, upperOnly)
}

// Num returns a string of random numbers of a given length.
//
// length defines the length of the output string.
func Num(lenght int) string {
	return internal.GenerateNum(lenght)
}

// Password returns a string of random characters, numbers and special signs of a given length.
//
// length defines the length of the output string.
//
// lowerOnly and upperOnly define whether only lower case or upper case characters are present in the output.
func Password(lenght int, lowerOnly, upperOnly bool) string {
	return internal.GeneratePassword(lenght, lowerOnly, upperOnly)
}

// Special returns a string of random special characters of a given length.
//
// length defines the length of the output string.
func Special(lenght int) string {
	return internal.GenerateSpecial(lenght)
}

// UUID returns an UUID
func UUID() string {
	return internal.GenerateUUID()
}
