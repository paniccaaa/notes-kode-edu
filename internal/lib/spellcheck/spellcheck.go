package spellcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SpellerResponse struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func CheckTexts(texts []string) ([][]SpellerResponse, error) {
	const op = "lib.spellcheck.CheckTexts"
	apiURL := "https://speller.yandex.net/services/spellservice.json/checkTexts"

	data := url.Values{}
	for _, text := range texts {
		data.Add("text", text)
	}

	url := fmt.Sprintf("%s?%s", apiURL, data.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	var result [][]SpellerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}
