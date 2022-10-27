package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"sync"
)

type Transale struct {
	FileAddesJson string
	Jsontmp       JsonType
	JsonMap       map[string]string
}
type JsonType []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//存放置换规则

//
var RWLcok sync.RWMutex

func (tra *Transale) writeFileJson(address string) {
	Mar, err := json.Marshal(tra.Jsontmp)
	if err != nil {
		log.Println("err:", err)
	}
	err = ioutil.WriteFile(address, Mar, 0666) //写入文件(字节数组)
	if err != nil {
		log.Println("err:", err)
	}

}
func (tra *Transale) OpenFileTxt(address, language string) {
	data, err := ioutil.ReadFile(address)
	if err != nil {
		log.Println("error:", err)
	}

	result := strings.Split(string(data), string('\n'))
	sync := sync.WaitGroup{}
	number := len(result)
	for i := 0; i < len(result); i++ {
		//排除空的
		if result[i] == "" || len(result[i]) == 0 {
			continue
		}
		go func(index int) {
			sync.Add(1)
			type tmp struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}
			var value tmp
			value.Key = result[index]
			valueText, _, _ := TranSlateApi(language, value.Key)
			value.Value = valueText
			tra.Jsontmp = append(tra.Jsontmp, value)
			log.Println("剩下请求数", number)
			number--
			sync.Done()
		}(i)
	}
	sync.Wait()
	log.Println("end")
	tra.writeFileJson(tra.FileAddesJson)
}
func (tra *Transale) ReadJsonFile(language string) {
	data, err := ioutil.ReadFile(tra.FileAddesJson)
	if err != nil {
		log.Println("error:", err)
	}
	err = json.Unmarshal(data, &tra.Jsontmp)

	if err != nil {
		log.Println("error:", err)
	} else {

		tra.JsonMap = make(map[string]string, 0)
		sync := sync.WaitGroup{}
		for i := 0; i < len(tra.Jsontmp); i++ {
			if tra.Jsontmp[i].Key == "" && len(tra.Jsontmp[i].Key) == 0 {
				continue
			}
			sync.Add(1)
			go func(index int) {
				if tra.Jsontmp[index].Value == "" {
					value, _, _ := TranSlateApi(language, tra.Jsontmp[index].Key)
					tra.Jsontmp[index].Value = value
				}
				RWLcok.Lock()
				tra.JsonMap[tra.Jsontmp[index].Key] = tra.Jsontmp[index].Value
				RWLcok.Unlock()
				sync.Done()
			}(i)
		}
		//sort  length more priority
		sort.Slice(tra.Jsontmp, func(i, j int) bool {
			if len(tra.Jsontmp[i].Key) > len(tra.Jsontmp[j].Key) {
				return true
			}
			return false
		})
		sync.Wait()
		log.Println(tra.Jsontmp)
		tra.writeFileJson(tra.FileAddesJson)

	}
	log.Println(tra.FileAddesJson, "json status is success")
}
