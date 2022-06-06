package edemrf

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Route struct {
	Id         CustomUint32 `json:"id"`
	UserId     string       `json:"userId"`
	CarId      string       `json:"carId"`
	FromCityId string       `json:"fromCityId"`
	ToCityId   string       `json:"toCityId"`
	StartTime  CustomTime   `json:"startTime"`
	EndTime    CustomTime   `json:"endTime"`
	Cost       CustomUint32 `json:"Cost"`
	FreePlaces CustomUint32 `json:"freePlaces"`
}

type RouteUser struct {
	Id     CustomUint32  `json:"id"`
	Name   string        `json:"name"`
	Rating CustomFloat32 `json:"rating"`
	Thumbs struct {
		Maxres string `json:"maxres"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"thumbs"`
}

func (user *RouteUser) GetBestQualityThumbUrl() (string, error) {
	thumbsRelUrls := [...]string{user.Thumbs.Maxres, user.Thumbs.Large, user.Thumbs.Medium, user.Thumbs.Small}
	for _, thumbRelUrl := range thumbsRelUrls {
		if thumbRelUrl != "" {
			return getImageFullUrl(thumbRelUrl), nil
		}
	}
	return "", fmt.Errorf("no thumbs for user %v", user)
}

type RouteCity struct {
	Id        CustomUint32  `json:"id"`
	Name      string        `json:"name"`
	Latitude  CustomFloat32 `json:"lat"`
	Longitude CustomFloat32 `json:"lng"`
}

type Routes struct {
	reponseStatus
	Data struct {
		Routes []Route                    `json:"routes"`
		Users  map[CustomUint32]RouteUser `json:"routesUsers"`
		Cities map[CustomUint32]RouteCity `json:"routesCities"`
	} `json:"data"`
	Meta struct {
		TotalCount CustomUint32 `json:"totalCount"`
		Page       CustomUint32 `json:"page"`
		PageSize   CustomUint32 `json:"pageSize"`
		PageCount  CustomUint32 `json:"pageCount"`
	} `json:"meta"`
}

func GetRoutes(fromCityId uint32, toCityId uint32, createdDate time.Time, pageSize uint32, page uint32) (Routes, error) {
	parsedUrl := getEndpoint("/routes")
	query := url.Values{}
	query.Set("fromCityId", strconv.FormatUint(uint64(fromCityId), 32))
	query.Set("toCityId", strconv.FormatUint(uint64(toCityId), 32))
	query.Set("createdDate", createdDate.Format("2006-01-02"))
	query.Set("pageSize", strconv.FormatUint(uint64(pageSize), 32))
	query.Set("page", strconv.FormatUint(uint64(page), 32))
	parsedUrl.RawQuery = query.Encode()

	var routes Routes
	if err := sendGetRequest(parsedUrl.String(), &routes); err != nil {
		return Routes{}, fmt.Errorf("failed to get routes: %v", err)
	}
	return routes, nil
}
