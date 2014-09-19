//用于测试
package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func BasicEncode(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))
}

func main() {
	key := "ID#Mi*eO(32"

	data := make(map[string]interface{})
	data["c"] = 1
	data["p"] = map[string]string{"Mac": "C8-1F-66-1E-FF-E4", "DeviceType": "abc"}
	data["u"] = "-1"
	data["t"] = 1410937034

	ss := fmt.Sprintf("%v|%v|%v|%v", data["u"], data["c"], data["t"], key)
	h := md5.New()
	io.WriteString(h, ss)
	data["s"] = fmt.Sprintf("%x", h.Sum(nil))
	data["d"] = "..."

	//json encode
	j_data, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	//base64 encode
	b_data := base64.URLEncoding.EncodeToString([]byte(j_data))

	url := fmt.Sprintf("http://192.168.51.128:8877/api/%s", b_data)

	//http
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, nil)

	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*;q=0.8")
	request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		returnData := fmt.Sprintf("%v", string(body))
		j_returnData, err := base64.URLEncoding.DecodeString(returnData)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(j_returnData))
	}
}
