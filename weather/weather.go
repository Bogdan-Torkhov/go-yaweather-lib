package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Key string

func GetWeather(apikey Key, wid, let float64, lang string) (stc *Weather, err error) {
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=%f&lon=%f&lang=%s", wid, let, lang)
	stc = new(Weather)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-Yandex-API-Key", string(apikey))
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
