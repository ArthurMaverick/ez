package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Continent struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"continent"`
	Country         string `json:"country"`
	Loc             string `json:"loc"`
	Anycast         bool   `json:"anycast"`
	City            string `json:"city"`
	Org             string `json:"org"`
	Timezone        string `json:"timezone"`
	IP              string `json:"ip"`
	CountryCurrency struct {
		Symbol string `json:"symbol"`
		Code   string `json:"code"`
	} `json:"country_currency"`
	Hostname    string `json:"hostname"`
	CountryName string `json:"country_name"`
	Postal      string `json:"postal"`
	Region      string `json:"region"`
	CountryFlag struct {
		Emoji   string `json:"emoji"`
		Unicode string `json:"unicode"`
	} `json:"country_flag"`
}

type Post struct {
	IpAdress string `json:"ip_address"`
}

func GetIpInfo(ip string) (Response, error) {
	var R Response
	endpoint := "https://bpcevj2u7xdw6mqrkgdos7ns5i0nxbwo.lambda-url.us-east-1.on.aws"
	postBody, _ := json.Marshal(map[string]string{
		"ip_address": ip,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(endpoint, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	json_err := json.Unmarshal([]byte(sb), &R)

	if json_err != nil {
		log.Fatalln(json_err)
	}

	fmt.Printf(`Continent: 
  Code: %s 
  Name: %s
Contry: %s
Loc: %s
Anycast: %t
City: %s
Org: %s
Timezone: %s
Ip: %s
CountryCurrency:
  Symbol: %s
  Code: %s
Hostname: %s
CountryName: %s
Postal: %s
Region: %s
CountryFlag:
  Emoji: %s
  Unicode: %s
`, R.Continent.Code, R.Continent.Name, R.Country, R.Loc, R.Anycast, R.City, R.Org, R.Timezone, R.IP, R.CountryCurrency.Symbol, R.CountryCurrency.Code, R.Hostname, R.CountryName, R.Postal, R.Region, R.CountryFlag.Emoji, R.CountryFlag.Unicode)

	return R, nil
}
