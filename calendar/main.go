package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type day struct {
	Suit  string `json:"suit"`
	Avoid string `json:"avoid"`
}

type data struct {
	Almanac []*day `json:"almanac"`
}

type result struct {
	Data []*data `json:"data"`
}

func main() {
	suit := make(map[string]int)
	avoid := make(map[string]int)
	url := "https://sp0.baidu.com/8aQDcjqpAAV3otqbppnN2DJv/api.php?query=%v年%v月&resource_id=6018&format=json"

	for year := 2019; year <= 2020; year++ {
		for month := 1; month <= 12; month++ {
			body, err := httpSend(fmt.Sprintf(url, year, month))
			if err != nil {
				panic(err)
			}

			model := result{}
			err = json.Unmarshal(body, &model)
			if err != nil {
				panic(err)
			}
			for _, v := range model.Data[0].Almanac {
				tmpSuits := strings.Split(v.Suit, ".")
				for _, item := range tmpSuits {
					if item != "" {
						suit[item]++
					}
				}
				tmpAvoid := strings.Split(v.Avoid, ".")
				for _, item := range tmpAvoid {
					if item != "" {
						avoid[item]++
					}
				}
			}
			fmt.Println(year, month)
		}
	}
	for k, v := range suit {
		fmt.Println(k, v)
	}
	fmt.Println("===========================================================")
	for k, v := range avoid {
		fmt.Println(k, v)
	}
}

func httpSend(url string) ([]byte, error) {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url) // Content-Type post请求必须设置
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}
