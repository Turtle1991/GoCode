package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
)

func main() {
	// response, _ := http.Get("http://www.baidu.com")
	response, _ := http.Get("http://data.5054399.net")

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	code := response.StatusCode
	fmt.Println("statusCode: ", code)
}*/

/*package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) say() {
	fmt.Printf("name is %s, age is %d.", p.name, p.age)
}

func main() {
	my := &person{}
	my.say()
}*/

/*package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://www.baidu.com", nil)*/

// request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
/*	request.Header.Set("Accept-Charset", "utf-8,GBK;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Set("Accept-Language", "utf-8,GBK;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}*/

/*************************************************************************/
//创建web服务端！！！！重要 学习！！！！
/*package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Hello turtle.<a href='http://www.baidu.com' target='_blank'>你好！</a></h1>"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)
}*/

/*package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)*/

// func main() {
	// client := &http.Client{}
	// request, _ := http.NewRequest("POST", "http://192.168.51.128:8877/api/eyJjIjotMTAwMCwicCI6bnVsbCwidSI6IiIsInQiOjE0MTA5NTA4NjYsInMiOiIiLCJkIjoi6YCa55So6ZSZ6K-vIn0=", nil)
	// request, _ := http.NewRequest("POST", "http://192.168.51.128:8877/api/eyJjIjoxLCJwIjp7Ik1hYyI6IkM4LTFGLTY2LTFFLUZGLUU4IiwiRGV2aWNlVHlwZSI6ImFhYWEifSwidSI6Ii0xIiwidCI6MTQxMDkzNzAzNCwicyI6ImE3MzhhOWYyOGRhZDZkNDAxNDg2YmIyYTRiMTg4M2Q4IiwiZCI6Ii4uLi4ifQ==", nil)

	// request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*;q=0.8")
	// request.Header.Set("Accept-Charset", "utf-8,GBK;q=0.7,*;q=0.3")
	/*request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(string(body))

	}
}
