package collection

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// UploadFileHTTP uploads files from disk to a remote HTTP/S server
func UploadFileHTTP(uri, destPath string) error {

	f, err := os.Open(destPath)
	if err != nil {
		return err
	}
	fileReader := bufio.NewReader(f)

	// create HTTP client
	client := &http.Client{}

	// create HTTP GET request
	req, err := http.NewRequest("POST", uri, fileReader)
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
	fmt.Println(string(fileContents))
	return err
}
