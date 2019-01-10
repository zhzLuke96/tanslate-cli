package transapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// tans API

type TransAPI interface {
	Trans(string) string
	TransEN(string) string
	TransCN(string) string
	TransTo(string, string) string
	Translate(string, string, string) string
}

func postData(target string, data url.Values) []byte {
	resp, err := http.PostForm(target, data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}
