package fireblocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	APIURL     string = "https://api.fireblocks.io"
	APIVERSION string = "v1"
)

var VaultId = "0"

func GetPath(path string) string {
	return fmt.Sprintf(`/%s%s`, APIVERSION, path)
}

func ClientGet(path string, v *url.Values) ([]byte, error) {
	query := ""
	path = GetPath(path)
	apiUrl := fmt.Sprintf(`%s%s`, APIURL, path)
	if v != nil {
		query = v.Encode()
		apiUrl = fmt.Sprintf(`%s%s?%s`, APIURL, path, query)
	}
	return doRequest(http.MethodGet, path, apiUrl, []byte(query))
}

func ClientPost(path string, v interface{}) ([]byte, error) {
	var result []byte
	path = GetPath(path)
	apiUrl := fmt.Sprintf(`%s%s`, APIURL, path)
	if v != nil {
		jsonData, err := json.Marshal(v)
		if err != nil {
			return result, err
		}
		return doRequest(http.MethodPost, path, apiUrl, jsonData)
	}
	return result, fmt.Errorf(`Sorry for post request must have json data`)
}

func ClientPut(path string, v interface{}) ([]byte, error) {
	var result []byte
	path = GetPath(path)
	apiUrl := fmt.Sprintf(`%s%s`, APIURL, path)
	if v != nil {
		jsonData, err := json.Marshal(v)
		if err != nil {
			return result, err
		}
		return doRequest(http.MethodPut, path, apiUrl, jsonData)
	}
	return result, fmt.Errorf(`Sorry for post request must have json data`)
}

func ClientDelt(path string) ([]byte, error) {
	path = GetPath(path)
	apiUrl := fmt.Sprintf(`%s%s`, APIURL, path)
	return doRequest(http.MethodDelete, path, apiUrl, []byte(``))
}

func doRequest(method string, path string, apiUrl string, body []byte) ([]byte, error) {
	var result []byte
	jwtToken, err := CreateJwt(path, body)
	if err != nil {
		fmt.Println("Error create jwtToken", err)
		return result, err
	}
	client := &http.Client{}
	var jsonData *bytes.Buffer
	if method == http.MethodGet {
		jsonData = bytes.NewBufferString("")
	} else {
		jsonData = bytes.NewBuffer(body)
	}
	req, _ := http.NewRequest(method, apiUrl, jsonData)
	req.Header.Add("X-API-Key", ApiKey)
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, jwtToken))
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error do request", err)
		return result, err
	}
	if response != nil {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error parse body result", err)
			return result, err
		}
		result = responseBody
	}
	return result, nil
}
