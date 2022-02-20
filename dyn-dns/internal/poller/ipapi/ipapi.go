package ipapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/larwef/my-platform/dyn-dns/internal/poller"
)

const ipAPIURL = "http://ip-api.com/json/"

type IPApi struct {
	client *http.Client
	url    string
}

func New(client *http.Client) *IPApi {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	return &IPApi{
		client: client,
		url:    ipAPIURL,
	}
}

func (i *IPApi) Poll() (net.IP, error) {
	res, err := i.client.Get(i.url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		payload, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("poller: unable to read payload: %v\n", err)
			payload = []byte{}
		}
		return nil, &poller.Error{Code: res.StatusCode, Message: string(payload)}
	}
	respObj := struct {
		Query string `json:"query"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&respObj); err != nil {
		return nil, err
	}
	return net.ParseIP(respObj.Query), nil
}
