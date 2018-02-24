package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomMovieIndex(total, perPage int) (int, int) {
	numMovie := rand.Intn(total)
	pageNumber := numMovie / perPage
	numMovie = int(numMovie % perPage)
	return numMovie, pageNumber
}
