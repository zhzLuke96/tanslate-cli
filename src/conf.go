package main

import (
	"encoding/json"
	"io/ioutil"
)

type rootConfig struct {
	DefaultUse string      `json:"default"`
	Baidu      baiduOption `json:"baidu"`
	Bing       bingOption  `json:"bing"`
}

type baiduOption struct {
	URL   string `json:"url"`
	AppID string `json:"appid"`
	Key   string `json:"key"`
}

type bingOption struct {
	URL   string `json:"url"`
	AppID string `json:"appid"`
	Key   string `json:"key"`
}

func loadConf(filename string) rootConfig {
	ret := rootConfig{}
	data, _ := ioutil.ReadFile(filename)

	json.Unmarshal(data, &ret)
	return ret
}
