package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type ResponseJSON struct {
	Page         int       `json:"page"`
	TotalResults int       `json:"total_results"`
	TotalPages   int       `json:"total_pages"`
	Results      []Results `json:"results"`
}

type Results struct {
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
	Overview   string `json:"overview"`
}

func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		fmt.Println("Deu Erro")
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	apiKey := flag.String("apikey", "foo", "a string")

	flag.Parse()

	x := RandomString(1)
	movieDbURL := "https://api.themoviedb.org/3/search/movie/?api_key=" + *apiKey + "&language=pt-BR&query=" + x
	resp := new(ResponseJSON)
	getJSON(movieDbURL, resp)
	numMovie := rand.Intn(resp.TotalResults)
	pageNumber := numMovie / 20
	numMovie = int(numMovie % 20)
	movieDbURL = movieDbURL + "&page=" + string(pageNumber)
	newResp := new(ResponseJSON)
	getJSON(movieDbURL, newResp)
	result := Results{}
	for index, v := range newResp.Results {
		if index == numMovie {
			result = v
			result.PosterPath = "https://image.tmdb.org/t/p/w500" + result.PosterPath
			break
		}
	}

	fmt.Printf("%+v\n", result)
}
