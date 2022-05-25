package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/cher-di/cabman/internal/edemrf"
)

type EdemrfProvider struct {
	FromCityId string
	ToCityId   string
	StartTime  time.Time
	UserId     string
}

func IsDatesEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (provider *EdemrfProvider) IsRightRoute(route *edemrf.Route) bool {
	return route.FromCityId == provider.FromCityId &&
		route.ToCityId == provider.ToCityId &&
		IsDatesEqual(route.StartTime.Time, provider.StartTime) &&
		route.UserId == provider.UserId
}

func MakeProviderRoute(data *edemrf.Routes, route *edemrf.Route) Route {
	eDriver := data.Data.Users[route.UserId]
	eFromCity := data.Data.Cities[route.FromCityId]
	eToCity := data.Data.Cities[route.ToCityId]

	return Route{
		FromCity:   City{Name: eFromCity.Name},
		ToCity:     City{Name: eToCity.Name},
		StartTime:  route.StartTime.Time,
		FinishTime: route.EndTime.Time,
		Driver: User{
			Name:   eDriver.Name,
			Rating: eDriver.Rating.Float32,
		},
		FreePlaces: route.FreePlaces.Uint32,
		Cost:       route.Cost.Uint32,
	}
}

func (provider *EdemrfProvider) FindRoute() (Route, error) {
	pageCount := uint32(1)
	pageSize := uint32(10)
	for page := uint32(1); page <= pageCount; page++ {
		data, err := edemrf.GetRoutes(provider.FromCityId, provider.ToCityId,
			provider.StartTime, pageSize, page)
		if err != nil {
			return Route{}, fmt.Errorf("failed to fetch data to search route: %v", err)
		}
		pageCount = data.Meta.PageCount.Uint32
		for _, route := range data.Data.Routes {
			if provider.IsRightRoute(&route) {
				return MakeProviderRoute(&data, &route), nil
			}
		}
	}
	return Route{}, errors.New("not found appropriate route")
}
