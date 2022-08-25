package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	status = map[string]string{}
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		var website []string
		json.NewDecoder(req.Body).Decode(&website)
		go checkWebsiteStatus(website)
	case "GET":
		value := req.URL.Query().Get("website")
		if value != "" {
			jsonVal, err := json.Marshal(status[value])
			if err != nil {
				fmt.Println("Error occurred=", err.Error())
			}
			writer.Write(jsonVal)
			return
		}
		fmt.Println("Before Get Logic is executed=", status)
		jsonVal, err := json.Marshal(status)
		if err != nil {
			fmt.Println("Error occurred=", err.Error())
		}
		writer.Write(jsonVal)
	}
}

func checkWebsiteStatus(website []string) {
	for {
		for _, val := range website {
			_, err := http.Get("http://" + val)
			if err != nil {
				status[val] = "DOWN"
			} else {
				status[val] = "UP"
			}
		}
		time.Sleep(time.Minute)
	}
}
