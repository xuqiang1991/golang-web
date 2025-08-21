#!/bin/bash

# Golang Web 应用启动脚本

echo "=== Golang Web 应用启动脚本 ==="

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "错误: Go 未安装，请先安装 Go 1.21 或更高版本"
    exit 1
fi

# 检查 Go 版本
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "检测到 Go 版本: $GO_VERSION"

# 设置环境变量（默认开发环境）
if [ -z "$GO_ENV" ]; then
    export GO_ENV=development
    echo "设置环境变量: GO_ENV=$GO_ENV"
fi

# 安装依赖
echo "正在安装依赖..."
go mod tidy

# 检查依赖安装是否成功
if [ $? -ne 0 ]; then
    echo "错误: 依赖安装失败"
    exit 1
fi

echo "依赖安装完成"

# 启动应用
echo "正在启动应用..."
echo "环境: $GO_ENV"
echo "应用将在 http://localhost:8080 启动"
echo "按 Ctrl+C 停止应用"

go run main.go
