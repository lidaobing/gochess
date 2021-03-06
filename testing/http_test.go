package testing

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/jonpchin/gochess/gostuff"
)

func TestLocale(t *testing.T) {

	client := gostuff.TimeOutHttp(5)
	response, err := client.Get("http://freegeoip.net/json/77.124.0.0")
	if response == nil {
		t.Error("URL time out for http://freegeoip.net/json/77.124.0.0 in TestLocale")
	}
	if err != nil {
		t.Error("Failed TestLocale http get")
	}
	defer response.Body.Close()
	htmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("Failed TestLocale read body")
	}

	var ipLocation gostuff.IPLocation

	if err := json.Unmarshal(htmlData, &ipLocation); err != nil {
		t.Error("Failed in JSON unmarshal", string(htmlData), err)
	}

	if ipLocation.Country_name != "Israel" {
		t.Error("Failed country name")
	}
}
