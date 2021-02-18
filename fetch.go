package rest

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Fetch data from a server and Unmarshal data into v.
// defaults to JSON if the content type is not specified.
func Fetch(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.Header.Get(HeaderContentType) {
	case ContentXML:
		if err = xml.Unmarshal(body, v); err != nil {
			return err
		}
	case ContentYAML:
		if err = yaml.Unmarshal(body, v); err != nil {
			return err
		}
	default:
		if err = json.Unmarshal(body, v); err != nil {
			return err
		}
	}
	return nil
}
