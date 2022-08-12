package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	status = map[string]string{}
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		var website []string
		json.NewDecoder(req.Body).Decode(&website)
		for _, val := range website {
			_, err := http.Get("http://" + val)
			if err != nil {
				status[val] = "Down"
			} else {
				status[val] = "UP"
			}
		}
		fmt.Println("After Post Logic is executed=", status)
	case "GET":
		fmt.Println("Before Get Logic is executed=", status)
		jsonVal, err := json.Marshal(status)
		if err != nil {
			fmt.Println("Error occured=", err.Error())
		}
		writer.Write(jsonVal)
	}
}
