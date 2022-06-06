package edemrf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

const API_SERVER_URL = "https://api.edemrf.com/v23"
const WEB_SERVER_URL = "https://едем.рф"

func getEndpoint(endpoint string) *url.URL {
	parsedUrl, _ := url.Parse(API_SERVER_URL)
	parsedUrl.Path = path.Join(parsedUrl.Path, endpoint)
	return parsedUrl
}

func getImageFullUrl(relativeUrl string) string {
	return WEB_SERVER_URL + relativeUrl
}

type reponseStatus struct {
	Success bool `json:"success"`
}

func sendGetRequest(url string, resultStorage interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to make request to %s: %v", url, err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to %s: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request to %s finished with unexpected status code %s", url, resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from %s: %v", url, err)
	}

	var status reponseStatus
	if err := json.Unmarshal(body, &status); err != nil {
		return fmt.Errorf("failed to parse response status from url %s: %v", url, err)
	}
	if !status.Success {
		return fmt.Errorf("request to %s was not successfull", url)
	}
	if err := json.Unmarshal(body, resultStorage); err != nil {
		return fmt.Errorf("failed to parse response from %s: %v", url, err)
	}
	return nil
}

type WithThumbUrls interface {
	GetBestQualityThumbUrl() (string, error)
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
