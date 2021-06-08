package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_API_Login(t *testing.T) {

	var (
		uname  = "quangtue"
		pwd    = "1234567"
		method = http.MethodGet
	)

	req, err := http.NewRequest(method, path["login"], nil)
	if err != nil {
		panic(err)
	}

	params := req.URL.Query()
	params.Add("username", uname)
	params.Add("password", pwd)
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	client.CloseIdleConnections()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---response: ", string(body))

	if res.StatusCode != 200 {
		t.Errorf("Test Login Failed")
	}

	defer res.Body.Close()
}
