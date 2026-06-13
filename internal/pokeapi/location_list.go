package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl string) (LocationArea, error) {
	url := baseURL + "/location-area/"
	if pageUrl != "" {
		url = pageUrl
	}

	if data, ok := c.cache.Get(url); ok {
		var locationArea LocationArea
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		return LocationArea{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)

	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
