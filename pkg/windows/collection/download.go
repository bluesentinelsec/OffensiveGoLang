package collection

import (
	"io/ioutil"
	"net/http"
)

// DownloadFileHTTP downloads remote files over HTTP/S and saves the file to disk.
// use an absolute file path for destPath, for example, "C:\\Windows\\Temp\\file.txt"
func DownloadFileHTTP(uri, destPath string) error {

	// create HTTP client
	client := &http.Client{}

	// create HTTP GET request
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}

	// set user agent string
	userAgentString := `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18362`
	req.Header.Set("User-Agent", userAgentString)

	// execute HTTP GET request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// read HTTP response
	defer resp.Body.Close()
	fileContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// write download contents to a file
	err = ioutil.WriteFile(destPath, fileContents, 0)
	if err != nil {
		return err
	}
	return nil
}
