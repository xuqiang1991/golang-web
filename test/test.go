package main

import (
	"fmt"
	"os"

	"golang-web/config"
)

func main() {
	// 测试配置加载
	fmt.Println("测试配置加载...")

	// 设置环境变量
	os.Setenv("GO_ENV", "development")

	// 加载配置
	cfg := config.LoadConfig()

	fmt.Printf("服务器配置: 端口=%s, 模式=%s\n", cfg.Server.Port, cfg.Server.Mode)
	fmt.Printf("数据库配置: 主机=%s, 端口=%s, 数据库=%s\n",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
	fmt.Printf("JWT配置: 过期时间=%d小时\n", cfg.JWT.Expire)

	// 测试数据库连接字符串
	dsn := cfg.GetDSN()
	fmt.Printf("数据库连接字符串: %s\n", dsn)

	fmt.Println("配置测试完成!")
}
