// 模仿http post
package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 鉴权密钥
const AccessToken = "Bearer X6Qhx6-jQ0KZl9dnbR03Ew"

func HttpPost(url string, data map[string]interface{}) ([]byte, error) {
	url = "http://127.0.0.1:9999" + url
	dataJson, _ := json.Marshal(data)
	body := bytes.NewReader(dataJson)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	fmt.Println(err)
	req.Header.Add("Authorization", "Bearer X6Qhx6-jQ0KZl9dnbR03Ew")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
}

func main(){
	url := "/v5/apigw/client/api4"
	var data =make(map[string]interface{})
	data["name"] = "http client"
	data["group_name"] = "auth"
	data["status"] = "enable"
	data["uri"] = "sadasd"
	resp,err:=HttpPost(url, data)
	fmt.Println(err)
	fmt.Println(string(resp))
}