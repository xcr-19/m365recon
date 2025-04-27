package pkg

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/xcr-19/m365recon/utils"
)

func SetupHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
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
