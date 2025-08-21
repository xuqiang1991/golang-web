package routes

import (
	"golang-web/config"
	"golang-web/handlers"
	"golang-web/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(cfg *config.Config) *gin.Engine {
	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())   // 日志中间件
	r.Use(gin.Recovery()) // 恢复中间件

	// 创建处理器
	authHandler := handlers.NewAuthHandler(cfg)

	// API路由组
	api := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)       // 用户登录
			auth.POST("/register", authHandler.Register) // 用户注册
		}

		// 需要认证的路由
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			// 用户相关
			user := protected.Group("/user")
			{
				user.GET("/profile", authHandler.GetProfile) // 获取用户信息
			}

			// 令牌相关
			token := protected.Group("/token")
			{
				token.POST("/refresh", authHandler.RefreshToken) // 刷新令牌
			}
		}
	}

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
		})
	})

	return r
}
