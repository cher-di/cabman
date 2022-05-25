package edemrf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type CustomTime struct {
	Time time.Time
}

func (customTime *CustomTime) UnmarshalJSON(data []byte) error {
	stringTime := strings.Trim(string(data), `"`)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", stringTime)
	if err != nil {
		return fmt.Errorf("failed to parse custom time: %v", err)
	}
	customTime.Time = parsedTime
	return nil
}

type CustomUint32 struct {
	Uint32 uint32
}

func (customUint32 *CustomUint32) UnmarshalJSON(data []byte) error {
	stringUint32 := strings.Trim(string(data), `"`)
	parsedUint32, err := strconv.ParseUint(stringUint32, 10, 32)
	if err != nil {
		return fmt.Errorf("failed to parse custom uint32: %v", err)
	}
	customUint32.Uint32 = uint32(parsedUint32)
	return nil
}

type CustomFloat32 struct {
	Float32 float32
}

func (customFloat32 *CustomFloat32) UnmarshalJSON(data []byte) error {
	stringFloat32 := strings.Trim(string(data), `"`)
	parsedFloat32, err := strconv.ParseFloat(stringFloat32, 32)
	if err != nil {
		return fmt.Errorf("failed to parse custom time: %v", err)
	}
	customFloat32.Float32 = float32(parsedFloat32)
	return nil
}

type Route struct {
	Id         string       `json:"id"`
	UserId     string       `json:"userId"`
	CarId      string       `json:"carId"`
	FromCityId string       `json:"fromCityId"`
	ToCityId   string       `json:"toCityId"`
	StartTime  CustomTime   `json:"startTime"`
	EndTime    CustomTime   `json:"endTime"`
	Cost       CustomUint32 `json:"Cost"`
	FreePlaces CustomUint32 `json:"freePlaces"`
}

type User struct {
	Id               string        `json:"id"`
	Name             string        `json:"name"`
	Rating           CustomFloat32 `json:"rating"`
	ImageRelativeUrl string        `json:"image"`
	Thumbs           struct {
		Maxres string `json:"maxres"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"thumbs"`
}

type City struct {
	Id        string        `json:"id"`
	CountryId string        `json:"countryId"`
	RegionId  string        `json:"regionId"`
	Name      string        `json:"name"`
	Latitude  CustomFloat32 `json:"lat"`
	Longitude CustomFloat32 `json:"lng"`
}

type Routes struct {
	Success bool `json:"success"`
	Data    struct {
		Routes       []Route         `json:"routes"`
		RoutesUsers  map[string]User `json:"routesUsers"`
		RoutesCities map[string]City `json:"routesCities"`
	} `json:"data"`
	Meta struct {
		TotalCount CustomUint32 `json:"totalCount"`
		Page       CustomUint32 `json:"page"`
		PageSize   CustomUint32 `json:"pageSize"`
		PageCount  CustomUint32 `json:"pageCount"`
	} `json:"meta"`
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
