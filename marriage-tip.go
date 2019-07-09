package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MarriageTip struct {
	Date string
	Tip  string
}

func main() {
	url := "http://clickanthem.com/rcia/api/marriagetip"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var marriageTip MarriageTip
	err := json.Unmarshal(body, &marriageTip)

	//fmt.Println(res)
	//fmt.Println(string(body))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(marriageTip.Date)
	fmt.Println(marriageTip.Tip)
}
