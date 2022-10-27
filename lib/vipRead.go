package lib

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type ConfigDefine struct {
	TransaleAddresJson string   `json:"transale_addres_json"mapstructure:"transale_addres_json"`
	TransaleAddres     string   `json:"transale_addres"mapstructure:"transale_addres"`
	NotIgnoreAddress   string   `json:"not_ignore_address"mapstructure:"not_ignore_address"`
	Language           string   `json:"language"mapstructure:"language"`
	Catalog_address    []string `json:"language"mapstructure:"catalog_address"`
}

var Config ConfigDefine

func init() {
	viper.AddConfigPath("./")
	viper.AddConfigPath("config.yaml")
	err := viper.ReadInConfig()
	// 读取配置信息
	if err != nil { // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	if err = viper.Unmarshal(&Config); err != nil {
		fmt.Println("viper bug", err)
	}
	log.Println(Config)
}
