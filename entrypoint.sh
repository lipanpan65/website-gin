#!/bin/bash
set -e

echo "🚀 启动Go Gin开发环境..."
echo "📍 应用端口: 9899" 
echo "🐛 调试端口: 2345"
echo "🔥 热重载: 已启用"
echo "⚙️  环境: ${GO_ENV}"

# 检查配置文件
if [ ! -f "config/config-${GO_ENV}.yaml" ]; then
    echo "⚠️ 警告: 配置文件 config/config-${GO_ENV}.yaml 不存在"
fi

# 启动应用
if [ "$1" = "debug" ]; then
    echo "🐛 启动调试模式..."
    exec dlv debug ./cmd/main.go --headless --listen=:2345 --api-version=2 --accept-multiclient
elif [ "$1" = "build" ]; then
    echo "🏗️ 构建模式..."
    exec go build -o ./tmp/main ./cmd/main.go
else
    echo "🔥 启动热重载模式..."
    exec air
fi