@echo off
chcp 65001 >nul

echo === Golang Web 应用启动脚本 ===

REM 检查 Go 是否安装
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo 错误: Go 未安装，请先安装 Go 1.21 或更高版本
    pause
    exit /b 1
)

REM 检查 Go 版本
for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
set GO_VERSION=%GO_VERSION:go=%
echo 检测到 Go 版本: %GO_VERSION%

REM 设置环境变量（默认开发环境）
if "%GO_ENV%"=="" (
    set GO_ENV=development
    echo 设置环境变量: GO_ENV=%GO_ENV%
)

REM 安装依赖
echo 正在安装依赖...
go mod tidy

REM 检查依赖安装是否成功
if %errorlevel% neq 0 (
    echo 错误: 依赖安装失败
    pause
    exit /b 1
)

echo 依赖安装完成

REM 启动应用
echo 正在启动应用...
echo 环境: %GO_ENV%
echo 应用将在 http://localhost:8080 启动
echo 按 Ctrl+C 停止应用

go run main.go

pause
