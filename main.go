package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cher-di/cabman/internal/providers"
)

const (
	EkbCityId = "70079"
	PvkCityId = "70597"
	Alexander = "935984"
)

func main() {
	provider := providers.EdemrfProvider{
		FromCityId: PvkCityId, ToCityId: EkbCityId,
		StartTime: time.Now(), UserId: Alexander}
	route, err := provider.FindRoute()
	if err != nil {
		log.Fatalf("failed to find route: %v", err)
	}
	jsonRoute, _ := json.MarshalIndent(route, "", "  ")
	os.Stdout.Write(jsonRoute)
	fmt.Print("\n")
}
