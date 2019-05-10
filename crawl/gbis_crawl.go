package crawl

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get333() string {
	urlStr := "http://www.gbis.go.kr/gbis2014/schBusAPI.action"

	resp, err := http.PostForm(urlStr, url.Values{
		"cmd":     {"searchRouteJson"},
		"routeId": {"205000002"}, // 333 번 버스 id
	})

	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
		return str
	}
	return ""
}
