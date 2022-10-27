package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//https://www.alapi.cn/api/view/42
//to Language type
//q  word content

type TranSlateApiType struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  Data   `json:"data"`
	Time  int    `json:"time"`
	LogID int64  `json:"log_id"`
}
type Data struct {
	From string `json:"from"`
	To   string `json:"to"`
	Src  string `json:"src"`
	Dst  string `json:"dst"`
}

func TranSlateApi(to, q string) (result string, status int, err error) {
	if len(q) == 0 || q == "" {
		return
	}

	mapReq := make(map[string]string, 0)
	mapReq["form"] = "auto"
	mapReq["token"] = "DWWyL9u3bFBRb5ku"
	mapReq["to"] = to
	mapReq["q"] = q
	value, err := json.Marshal(mapReq)
	var valuebody []byte
	var body TranSlateApiType
	for {
		valuebody, status = Post("https://v2.alapi.cn/api/fanyi", value)
		//log.Println(string(valuebody))
		json.Unmarshal(valuebody, &body)
		result = body.Data.Dst
		if status == 200 && body.Code == 200 {
			log.Println("请求成功：200")
			break
		}
	}
	//log.Println(status)
	//if len(valuebody)==0{}

	return
}

func Post(urlStr string, json []byte) (body []byte, status int) {
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		//Transport: netTransport,
		Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client err:", resp)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	status = resp.StatusCode
	return
}
