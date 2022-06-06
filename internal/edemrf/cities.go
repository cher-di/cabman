package edemrf

import (
	"fmt"
	"net/url"
)

type CityItem struct {
	Id       CustomUint32 `json:"id"`
	Name     string       `json:"name"`
	Locality string       `json:"locality"`
	Address  string       `json:"address"`
	Priority CustomUint32 `json:"priority"`
}

type Cities struct {
	reponseStatus
	Data struct {
		Items []CityItem `json:"items"`
	} `json:"data"`
}

func GetCitiesSearchResutl(cityName string) (Cities, error) {
	parsedUrl := getEndpoint("/cities")
	query := url.Values{}
	query.Set("searchQuery", cityName)
	parsedUrl.RawQuery = query.Encode()

	var cities Cities
	if err := sendGetRequest(parsedUrl.String(), &cities); err != nil {
		return Cities{}, fmt.Errorf("failed to get citites with city name %s: %v", cityName, err)
	}
	return cities, nil
}
