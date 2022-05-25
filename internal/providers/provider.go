package providers

import "time"

type User struct {
	Name   string
	Rating float32
}

type City struct {
	Name string
}

type Route struct {
	FromCity   City
	ToCity     City
	StartTime  time.Time
	FinishTime time.Time
	Driver     User
	FreePlaces uint32
	Cost       uint32
}

type Provider interface {
	FindRoute() (*Route, error)
}
