package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang-web/config"
	"golang-web/database"
	"golang-web/routes"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()
	log.Printf("应用启动，环境: %s, 端口: %s", os.Getenv("GO_ENV"), cfg.Server.Port)

	// 初始化数据库连接
	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer database.CloseDB()
	log.Println("数据库连接成功")

	// 设置路由
	router := routes.SetupRoutes(cfg)

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// 在goroutine中启动服务器
	go func() {
		log.Printf("HTTP服务器启动在端口 %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	// 设置5秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器强制关闭:", err)
	}

	log.Println("服务器已退出")
}
