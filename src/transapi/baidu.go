package transapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type transRes struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}

type baiduRespon struct {
	From   string     `json:"from"`
	To     string     `json:"to"`
	Result []transRes `json:"trans_result"`
}

func getMd5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	ret := m.Sum(nil)
	return hex.EncodeToString(ret)
}

func getSign(appid, query, salt, key string) string {
	return getMd5(appid + query + salt + key)
}

type baiduTranslate struct {
	apiURL string
	appID  string
	key    string
}

func NewBaiduTrans(url, id, key string) *baiduTranslate {
	return &baiduTranslate{url, id, key}
}

func (btl *baiduTranslate) baseTrans(from, to, query string) string {
	if strings.Trim(query, "\n\t\r ") == "" {
		return ""
	}

	rand.Seed(time.Now().Unix())

	salt := strconv.Itoa(rand.Intn(10000000))

	sign := getSign(btl.appID, query, salt, btl.key)

	res := postData(btl.apiURL, url.Values{"from": {from}, "to": {to}, "q": {query}, "appid": {btl.appID}, "salt": {salt}, "sign": {sign}})

	resp := baiduRespon{}

	json.Unmarshal(res, &resp)

	ret := ""

	for _, val := range resp.Result {
		ret += val.Dst
	}

	return strings.Trim(ret, "\n\t\r ")
}

func (btl *baiduTranslate) Trans(text string) string {
	return btl.baseTrans("auto", "zh", text)
}

func (btl *baiduTranslate) TransCN(text string) string {
	return btl.Trans(text)
}

func (btl *baiduTranslate) TransEN(text string) string {
	return btl.baseTrans("auto", "en", text)
}

func (btl *baiduTranslate) TransTo(to, text string) string {
	return btl.baseTrans("auto", to, text)
}

func (btl *baiduTranslate) Translate(from, to, text string) string {
	return btl.baseTrans(from, to, text)
}
