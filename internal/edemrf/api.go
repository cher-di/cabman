package edemrf

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

const API_SERVER_URL = "https://api.edemrf.com/v23"
const WEB_SERVER_URL = "https://едем.рф"

func GetEndpoint(endpoint string) *url.URL {
	parsedUrl, _ := url.Parse(API_SERVER_URL)
	parsedUrl.Path = path.Join(parsedUrl.Path, endpoint)
	return parsedUrl
}

func GetImageFullUrl(relativeUrl string) string {
	return WEB_SERVER_URL + relativeUrl
}

func GetBestQualityThumbUrl(user User) (string, error) {
	thumbsRelUrls := [...]string{user.Thumbs.Maxres, user.Thumbs.Large, user.Thumbs.Medium, user.Thumbs.Small}
	for _, thumbRelUrl := range thumbsRelUrls {
		if thumbRelUrl != "" {
			return GetImageFullUrl(thumbRelUrl), nil
		}
	}
	return "", fmt.Errorf("no thumbs for user %v", user)
}

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
