package rest

import (
	"encoding/json"
	"github.com/lenistwo/model"
	"github.com/lenistwo/util"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	apiURL          = "https://mobileapi.x-kom.pl/api/v1/xkom/hotShots/current?onlyHeader=true"
	apiKeyHeader    = "x-api-key"
	userAgentHeader = "user-agent"
)

var client http.Client

func New(timeout int64) {
	client = http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func RetrievePromotion() model.Promotion {
	request, err := http.NewRequest("GET", apiURL, nil)
	util.CheckError(err)

	request.Header.Set(apiKeyHeader, os.Getenv("API_KEY"))
	request.Header.Set(userAgentHeader, os.Getenv("CHROME_USER_AGENT"))

	response, err := client.Do(request)
	util.CheckError(err)

	defer response.Body.Close()
	obj, err := ioutil.ReadAll(response.Body)
	util.CheckError(err)

	return unmarshal(obj)
}

func unmarshal(obj []byte) model.Promotion {
	var p model.Promotion
	util.CheckError(json.Unmarshal(obj, &p))
	return p
}
