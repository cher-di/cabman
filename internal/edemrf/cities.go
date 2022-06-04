package edemrf

import (
	"fmt"
	"net/url"
)

type City struct {
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Locality string       `json:"locality"`
	Address  string       `json:"address"`
	Priority CustomUint32 `json:"priority"`
}

type Cities struct {
	reponseStatus
	Data struct {
		Items []City `json:"items"`
	} `json:"data"`
}

func GetCities(cityName string) (Cities, error) {
	parsedUrl := GetEndpoint("/cities")
	query := url.Values{}
	query.Set("searchQuery", cityName)
	parsedUrl.RawQuery = query.Encode()

	var cities Cities
	if err := sendGetRequest(parsedUrl.String(), &cities); err != nil {
		return Cities{}, fmt.Errorf("failed to get citites with city name %s: %v", cityName, err)
	}
	return cities, nil
}
