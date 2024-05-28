package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationEncounters(pageURL *string) (RespEncounters, error) {
	url := baseURL + "/location-area/" + *pageURL

	if val, ok := c.cache.Get(url); ok {
		encResp := RespEncounters{}
		err := json.Unmarshal(val, &encResp)
		if err != nil {
			return RespEncounters{}, err
		}

		return encResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespEncounters{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespEncounters{}, err
	}

	encResp := RespEncounters{}
	err = json.Unmarshal(body, &encResp)
	if err != nil {
		return RespEncounters{}, err
	}

	c.cache.Add(url, body)
	return encResp, nil
}
