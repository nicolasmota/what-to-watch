package main

import (
	"flag"
	"fmt"
	"github.com/nicolasmota/what-to-watch/utils"
	"os"
)

// Results represents the results returned inside from ResponseJSON
type Results struct {
	VoteCount        int     `json:"vote_count"`
	ID               int     `json:"id"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	Title            string  `json:"title"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIds         []int   `json:"genre_ids"`
	BackdropPath     string  `json:"backdrop_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}

// ResponseJSON represents a response returned from themoviedb
type ResponseJSON struct {
	Page         int       `json:"page"`
	TotalResults int       `json:"total_results"`
	TotalPages   int       `json:"total_pages"`
	Results      []Results `json:"results"`
}

func main() {
	apiKey := flag.String("apikey", "", "")
	flag.Parse()

	if !(len(*apiKey) > 0) {
		fmt.Println("You need to add apikey as argument. ie: --apikey=xxxkkkwww000332")
		os.Exit(1)
	}

	x := utils.RandomString(1)

	movieDbURL := fmt.Sprintf("https://api.themoviedb.org/3/search/movie/?api_key=%s&language=pt-BR&query=%s", *apiKey, x)

	resp := new(ResponseJSON)
	utils.GetJSON(movieDbURL, resp)

	numMovie, pageNumber := utils.RandomMovieIndex(resp.TotalResults, len(resp.Results))
	movieDbURL = movieDbURL + "&page=" + string(pageNumber)

	newResp := new(ResponseJSON)
	utils.GetJSON(movieDbURL, newResp)

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
