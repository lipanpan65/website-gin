# Makefile for Go Gin Docker Development Environment

.PHONY: help check-docker watch debug docker-dev docker-build docker-up docker-stop docker-clean docker-logs docker-shell docker-db-shell docker-restart docker-test docker-migrate

# 默认目标
help:
	@echo "🚀 Go Gin Docker开发环境管理"
	@echo ""
	@echo "本地开发命令:"
	@echo "  make watch           - 启动热重载开发模式 (监听文件变化自动重启)"
	@echo "  make debug           - 启动本地调试模式 (使用Delve调试器)"
	@echo ""
	@echo "Docker 相关命令:"
	@echo "  make docker-dev      - 启动Docker开发环境 (构建并运行，支持热加载)"
	@echo "  make docker-build    - 构建Docker镜像"
	@echo "  make docker-up       - 启动容器"
	@echo "  make docker-stop     - 停止并删除容器"
	@echo "  make docker-clean    - 清理所有Docker资源"
	@echo "  make docker-logs     - 查看容器日志"
	@echo "  make docker-shell    - 进入Go容器shell"
	@echo "  make docker-db-shell - 进入数据库容器shell"
	@echo "  make docker-restart  - 重启开发环境"
	@echo "  make docker-test     - 运行测试"
	@echo "  make docker-migrate  - 运行数据库迁移"
	@echo ""
	@echo "其他命令:"
	@echo "  make help            - 显示此帮助信息"
	@echo ""
	@echo "💡 调试说明:"
	@echo "  本地调试: 在Cursor中按F5选择'🚀 Go Gin Debug'"
	@echo "  远程调试: 先运行'make debug'，再按F5选择'🔍 Go Gin Remote Debug'"

# 热重载开发模式
watch:
	@echo "🔥 启动热重载开发模式..."
	@echo "📁 监听文件变化，自动重启应用"
	@echo "🌐 服务将在 http://localhost:9899 启动"
	@if ! command -v air > /dev/null 2>&1; then \
		echo "❌ air 未安装，请先安装: go install github.com/cosmtrek/air@v1.49.0"; \
		exit 1; \
	fi
	@air

# 本地调试模式
debug:
	@echo "🐛 启动Go Gin本地调试模式..."
	@echo "📍 调试服务器将在 http://localhost:9899 启动"
	@echo "🔍 使用Delve调试器，在Cursor中按F5开始调试"
	@if ! command -v dlv > /dev/null 2>&1; then \
		echo "❌ Delve未安装，正在安装..."; \
		go install github.com/go-delve/delve/cmd/dlv@latest; \
		if command -v asdf > /dev/null 2>&1; then \
			echo "🔄 更新asdf shims..."; \
			asdf reshim golang; \
		fi; \
	fi
	@export GO_ENV=dev && dlv debug ./cmd/main.go --headless --listen=:2345 --api-version=2 --accept-multiclient

# 检查Docker是否运行
check-docker:
	@if ! docker info > /dev/null 2>&1; then \
		echo "❌ Docker未运行，请先启动Docker"; \
		exit 1; \
	fi

# 启动开发环境
docker-dev: check-docker
	@echo "🚀 启动Go Gin Docker开发环境..."
	@echo "📦 构建Docker镜像..."
	docker-compose -f docker-compose.dev.yml build
	@echo "🔄 启动开发服务器 (支持热加载)..."
	docker-compose -f docker-compose.dev.yml up

# 构建镜像
docker-build: check-docker
	@echo "📦 构建Docker镜像..."
	docker-compose -f docker-compose.dev.yml build

# 启动容器
docker-up: check-docker
	@echo "🔄 启动开发服务器..."
	docker-compose -f docker-compose.dev.yml up -d
	@echo "✅ 开发环境已启动！"
	@echo "🌐 Go Gin访问地址: http://localhost:8080"
	@echo "🗄️  MySQL访问地址: localhost:3306"
	@echo "🔥 热加载已启用，修改代码会自动重新编译"

# 停止容器
docker-stop:
	@echo "🛑 停止Docker开发环境..."
	docker-compose -f docker-compose.dev.yml down

# 清理所有Docker资源
docker-clean:
	@echo "🧹 清理Docker资源..."
	docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans
	docker system prune -f

# 查看日志
docker-logs:
	@echo "📋 查看容器日志..."
	docker-compose -f docker-compose.dev.yml logs -f

# 进入Go容器shell
docker-shell:
	@echo "🐚 进入Go容器shell..."
	docker-compose -f docker-compose.dev.yml exec website-gin-dev sh

# 进入数据库容器shell
docker-db-shell:
	@echo "🗄️ 进入数据库容器shell..."
	docker-compose -f docker-compose.dev.yml exec db-dev mysql -u website_user -pwebsite_password website_db

# 重启开发环境
docker-restart: docker-stop docker-up
	@echo "🔄 开发环境已重启！"

# 运行测试
docker-test:
	@echo "🧪 运行测试..."
	docker-compose -f docker-compose.dev.yml exec website-gin-dev go test ./...

# 运行数据库迁移
docker-migrate:
	@echo "🔄 运行数据库迁移..."
	docker-compose -f docker-compose.dev.yml exec website-gin-dev go run cmd/main.go migrate



