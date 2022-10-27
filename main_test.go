package main

import (
	"log"
	"regexp"
	"testing"
)

func TestName(t *testing.T) {
	//
	////test := reg.ReplaceAllString("*")
	strMap := make(map[string]string, 0)
	str := make(map[string][]string, 0)
	answer := make(map[string][]string, 0)
	result := make(map[string][]bool, 0)
	strMap["*"] = "[^\\s]*"
	str["*"] = []string{"123*", "*", "*.vue", "*.vue", "abc/*"}
	answer["*"] = []string{"123asd", "asd223", "abc.vue",
		"/12313/adsasd/abc.vue",
		"abc/abc/vue.vue"}
	result["*"] = []bool{true, true, true, true, true}

	log.Println(str["*"][0])
	//os.Exit(0)
	for key, input := range strMap {
		log.Println(input)
		for i := 0; i < len(str["*"]); i++ {
			strReg := "^"
			reg, err := regexp.Compile("[" + key + "]+")
			if err != nil {
				log.Println(err)
			}
			strReg += reg.ReplaceAllString(str[key][i], input)
			log.Println("匹配", strReg)
			reg, err = regexp.Compile(strReg)
			if result[key][i] == reg.Match([]byte(answer[key][i])) {
				log.Println(result[key][i])
			} else {
				t.Error("有问题")
			}
		}
		//reg, _ := regexp.Compile("")
		//test := reg.ReplaceAllString("*")
		//拼接
		//reg = "^"
	}

}

func TestPretreatRegCompile(t *testing.T) {
	//a := "bcssas*"
	//b := "a"
	//log.Println(a[strings.Index(a, b)+1:])
	//for i := 0; i < 100; i++ {
	//	value, _, _ := lib.TranSlateApi("en", "测试 我是人类")
	//	log.Println(value.Data.Src)
	//}
	value := "<html>asdasd 我 abc"
	reg, err := regexp.Compile("我")
	if err != nil {
		log.Println("error of reg:", err)
	}

	abc := reg.ReplaceAllString(value, "112233")
	log.Println(abc)
	//if a == result[key] {
	//	log.Println(a)
	//}
}
