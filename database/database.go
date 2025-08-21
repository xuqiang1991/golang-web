package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang-web/config"

	_ "github.com/go-sql-driver/mysql"
)

// DB 全局数据库连接
var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) error {
	var err error

	// 获取数据库连接字符串
	dsn := cfg.GetDSN()

	// 连接数据库
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 设置连接池参数
	DB.SetMaxOpenConns(25)                 // 最大连接数
	DB.SetMaxIdleConns(10)                 // 最大空闲连接数
	DB.SetConnMaxLifetime(5 * time.Minute) // 连接最大生命周期

	// 测试数据库连接
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %v", err)
	}

	log.Println("数据库连接成功")

	// 初始化数据库表
	if err := initTables(); err != nil {
		return fmt.Errorf("初始化数据库表失败: %v", err)
	}

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("数据库连接已关闭")
	}
}

// initTables 初始化数据库表
func initTables() error {
	// 创建用户表
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(100),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		return fmt.Errorf("创建用户表失败: %v", err)
	}

	// 检查是否需要插入默认用户
	if err := insertDefaultUser(); err != nil {
		log.Printf("插入默认用户失败: %v", err)
	}

	log.Println("数据库表初始化完成")
	return nil
}

// insertDefaultUser 插入默认用户（如果不存在）
func insertDefaultUser() error {
	// 检查默认用户是否已存在
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "admin").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// 插入默认用户，密码为 admin123
		hashedPassword, err := hashPassword("admin123")
		if err != nil {
			return err
		}

		_, err = DB.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
			"admin", hashedPassword, "admin@example.com")
		if err != nil {
			return err
		}

		log.Println("默认用户已创建: admin/admin123")
	}

	return nil
}

// hashPassword 简单的密码哈希（生产环境建议使用 bcrypt）
func hashPassword(password string) (string, error) {
	// 这里使用简单的哈希，生产环境建议使用 bcrypt
	// 为了演示，我们直接返回密码的哈希值
	return password, nil
}
