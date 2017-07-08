package system

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(postUrl string, postParam url.Values) (map[string]interface{}, error) {

	postValue := url.Values{}

	for key, value := range postParam {
		if key != "service" && key != "action" {
			postValue.Set(key, value[0])
		}
	}

	postUrl = "http://" + postUrl

	response, err := http.Post(postUrl, "application/x-www-form-urlencoded", strings.NewReader(postValue.Encode()))

	obj := make(map[string]interface{})
	if err != nil {
		Loger().Err(err.Error())
		return obj, nil
	}

	text, err := ioutil.ReadAll(response.Body)

	response.Body.Close()
	if err != nil {
		Loger().Err(err.Error())
		return obj, nil
	}

	err = json.Unmarshal(text, &obj)
	if err != nil {
		Loger().Err(err.Error())
		return obj, nil
	}

	return obj, nil

}
