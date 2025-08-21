package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang-web/database"
)

// User 用户模型
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"` // 密码不返回给前端
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, password, email, create_time, update_time FROM t_user WHERE username = ?`

	err := database.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return user, nil
}

// GetUserByID 根据用户ID获取用户
func GetUserByID(userID int) (*User, error) {
	user := &User{}
	query := `SELECT id, username, password, email, create_time, update_time FROM t_user WHERE id = ?`

	err := database.DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return user, nil
}

// CreateUser 创建新用户
func CreateUser(req *RegisterRequest) (*User, error) {
	// 检查用户名是否已存在
	existingUser, err := GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("用户名已存在")
	}

	// 对密码进行哈希加密
	hashedPassword, err := database.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 创建用户
	query := `INSERT INTO t_user (username, password, email, create_time, update_time) VALUES (?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query, req.Username, hashedPassword, req.Email, currentTime, currentTime)
	if err != nil {
		return nil, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 返回新创建的用户
	user := &User{
		ID:       int(userID),
		Username: req.Username,
		Email:    req.Email,
	}

	return user, nil
}

// ValidatePassword 验证用户密码
func (u *User) ValidatePassword(password string) bool {
	// 使用 bcrypt 进行安全的密码比较
	err := database.CheckPassword(password, u.Password)
	return err == nil
}
