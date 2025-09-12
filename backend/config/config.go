package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Addr            string
		User            string
		Password        string
		Name            string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime int
	}

	CORSConfig struct {
		AllowOrigins string
	}
}

var AppConfig Config

// 初始化配置
func InitConfig() {

	//读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("d:/Codefield/CODE_GO/goweblearning/golang_web/Todolist/backend/config/")
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
