package pkgGoSocketSender

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FeedInput struct {
	key string
	url string
}

func (f *FeedInput) Initialize(key string, secret string) {
	f.key = key
	f.url = secret
}

func (f *FeedInput) SendMessage(body interface{}) (map[string]interface{}, error) {

	request, err := http.NewRequest("POST", f.url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("x-api-key", f.key)
	request.Header.Add("Content-Type", "application/json")

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}
	json.Unmarshal(responseBytes, &responseData)

	return responseData, nil
}
