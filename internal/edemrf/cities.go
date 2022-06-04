package edemrf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CityInfo struct {
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Locality string       `json:"locality"`
	Address  string       `json:"address"`
	Priority CustomUint32 `json:"priority"`
}

type Cities struct {
	Success bool `json:"success"`
	Data    struct {
		Items []CityInfo `json:"items"`
	}
}

func GetCities(cityName string) (Cities, error) {
	parsedUrl := GetEndpoint("/cities")
	query := url.Values{}
	query.Set("searchQuery", cityName)
	parsedUrl.RawQuery = query.Encode()
	rawUrl := parsedUrl.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, rawUrl, nil)
	if err != nil {
		return Cities{}, fmt.Errorf("failed to make request to %s: %v", rawUrl, err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return Cities{}, fmt.Errorf("failed to send request to %s: %v", rawUrl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Cities{}, fmt.Errorf("request to %s finished with unexpected status code %s", rawUrl, resp.Status)
	}

	var cities Cities
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&cities); err != nil {
		return Cities{}, fmt.Errorf("failed to parse response from %s: %v", rawUrl, err)
	}
	return cities, nil
}
