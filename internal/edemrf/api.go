package edemrf

import (
	"net/url"
	"path"
)

const API_SERVER_URL = "https://api.edemrf.com/v22"

func GetEndpoint(endpoint string) *url.URL {
	parsedUrl, _ := url.Parse(API_SERVER_URL)
	parsedUrl.Path = path.Join(parsedUrl.Path, endpoint)
	return parsedUrl
}
