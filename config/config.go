package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key"`
	Expire    int    `mapstructure:"expire"` // 过期时间（小时）
}

// LoadConfig 加载配置文件
func LoadConfig() *Config {
	// 获取环境变量
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development" // 默认使用开发环境
	}

	// 设置配置文件路径
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("警告: 无法读取配置文件 %s.yaml，使用默认配置", env)
		// 使用默认配置
		return getDefaultConfig(env)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	return &config
}

// getDefaultConfig 获取默认配置
func getDefaultConfig(env string) *Config {
	if env == "production" {
		return &Config{
			Server: ServerConfig{
				Port: "8080",
				Mode: "release",
			},
			Database: DatabaseConfig{
				Host:     "localhost",
				Port:     "3306",
				Username: "root",
				Password: "123456",
				Database: "golang_web",
			},
			JWT: JWTConfig{
				SecretKey: "your-secret-key-change-in-production",
				Expire:    24, // 24小时
			},
		}
	}

	// 开发环境
	return &Config{
		Server: ServerConfig{
			Port: "8080",
			Mode: "debug",
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			Database: "golang_dev",
		},
		JWT: JWTConfig{
			SecretKey: "dev-secret-key",
			Expire:    24, // 24小时
		},
	}
}

// GetDSN 获取数据库连接字符串
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Database,
	)
}
