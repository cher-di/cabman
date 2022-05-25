package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DownloadFile(url string, writer io.Writer) error {
	req, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make get request to %s: %v", url, err)
	}
	defer req.Body.Close()
	bytes_wrote, err := io.Copy(writer, req.Body)
	log.Printf("Downloaded %d bytes from %s", bytes_wrote, url)
	if err != nil {
		return fmt.Errorf("failed to download all data from %s", url)
	}
	return nil
}
