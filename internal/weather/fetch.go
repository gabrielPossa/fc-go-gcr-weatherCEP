package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const baseUrl = "https://api.weatherapi.com/v1/current.json"

var key = os.Getenv("API_KEY")

func GetWeatherData(ctx context.Context, q string) (*Weather, error) {

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	cleanQ, _, err := transform.String(t, q)
	if err != nil {
		return nil, err // Handle errors as appropriate for your application
	}

	params := url.Values{}
	params.Add("q", cleanQ)
	params.Add("key", key)

	finalURL := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", finalURL, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var weatherData Weather
	err = json.NewDecoder(res.Body).Decode(&weatherData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &weatherData, nil
}
