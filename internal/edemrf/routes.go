package edemrf

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

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

type RUser struct {
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

type RCity struct {
	Id        string        `json:"id"`
	CountryId string        `json:"countryId"`
	RegionId  string        `json:"regionId"`
	Name      string        `json:"name"`
	Latitude  CustomFloat32 `json:"lat"`
	Longitude CustomFloat32 `json:"lng"`
}

type Routes struct {
	reponseStatus
	Data struct {
		Routes []Route          `json:"routes"`
		Users  map[string]RUser `json:"routesUsers"`
		Cities map[string]RCity `json:"routesCities"`
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

	var routes Routes
	if err := sendGetRequest(parsedUrl.String(), &routes); err != nil {
		return Routes{}, fmt.Errorf("failed to get routes: %v", err)
	}
	return routes, nil
}
