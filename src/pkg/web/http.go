package web

import (
	"golang.org/x/net/html"
	"net/http"
)

var client = http.Client{}

func RequestNodes(url string) (*html.Node, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("user-agent", "Mozilla/5.0")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()
	return html.Parse(response.Body)
}
