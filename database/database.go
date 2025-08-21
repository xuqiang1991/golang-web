package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang-web/config"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
	CREATE TABLE IF NOT EXISTS t_user (
	  id int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
	  username varchar(100) DEFAULT NULL COMMENT '用户名',
	  password varchar(255) DEFAULT NULL COMMENT '密码',
	  email varchar(32) DEFAULT '' COMMENT '邮箱',
	  create_time datetime DEFAULT NULL COMMENT '创建时间',
	  update_time datetime DEFAULT NULL COMMENT '更新时间',
	  PRIMARY KEY (id),
	  UNIQUE KEY idx_user (username) USING BTREE
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
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
	err := DB.QueryRow("SELECT COUNT(*) FROM t_user WHERE username = ?", "admin").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// 插入默认用户，密码为 admin123
		hashedPassword, err := HashPassword("admin123")
		if err != nil {
			return err
		}

		// 获取当前时间
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		_, err = DB.Exec("INSERT INTO t_user (username, password, email, create_time, update_time) VALUES (?, ?, ?, ?, ?)",
			"admin", hashedPassword, "admin@example.com", currentTime, currentTime)
		if err != nil {
			return err
		}

		log.Println("默认用户已创建: admin/admin123")
	}

	return nil
}

// HashPassword 使用 bcrypt 对密码进行哈希加密
func HashPassword(password string) (string, error) {
	// 使用 bcrypt 进行密码哈希，默认成本因子为 12
	// 成本因子越高，哈希越安全但计算时间越长
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("密码哈希失败: %v", err)
	}
	return string(hashedBytes), nil
}

// CheckPassword 验证密码是否匹配哈希值
func CheckPassword(password, hashedPassword string) error {
	// 使用 bcrypt 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("密码验证失败: %v", err)
	}
	return nil
}
