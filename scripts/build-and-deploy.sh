#!/bin/bash

# 项目根目录
PROJECT_DIR="/Users/lipanpan/github/website-gin"
# 服务器目标路径
REMOTE_DIR="/opt/website-gin"
SERVER_IP="47.93.43.223"
USER="root"

# 进入项目根目录
cd $PROJECT_DIR || { echo "Failed to enter project directory $PROJECT_DIR"; exit 1; }

# 设置Go环境变量，构建适用于Linux的二进制文件
GOOS=linux GOARCH=amd64 go build -o website-gin ./cmd || { echo "Go build failed"; exit 1; }
# 上传构建的二进制文件到服务器
scp website-gin $USER@$SERVER_IP:$REMOTE_DIR || { echo "SCP failed"; exit 1; }

# SSH到服务器并执行部署
ssh -t $USER@$SERVER_IP << EOF
  # 进入项目目录
  cd $REMOTE_DIR || { echo "Failed to enter remote directory $REMOTE_DIR"; exit 1; }

  # 设置可执行权限
  chmod +x website-gin || { echo "Failed to set executable permissions"; exit 1; }

  # 如果应用已在运行，先停止它
  if pgrep -f "website-gin" > /dev/null; then
    echo "Stopping the existing app..."
    pkill -f "website-gin" || { echo "Failed to stop the running app"; exit 1; }
  fi

  # 启动应用并将日志重定向到文件
  echo "Starting the app..."
  nohup ./website-gin 2>&1 | while read line; do echo \$(date "+%Y-%m-%d %H:%M:%S") "\$line"; done > $REMOTE_DIR/website-gin.log &

  # 确认应用已启动
  if pgrep -f "website-gin" > /dev/null; then
    echo "Gin app deployed and started successfully"
  else
    echo "Failed to start the Gin app"
    exit 1
  fi
EOF

echo "Build and deployment completed."
