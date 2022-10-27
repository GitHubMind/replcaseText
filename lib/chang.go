package lib

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func OpenFile(address string, transale Transale) {
	log.Println("check ：", address)
	data, err := ioutil.ReadFile(address)
	if err != nil {
		log.Println("error:", err)
	}
	//做分割
	result := strings.Split(string(data), string('\n'))
	for key, _ := range result {
		for i := 0; i < len(transale.Jsontmp); i++ {
			if transale.Jsontmp[i].Key == "" {
				continue
			}
			reg, err := regexp.Compile(transale.Jsontmp[i].Key)
			if err != nil {
				log.Println("error of reg:", err)
				continue
			}
			old := result[key]
			result[key] = reg.ReplaceAllString(result[key], transale.Jsontmp[i].Value)
			if old != result[key] {
				log.Println("Replace:", old, "=>", result[key])
			}
		}
	}
	err2 := ioutil.WriteFile(address, []byte(strings.Join(result, string('\n'))), 0666) //写入文件(字节数组)

	if err2 != nil {
		log.Fatal(err2.Error())

	}
}
