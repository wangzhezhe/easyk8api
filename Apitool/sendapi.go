package sendapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

func sendbase(client *http.Client, apitype string, url string, filename string) (int, map[string]interface{}) {
	var finalreqest *http.Request
	if apitype == "POST" {
		current_dir, _ := os.Getwd()
		filename = current_dir + "/" + filename
		fmt.Println(filename)
		byt, err := ioutil.ReadFile(filename)

		if err != nil {
			panic(err.Error())
		}

		reqest, _ := http.NewRequest(apitype, url, bytes.NewBuffer(byt))
		reqest.Header.Set("Content-Type", "application/json")
		finalreqest = reqest

	} else {
		reqest, err := http.NewRequest(apitype, url, nil)
		if err != nil {
			panic(err.Error())
		}
		finalreqest = reqest

	}

	response, _ := client.Do(finalreqest)

	//body为 []byte类型
	body, _ := ioutil.ReadAll(response.Body)

	//decoding the []body into the map
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	//fmt.Println(result)
	status := response.StatusCode
	return status, result

}

func sendGet(client *http.Client, url string) (int, map[string]interface{}) {
	var result map[string]interface{}
	var status int
	fmt.Println("sent get request ")
	//using "" to represent the nil of string
	status, result = sendbase(client, "GET", url, "")

	return status, result

}

func sendPost(client *http.Client, url string, filename string) (int, map[string]interface{}) {

	var result map[string]interface{}
	var status int
	fmt.Println("sent post request ")

	status, result = sendbase(client, "POST", url, filename)

	return status, result

}

func sendDelete(client *http.Client, url string) (int, map[string]interface{}) {
	var result map[string]interface{}
	var status int
	fmt.Println("sent delete request ")
	status, result = sendbase(client, "DELETE", url, "")
	return status, result

}

func sendPut(client *http.Client, url string) (int, map[string]interface{}) {
	var result map[string]interface{}
	var status int
	fmt.Println("sent put request ")
	status, result = sendbase(client, "PUT", url, "")
	return status, result

}

//problems in using patch
func sendPatch(client *http.Client, url string, filename string) (int, map[string]interface{}) {

	var result map[string]interface{}
	var status int
	fmt.Println("sent patch request ")
	status, result = sendbase(client, "PATCH", url, "")

	return status, result

}

func Sendapi(apitype string, host string, port string, version string, commands []string, filename string) (int, map[string]interface{}) {

	client := &http.Client{}
	fmt.Println(reflect.TypeOf(client))
	//注意前面要加上http://
	url := "http://" + host + ":" + port + "/api" + "/" + version
	for _, str := range commands {
		url = url + "/" + str

	}
	fmt.Println(url)

	var result map[string]interface{}
	var status int
	if apitype == "GET" {
		status, result = sendGet(client, url)
	} else if apitype == "POST" {
		status, result = sendPost(client, url, filename)

	} else if apitype == "DELETE" {
		status, result = sendDelete(client, url)

	} else if apitype == "PUT" {
		status, result = sendGet(client, url)

	} else if apitype == "PATCH" {
		status, result = sendPatch(client, url, filename)

	} else {
		panic("error api type")

	}

	return status, result
}
