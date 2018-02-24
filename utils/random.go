package utils

import (
	"math/rand"
	"time"
)

// RandomString generate a string based on n argument
func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// RandomMovieIndex generate a index of movie inside a pageNumber
func RandomMovieIndex(total, perPage int) (int, int) {
	numMovie := rand.Intn(total)
	pageNumber := numMovie / perPage
	numMovie = int(numMovie % perPage)
	return numMovie, pageNumber
}
