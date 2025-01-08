package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	AppMode string `mapstructure:"app_mode"`
	Port    string `mapstructure:"port"`
	DBUser  string `mapstructure:"db_user"`
	DBPass  string `mapstructure:"db_pass"`
	DBName  string `mapstructure:"db_name"`
	DBHost  string `mapstructure:"db_host"`
	DBPort  string `mapstructure:"db_port"`
}

var (
	Conf Config
	DB   *gorm.DB
	//go:embed config-dev.yaml
	//go:embed config-prod.yaml
	configFile embed.FS
)

func InitConfig() {
	// 获取环境变量，例如 `dev`, `prod` 等
	env := os.Getenv("GO_ENV")
	fmt.Println("env", env)
	if env == "" {
		env = "dev" // 默认加载开发环境配置
	}

	// 根据环境选择配置文件
	configFileName := fmt.Sprintf("config-%s.yaml", env)
	fmt.Println("configFileName", configFileName)

	data, err := configFile.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("failed to read embedded config file %s: %v", configFileName, err)
	}

	// 使用 viper 解析嵌入的配置文件
	// 设置配置文件名和类型
	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err = viper.ReadConfig(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("failed to parse embedded config file: %v", err)
	}

	//// 读取配置文件
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("Error reading config file: %v", err)
	//}

	// 将配置文件中的值映射到 Config 结构体
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// 设置 Gin 的模式
	if Conf.AppMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if Conf.AppMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	fmt.Printf("Loaded Config: %+v\n", Conf)
}

func InitDB() {
	// 构建 MySQL 数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Conf.DBUser, Conf.DBPass, Conf.DBHost, Conf.DBPort, Conf.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}

//GOOS=linux GOARCH=amd64 go build -o website-gin main.go
//scp website-gin root@47.93.43.223:/opt/website-gin
