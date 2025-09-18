package config

import (
	"fmt"

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
func InitConfig() error {

	//读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("d:/Codefield/CODE_GO/goweblearning/golang_web/Todolist/backend/config/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("can not found yaml, error: %v", err) //错误内容的开头是小写
		} else {
			return fmt.Errorf("read in config  error: %v", err)
		}
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("init config  error: %v", err)
	}
	return nil
}
