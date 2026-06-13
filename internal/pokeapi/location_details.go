package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationName string) (LocationAreaDetails, error) {
	url := baseURL + "/location-area/" + locationName + "/"
	if data, ok := c.cache.Get(url); ok {
		var locationDetails LocationAreaDetails
		err := json.Unmarshal(data, &locationDetails)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		return locationDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaDetails{}, fmt.Errorf("bad status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	c.cache.Add(url, data)

	var locationDetails LocationAreaDetails
	err = json.Unmarshal(data, &locationDetails)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	return locationDetails, nil
}