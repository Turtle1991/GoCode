package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("err: %v", x)
		}
	}()

	//json encode
	j1 := make(map[string]interface{})
	j1["name"] = "turtle"
	j1["url"] = "http://192.168.51.128:8877/api/"

	js1, err := json.Marshal(j1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js1))

	//json decode
	j2 := make(map[string]interface{})
	err = json.Unmarshal(js1, &j2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", j2)
	fmt.Printf("%v\n", j2)
}
