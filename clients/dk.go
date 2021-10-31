package clients

import (
	"Sportsbook/dtos"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetDraftKingsPromotions() (*dtos.DraftKingsResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		"https://sportsbook-us-va.draftkings.com//sites/US-VA-SB/api/v4/eventgroups/88670561",
		nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil, err
	}


	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	target := dtos.DraftKingsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}