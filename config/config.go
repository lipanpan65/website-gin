package config

import (
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
)

func InitConfig() {
	// 获取环境变量，例如 `dev`, `prod` 等
	env := os.Getenv("GO_ENV")
	fmt.Println("env", env)
	if env == "" {
		env = "dev" // 默认加载开发环境配置
	}
	// 设置配置文件名和类型
	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 将配置文件中的值映射到 Config 结构体
	err := viper.Unmarshal(&Conf)
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
