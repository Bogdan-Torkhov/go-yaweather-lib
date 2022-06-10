package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const yaKey = "47ec4e13-b304-4e7b-bd7c-0ca727d87fbc"

type key string

func GetWeather(apikey key, wid, let float64) (stc *Weather, err error) {
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=%f?lon=%f", wid, let)
	stc = new(Weather)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-Yandex-API-Key", yaKey)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode == 403 || resp.StatusCode == 404 {
		return stc, errors.New(fmt.Sprint(resp.StatusCode))
	}
	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(output, &stc)
	if err != nil {
		return
	}
	return
}
