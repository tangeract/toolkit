package toolkit

import "crypto/rand"

const randomStringSource = "abcdefeghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

// Tools is the type used to instantiate this module. Any variable of this type will
// have access to all methods with receiver *Tools.
type Tools struct{}

// RandomString returns a string of random characters of length n, using randomStringSource as source.
func (t *Tools) RandomString(n int) string {
	// s is slice of Rune, of length - length
	// r is a slice of Rune from randomStringSource
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
