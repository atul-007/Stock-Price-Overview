package data

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile downloads a file from a URL and saves it to the specified path
func DownloadFile(filePath, url string) error {
	fmt.Printf("Downloading file from %s to %s\n", url, filePath)

	// Create a new HTTP client with a custom user-agent header
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow redirects
			return nil
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set a user-agent header (replace "YourUserAgent" with an appropriate user agent)
	req.Header.Set("User-Agent", "YourUserAgent")

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file. Status code: %d", response.StatusCode)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("File downloaded: %s\n", filePath)
	return nil
}
