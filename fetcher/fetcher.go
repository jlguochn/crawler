package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get(url)
	resp,err := HttpGetFromProxy(url,"//127.0.0.1:1087")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func HttpGetFromProxy(urll, proxyURL string) (*http.Response, error) {
	reqest, _ := http.NewRequest("GET", urll, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy : http.ProxyURL(proxy),
		},
	}
	return client.Do(reqest)
}
