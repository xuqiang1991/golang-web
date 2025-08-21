-- Golang Web 应用数据库初始化脚本

-- 创建开发环境数据库
CREATE DATABASE IF NOT EXISTS golang_dev CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建生产环境数据库
CREATE DATABASE IF NOT EXISTS golang_web CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用开发环境数据库
USE golang_dev;

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码',
    email VARCHAR(100) COMMENT '邮箱',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_username (username),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 插入默认用户（密码: admin123）
INSERT INTO users (username, password, email) VALUES 
('admin', 'admin123', 'admin@example.com')
ON DUPLICATE KEY UPDATE updated_at = CURRENT_TIMESTAMP;

-- 使用生产环境数据库
USE golang_web;

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码',
    email VARCHAR(100) COMMENT '邮箱',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_username (username),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 插入默认用户（密码: admin123）
INSERT INTO users (username, password, email) VALUES 
('admin', 'admin123', 'admin@example.com')
ON DUPLICATE KEY UPDATE updated_at = CURRENT_TIMESTAMP;

-- 显示创建的表结构
SELECT 'golang_dev 数据库表结构:' AS info;
USE golang_dev;
SHOW TABLES;
DESCRIBE users;

SELECT 'golang_web 数据库表结构:' AS info;
USE golang_web;
SHOW TABLES;
DESCRIBE users;

-- 显示用户数据
SELECT 'golang_dev 用户数据:' AS info;
USE golang_dev;
SELECT id, username, email, created_at FROM users;

SELECT 'golang_web 用户数据:' AS info;
USE golang_web;
SELECT id, username, email, created_at FROM users;
