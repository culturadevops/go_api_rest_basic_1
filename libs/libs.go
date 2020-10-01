package libs

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
}

func Configure(ConfigPath string, ConfigName string) DbConfig {
	config.AddConfigPath(ConfigPath)
	config.SetConfigName(ConfigName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	response := DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
		"",
	}
	return response
}
