package config

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestInitconf(t *testing.T) {

	//读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("can not found yaml, error: %v\n", err)
		} else {
			log.Fatalf("Read in config  error: %v\n", err)
		}
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Init config  error: %v\n", err)
	}

}
