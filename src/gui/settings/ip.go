package settings

import (
	"encoding/json"
	"io"
	"net/http"
)

const apiUrl = "https://ipinfo.io/json"

func getMyRegion() (string, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	as := &struct {
		Timezone string `json:"timezone"`
	}{}
	if err = json.Unmarshal(all, as); err != nil {
		return "", err
	}
	b := make([]rune, 0, len(as.Timezone))
	for _, char := range as.Timezone {
		if char == '/' {
			break
		}
		b = append(b, char)
	}
	return string(b), nil
}
