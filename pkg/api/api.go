package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Url string
}

func NewApi(url string) *Api {
	return &Api{Url: url}
}

func (a *Api) MakeRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d %s", response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (a *Api) Get(path string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, a.Url+path, nil)
	if err != nil {
		return nil, err
	}
	return a.MakeRequest(request)
}

func (a *Api) Post(path string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, a.Url+path, body)
	if err != nil {
		return nil, err
	}
	return a.MakeRequest(request)
}
