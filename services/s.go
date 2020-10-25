package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetRequest() interface{} {

	endpoint := `http://httpbin.org/get`

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var m interface{}

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println(err)
	}

	return m

}

func PostWithHeader() interface{} {

	endpoint := `http://httpbin.org/post`

	reqBody, err := json.Marshal(map[string]string{
		"name":  "username",
		"email": "email@domain.com",
	})

	if err != nil {
		log.Println(err)
	}

	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var m interface{}

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println(err)
	}

	return m
}

func CustomRequest() interface{} {

	endpoint := `http://httpbin.org/post`

	reqBody, err := json.Marshal(map[string]string{
		"name":  "username",
		"email": "email@domain.com",
	})

	client := &http.Client{}

	rq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
	}

	rq.Header.Add("Content-Type", "application/json")
	rq.Header.Add("x-api-key", "your api key")

	res, err := client.Do(rq) //make request
	if err != nil {

		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		log.Fatal(err)
	}
	// log.Println(string(body))

	var m interface{}

	e := json.Unmarshal(body, &m)

	if e != nil {
		log.Fatal(e)
	}

	return m

}

func UrlEncode() interface{} {
	tenant := `tenant_id`
	clientID := `client_id`
	clientSecret := `client_secret`
	scope := `https://graph.microsoft.com/.default mail info`
	grantType := `client_credentials`

	// endpoint := `https://login.microsoftonline.com/` + tenant + `/oauth2/v2.0/token`
	endpoint := `http://httpbin.org/post`

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("scope", scope)
	data.Set("grant_type", grantType)
	data.Set("tenant", tenant)

	client := &http.Client{}
	rq, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}

	rq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rq.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(rq) //request
	if err != nil {

		log.Fatal(err)
	}

	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		log.Fatal(err)
	}
	log.Println(string(body))

	var m interface{}

	err = json.Unmarshal(body, &m)

	if err != nil {
		log.Println(err)

	}

	return m

}
