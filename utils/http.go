package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func GetJSON(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		fmt.Println("Occurred an error: ", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
