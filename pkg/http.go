package pkg

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/xcr-19/m365recon/utils"
)

func SetupHTTPClient(proxy string) *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		fmt.Println("Proxy: ", proxyURL)
		if err != nil {
			fmt.Println("Unable to parse Proxy: ", err)
			os.Exit(1)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	return client
}
func GetRequest(url string, method string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", utils.UserAgent[rand.Intn(len(utils.UserAgent))])
	return request, nil
}
