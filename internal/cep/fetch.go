package cep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://viacep.com.br/ws/"

func FetchCEPData(ctx context.Context, cep string) (*CEP, error) {

	url := fmt.Sprintf("%s%s/json/", baseUrl, cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var cepData CEP
	err = json.NewDecoder(res.Body).Decode(&cepData)
	if err != nil {
		return nil, err
	}

	return &cepData, nil
}
