package edemrf

import (
	"net/url"
	"path"
)

const API_SERVER_URL = "https://api.edemrf.com/v22"
const WEB_SERVER_URL = "https://едем.рф"

func GetEndpoint(endpoint string) *url.URL {
	parsedUrl, _ := url.Parse(API_SERVER_URL)
	parsedUrl.Path = path.Join(parsedUrl.Path, endpoint)
	return parsedUrl
}

func GetImageFullUrl(relativeUrl string) string {
	return WEB_SERVER_URL + relativeUrl
}
