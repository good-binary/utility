package random

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomNumber generates a random integer between 'from' and 'to' (inclusive)
func RandomNumber(from, to int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(to-from+1) + from
}

// RandomString generates a random string of 'noOfChars' length, with a prefix and suffix
func RandomString(noOfChars int, prefix, suffix string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune(chars)
	s := make([]rune, noOfChars)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return prefix + string(s) + suffix
}

func RandomFullName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select a random first name and last name
	firstName := firstNames[r.Intn(len(firstNames))]
	lastName := lastNames[r.Intn(len(lastNames))]

	// return the random full name
	return fmt.Sprintf("%s %s\n", firstName, lastName)
}
