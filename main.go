package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cher-di/cabman/internal/edemrf"
)

const (
	EkbCityId = "70079"
	PvkCityId = "70597"
	PageSize  = 10
)

func main() {
	fmt.Println("This is a cabman tool!")
	res, err := edemrf.GetRoutes(PvkCityId, EkbCityId, time.Now(), PageSize, 1)
	if err != nil {
		log.Fatalf("failed to make request: %v", err)
	}
	fmt.Printf("%v\n", res)
}
