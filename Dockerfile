# 使用Go 1.22作为基础镜像
FROM golang:1.22-alpine

# 设置工作目录
WORKDIR /app

# 安装必要的系统依赖
RUN apk add --no-cache \
    git \
    gcc \
    musl-dev

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -o main cmd/main.go

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./main"]

