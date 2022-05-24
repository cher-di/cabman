package edemrf

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Ruble uint32

func (ruble Ruble) String() string {
	return fmt.Sprintf("%d RUB", ruble)
}

type Route struct {
	Id         string    `json:"id"`
	UserId     string    `json:"userId"`
	CarId      string    `json:"carId"`
	FromCityId string    `json:"fromCityId"`
	ToCityId   string    `json:"toCityId"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Cost       Ruble     `json:"Cost"`
	FreePlaces uint8     `json:"freePlaces"`
}

type User struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	BirthDate        time.Time `json:"birthDate"`
	Rating           float32   `json:"rating"`
	imageRelativeUrl string    `json:"image"`
}

type City struct {
	Id        string  `json:"id"`
	CountryId string  `json:"countryId"`
	RegionId  string  `json:"regionId"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

type Routes struct {
	Success bool `json:"success"`
	Data    struct {
		Routes       []Route         `json:"routes"`
		RoutesUsers  map[string]User `json:"routesUsers"`
		RoutesCities map[string]City `json:"routesCities"`
	}
	Meta struct {
		TotalCount int `json:"totalCount"`
		Page       int `json:"page"`
		PageSize   int `json:"pageSize"`
		PageCount  int `json:"pageCount"`
	}
}

func GetRoutes(fromCityId string, toCityId string, createdDate time.Time, PageSize uint32, page uint32) (Routes, error) {
	parsedUrl := GetEndpoint("/routes")
	query := url.Values{}
	query.Set("fromCityId", fromCityId)
	query.Set("toCityId", toCityId)
	query.Set("createdDate", createdDate.Format("2006-01-02"))
	query.Set("pageSize", strconv.Itoa(int(PageSize)))
	query.Set("page", strconv.Itoa(int(page)))
	parsedUrl.RawQuery = query.Encode()
	rawUrl := parsedUrl.String()
	log.Printf("Result URL: %s", rawUrl)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, rawUrl, nil)
	if err != nil {
		return Routes{}, fmt.Errorf("failed to make request to %s: %v", rawUrl, err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return Routes{}, fmt.Errorf("failed to send request to %s: %v", rawUrl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Routes{}, fmt.Errorf("request to %s finished with status code %s", rawUrl, resp.Status)
	}

	var routes Routes
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&routes); err != nil {
		return Routes{}, fmt.Errorf("failed to parse response from %s: %v", rawUrl, err)
	}
	return routes, nil
}
