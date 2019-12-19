package util

import (
	"encoding/json"
	"log"
)

func ToJson(response interface{}) interface{} {
	response, err := json.Marshal(response)
	if err != nil {
		log.Println("Could not parse response to json : " + err.Error())
	}
	return response
}

func FromJson(_json []byte) interface{} {
	var e interface{}
	err := json.Unmarshal(_json, &e)
	if err != nil {
		log.Fatal("Could not decode json : " + err.Error())
	}
	return e
}
