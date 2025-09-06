# Go Gin 开发环境部署指南

## 项目结构说明

本项目采用多项目工作区结构：
```
website/                          # 工作区根目录
├── website-django/               # Django 后端项目
├── website-react/                # React 前端项目
├── website-gin/                  # Go 后端项目
│   ├── cmd/                      # 主程序入口
│   ├── config/                   # 配置文件
│   ├── controllers/              # 控制器
│   ├── routes/                   # 路由配置
│   ├── middleware/               # 中间件
│   ├── internal/                 # 内部包
│   ├── utils/                    # 工具函数
│   ├── dto/                      # 数据传输对象
│   ├── go.mod                    # Go 模块文件
│   ├── go.sum                    # Go 依赖校验文件
│   ├── .tool-versions            # asdf 版本管理文件
│   └── Makefile                  # 构建脚本
└── website.code-workspace        # VS Code/Cursor 工作区配置
```

## 基于本地 Go 环境的开发部署

### 环境要求
- Go 1.25.0+（推荐使用 asdf 版本管理）
- MySQL 数据库（共享 Django 项目数据库）
- VS Code 或 Cursor IDE

### 快速开始

#### 1. 克隆项目并进入工作区根目录

```bash
git clone <repository-url>
cd website/website-gin
```

#### 2. 配置 Go 版本（推荐使用 asdf）

**如果使用 asdf 管理 Go 版本：**
```bash
# 项目已包含 .tool-versions 文件，指定 Go 1.25.0
# 如果未安装该版本，执行：
asdf install golang 1.25.0

# 验证版本
go version
# 应显示：go version go1.25.0 darwin/arm64

# 确认 asdf 配置
asdf current golang
# 应显示：golang  1.25.0  /path/to/website-gin/.tool-versions  true
```

#### 3. 安装 Go 依赖

```bash
# 下载并安装项目依赖
go mod download
go mod tidy
```

#### 4. 配置数据库连接

项目使用嵌入式配置文件管理，配置文件已预配置。

> **注意**：项目共享 Django 后端的阿里云 RDS MySQL 数据库

#### 5. 启动开发服务器

**推荐：使用热重载模式（自动重启）**

```bash
# 启动热重载开发模式（推荐）
make watch
```

**传统方式：手动启动**

```bash
# 启动开发环境（默认）
go run cmd/main.go

# 或者显式指定开发环境
export GO_ENV=dev
go run cmd/main.go

# 启动生产环境模式
export GO_ENV=prod
go run cmd/main.go
```

#### 6. 验证服务启动

服务启动成功后会显示：
```
Configuration initialized: {AppMode:debug Port::9899 ...}
successfully connected to database
Database initialized
Routes initialized
Starting server on port :9899
[GIN-debug] Listening and serving HTTP on :9899
```

#### 7. 访问应用

打开浏览器访问：http://localhost:9899

## API 端点

项目启动后可用的 API 端点：

### 用户管理
- `POST /api/v1/users/` - 创建用户
- `GET /api/v1/users/` - 获取所有用户

### 主题管理
- `POST /api/v1/subject/` - 创建主题

### 话题管理
- `GET /api/v1/topics/` - 查询话题
- `POST /api/v1/topics/` - 创建话题

## 基于 Docker 的开发环境

### Docker 相关命令

```bash
# 查看所有可用命令
make help

# 本地开发命令
make watch             # 启动热重载开发模式 (监听文件变化自动重启)

# Docker 环境管理
make docker-build      # 构建Docker镜像
make docker-dev        # 启动Docker开发环境 (构建并运行，支持热加载)
make docker-debug      # 启动Docker调试环境 (支持远程调试)
make docker-up         # 启动容器
make docker-stop       # 停止并删除容器
make docker-stop-debug # 停止调试容器
make docker-clean      # 清理所有Docker资源
make docker-logs       # 查看容器日志
make docker-shell      # 进入Go容器shell
make docker-db-shell   # 进入数据库容器shell
make docker-restart    # 重启开发环境
make docker-test       # 运行测试
make docker-migrate    # 运行数据库迁移
```

### 启动 Docker 开发环境

#### 热重载开发模式
```bash
# 构建并启动完整的开发环境
make docker-dev
```

#### Docker调试模式
```bash
# 启动Docker调试环境
make docker-debug
```

Docker 环境包含：
- Go Gin 应用容器（支持热加载/调试）
- MySQL 数据库容器
- 端口映射：
  - Go 应用：`http://localhost:9899`
  - 调试端口：`localhost:2345` (仅调试模式)
  - MySQL：`localhost:3306`

#### Docker调试环境特性
- **自动化管理**：自动构建、启动、停止调试容器
- **代码同步**：实时挂载本地代码到容器
- **网络隔离**：独立的Docker网络环境
- **调试就绪**：Delve调试服务器自动启动
- **容器命名**：调试容器名为`gin-debug`，便于管理

## 开发工具和配置

### VS Code/Cursor 工作区配置

1. **打开工作区文件**
   
   使用 Cursor/VS Code 打开 `website.code-workspace` 文件，该文件包含多项目配置。

2. **Go 扩展推荐**
   
   确保安装以下扩展：
   - Go (Google)
   - Go Test Explorer

#### 热重载开发模式

项目已集成 [Air](https://github.com/cosmtrek/air) 热重载工具，支持代码自动重启功能。

#### 1. 安装 Air（如果未安装）

```bash
# 安装最新版本（适配Go 1.25.0）
go install github.com/air-verse/air@latest

# 重新生成 shims（如果使用 asdf）
asdf reshim golang

# 验证安装
air -v
# 应该显示：v1.62.0, built with Go go1.25.0
```

> **⚠️ 注意**：如果遇到 `No version is set for command air` 错误，请查看故障排除部分的 "Air 热重载工具找不到" 解决方案。

#### 2. 启动热重载模式

```bash
make watch
```

#### 3. 热重载特性

- **自动监听**：监听 `.go`、`.yaml`、`.yml`、`.json`、`.html` 等文件变化
- **自动重启**：文件变化时自动重新编译并重启应用
- **排除目录**：自动排除 `tmp`、`vendor`、`node_modules`、`.git`、`docs` 等目录
- **构建日志**：编译错误会保存在 `build-errors.log` 文件中
- **彩色输出**：不同类型的日志使用不同颜色区分

#### 4. Air 配置文件

项目包含 `.air.toml` 配置文件，主要配置：

```toml
[build]
  cmd = "go build -o ./tmp/main ./cmd/main.go"
  bin = "./tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html", "yaml", "yml", "json"]
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "node_modules", ".git", "docs"]
  exclude_regex = ["_test.go"]
```

## Go调试环境配置

### 环境准备和安装

#### 1. 安装Delve调试器

**自动安装（推荐）：**
```bash
# 使用make命令自动安装
make debug  # 首次运行会自动安装dlv并更新asdf shims
```

**手动安装：**
```bash
# 安装Delve调试器
go install -v github.com/go-delve/delve/cmd/dlv@latest

# 如果使用asdf管理Go版本，更新shims
asdf reshim golang

# 验证安装
dlv version
# 应显示：Delve Debugger Version: 1.25.2 (latest)
```

#### 2. 项目调试配置文件

项目已自动创建以下调试配置文件：

**`.vscode/launch.json`** - 调试启动配置：
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "🚀 Go Gin Debug",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/main.go",
            "env": {"GO_ENV": "dev"}
        },
        {
            "name": "🔍 Go Gin Remote Debug", 
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "port": 2345,
            "host": "127.0.0.1"
        }
    ]
}
```

**`.vscode/settings.json`** - Go扩展配置：
- 自动更新Go工具
- 配置Delve调试参数
- 优化调试体验设置

### 调试方式

项目支持三种调试方式：**本地直接调试**、**本地远程调试**和**Docker远程调试**。

- **本地直接调试**：在本地环境中直接启动调试，最简单
- **本地远程调试**：本地启动调试服务器，支持热连接
- **Docker远程调试**：在Docker容器中启动调试服务器，隔离环境

#### 1. 直接调试（推荐新手）

使用Cursor内置的Go调试功能，最简单的调试方式。

**步骤：**
1. 在Cursor中打开 `website-gin` 项目文件夹
2. 在 `cmd/main.go` 或其他Go文件中设置断点（点击行号左侧）
3. 按 `F5` 选择 `🚀 Go Gin Debug`
4. 应用会在调试模式下启动，遇到断点时自动暂停

**特点：**
- 🟢 简单易用，一键启动
- 🟢 支持断点、变量查看、调用栈
- 🟢 完全集成在Cursor IDE中
- 🟢 自动设置环境变量
- 🟡 每次调试需要重新启动应用

**示例调试位置：**
```go
func main() {
    // 在这里设置断点测试
    config.InitConfig()  // <- 第22行，推荐的第一个断点位置
    log.Printf("Configuration initialized: %+v", config.Conf)
}
```

#### 2. 远程调试（推荐高级用户）

使用Delve调试器的远程调试功能，支持热连接调试。

**步骤：**
1. **启动调试服务器**
   ```bash
   make debug
   ```
   
   等待服务器启动成功：
   ```
   🐛 启动Go Gin本地调试模式...
   📍 调试服务器将在 http://localhost:9899 启动  
   🔍 使用Delve调试器，在Cursor中按F5开始调试
   API server listening at: [::]:2345
   warn layer=rpc Listening for remote connections (connections are not authenticated nor encrypted)
   ```

2. **连接调试器**
   - 在代码中设置断点
   - 按 `F5` 选择 `🔍 Go Gin Remote Debug`
   - 调试器连接到运行中的应用

3. **断开和重连**
   - 可以随时断开调试器（应用继续运行）
   - 需要调试时重新连接即可

**特点：**
- 🟢 支持热连接/断开调试器
- 🟢 应用持续运行，无需重启
- 🟢 支持多客户端同时调试
- 🟢 适合长期开发调试
- 🟡 需要额外启动调试服务器

#### 3. Docker远程调试（推荐团队开发）

使用Docker容器中的Delve调试器，提供隔离的调试环境。

**步骤：**
1. **启动Docker调试环境**
   ```bash
   make docker-debug
   ```
   
   系统会自动执行以下操作：
   ```
   🐛 启动Go Gin Docker调试环境...
   📍 应用端口: 9899
   🔍 调试端口: 2345 (Delve API)
   🔄 停止现有容器...
   🚀 启动调试容器...
   ✅ 调试环境已启动！
   ```

2. **连接调试器**
   - 在代码中设置断点
   - 按 `F5` 选择 `🔍 Go Gin Remote Debug`
   - 调试器连接到Docker容器中的应用

3. **管理调试容器**
   ```bash
   # 查看调试日志
   docker logs gin-debug
   
   # 停止调试环境
   make docker-stop-debug
   ```

**特点：**
- 🟢 完全隔离的调试环境
- 🟢 与生产环境更接近
- 🟢 支持热连接/断开调试器
- 🟢 团队成员环境一致
- 🟢 自动化容器管理
- 🟡 需要Docker环境支持

**适用场景：**
- 团队协作开发
- 环境依赖复杂的调试
- 需要模拟生产环境
- 多服务联调场景

#### 4. F5按键调试选择

**问题解决**：项目已解决多项目工作区中F5默认选择Django调试的问题。

**当前行为**：
- 在Go项目目录中按 `F5`：优先显示Go调试选项
- 在Django项目目录中按 `F5`：显示Python/Django调试选项
- 在工作区根目录按 `F5`：显示所有项目的调试选项

### 调试功能详解

#### 1. 断点功能

**设置断点**：
- 普通断点：在代码行号左侧点击红点
- 条件断点：右键断点设置条件表达式
- 日志断点：输出调试信息而不暂停执行

**断点管理**：
- 断点窗口查看所有断点
- 支持启用/禁用断点
- 支持导入/导出断点配置

#### 2. 变量查看

**变量窗口功能**：
- 自动显示当前作用域所有变量
- 支持展开复杂数据结构（struct、map、slice）
- 实时变量值更新
- 支持变量值修改（在暂停时）

**监视表达式**：
- 添加自定义表达式监视
- 支持函数调用和复杂表达式
- 实时求值显示

#### 3. 调用栈

**栈帧导航**：
- 查看完整的函数调用链
- 点击栈帧跳转到对应代码位置
- 查看每个栈帧的局部变量

**协程调试**：
- 支持多协程调试
- 可在不同协程间切换
- 查看协程状态和栈信息

#### 4. 调试控制

**执行控制**：
- Continue (`F5`)：继续执行直到下一个断点
- Step Over (`F10`)：单步执行，不进入函数
- Step Into (`F11`)：单步执行，进入函数内部
- Step Out (`Shift+F11`)：跳出当前函数
- Restart (`Ctrl+Shift+F5`)：重启调试会话
- Stop (`Shift+F5`)：停止调试

### 调试配置说明

#### 1. 环境变量配置

**自动设置的环境变量**：
- `GO_ENV=dev`：强制使用开发环境配置
- 使用配置文件：`config/config-dev.yaml`
- 数据库连接：自动连接开发环境数据库

#### 2. 端口配置

**调试相关端口**：
- 应用服务端口：`9899` (http://localhost:9899)
- 远程调试端口：`2345` (Delve API服务器)
- 数据库端口：`3306` (MySQL)

**Docker调试端口映射**：
- 本地 `9899` → 容器 `9899` (应用端口)
- 本地 `2345` → 容器 `2345` (调试端口)

#### 3. Delve调试器配置

**调试器参数**：
```bash
dlv debug ./cmd/main.go \
  --headless \              # 无GUI模式
  --listen=:2345 \          # 监听端口
  --api-version=2 \         # API版本
  --accept-multiclient      # 支持多客户端
```

**性能优化配置**：
- `followPointers: true`：跟踪指针引用
- `maxVariableRecurse: 3`：变量递归深度
- `maxStringLen: 400`：字符串最大长度
- `maxArrayValues: 400`：数组最大元素数
- `maxStructFields: -1`：结构体字段无限制

### 常见调试场景

#### 1. API请求调试

**本地调试：**
在路由处理函数中设置断点：
```go
// 在controllers中设置断点
func GetUsers(c *gin.Context) {
    // 设置断点检查请求参数
    users := []models.User{}  // <- 断点位置
    // ... 业务逻辑
}
```

**Docker调试：**
```bash
# 启动Docker调试环境
make docker-debug

# 在Cursor中连接调试器，设置断点后：
# 使用curl或Postman测试API
curl http://localhost:9899/api/v1/users/
```

#### 2. 数据库查询调试

**本地调试：**
在数据库操作前后设置断点：
```go
// 在数据库初始化时调试
db, err := config.InitDB()
if err != nil {
    log.Fatalf("Failed to initialize database: %v", err)  // <- 断点位置
}
```

**Docker调试：**
```bash
# Docker环境中数据库连接调试
make docker-debug

# 查看容器内数据库连接状态
docker logs gin-debug | grep -i database
```

#### 3. 中间件调试

**本地调试：**
在中间件函数中设置断点：
```go
// 在middleware中调试请求处理
func GlobalErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 设置断点检查请求信息
        c.Next()  // <- 断点位置
    }
}
```

**Docker调试：**
```bash
# Docker环境中中间件调试
make docker-debug

# 监听容器日志查看中间件执行
docker logs gin-debug -f
```

#### 4. Docker特有调试场景

**容器环境变量调试：**
```bash
# 检查容器内环境变量
docker exec gin-debug env | grep GO_

# 检查配置文件加载
docker exec gin-debug ls -la config/
```

**网络连接调试：**
```bash
# 检查容器网络状态
docker inspect gin-debug | grep -A 10 NetworkMode

# 测试容器间通信
docker exec gin-debug ping db-dev
```

**文件挂载调试：**
```bash
# 检查代码挂载是否正常
docker exec gin-debug ls -la /app/

# 验证代码同步
echo "// test comment" >> controllers/user.go
docker exec gin-debug grep -n "test comment" /app/controllers/user.go
```

## 常见开发任务

#### 1. 运行测试

```bash
# 本地运行测试
go test ./...

# Docker 环境运行测试
make docker-test
```

#### 2. 代码格式化

```bash
# 格式化代码
go fmt ./...

# 整理导入
go mod tidy
```

#### 3. 构建项目

```bash
# 本地构建
go build -o website-gin cmd/main.go

# 跨平台构建（Linux）
GOOS=linux GOARCH=amd64 go build -o website-gin cmd/main.go
```

## Go版本升级指南

### 升级到最新稳定版本

项目当前使用 Go 1.25.0，如需升级到更新版本：

#### 1. 检查当前版本和可用版本

```bash
# 检查当前版本
go version
asdf current golang

# 查看可用版本
asdf list all golang | grep -E '^[0-9]+\.[0-9]+\.[0-9]+$' | tail -10
```

#### 2. 安装新版本

```bash
# 安装最新稳定版本（例如：1.25.0）
asdf install golang 1.25.0

# 更新项目版本配置
# 编辑 .tool-versions 文件
echo "golang 1.25.0" > .tool-versions

# 更新 go.mod 文件
sed -i '' 's/go [0-9]\+\.[0-9]\+\.[0-9]\+/go 1.25.0/' go.mod
```

#### 3. 更新依赖和工具

```bash
# 更新Go模块依赖
go mod tidy
go get -u ./...

# 重新安装Go工具
go install github.com/air-verse/air@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest

# 更新asdf shims
asdf reshim golang
```

#### 4. 验证升级

```bash
# 验证Go版本
go version
# 应显示新版本

# 验证工具版本
air -v
# 应显示：v1.62.0, built with Go go1.25.0

dlv version
# 应显示：Delve Debugger Version: 1.25.2

# 测试项目构建
go build -o /tmp/test-gin cmd/main.go && rm /tmp/test-gin
echo "✅ 构建成功！"
```

#### 5. 升级注意事项

**兼容性检查**：
- 检查第三方依赖是否支持新版本Go
- 运行完整测试套件确保功能正常
- 检查CI/CD配置中的Go版本

**常见问题**：
- 如果构建失败，检查 `go mod tidy` 输出
- 工具无法使用时，重新运行 `asdf reshim golang`
- 依赖冲突时，可能需要手动更新特定包版本

## 故障排除

### 常见问题与解决方案

#### 1. 配置文件找不到

**问题**：`pattern config-dev.yaml: no matching files found`

**解决方案**：
- 确保 `config/config-dev.yaml` 和 `config/config-prod.yaml` 文件存在
- 检查文件权限和路径

#### 2. 数据库连接失败

**问题**：`dial tcp: connect: connection refused`

**解决方案**：
- 确认数据库服务器正在运行
- 检查网络连接和防火墙设置
- 验证数据库配置信息（用户名、密码、主机、端口）

#### 3. 端口被占用

**问题**：`bind: address already in use`

**解决方案**：
- 检查端口占用：`lsof -i :9899`
- 修改配置文件中的端口号
- 或者停止占用端口的进程

#### 4. Go 版本不匹配

**问题**：`No version is set for command go`

**解决方案**：
```bash
# 安装指定版本
asdf install golang 1.22.3

# 确认项目目录包含 .tool-versions 文件
cat .tool-versions
# 应显示：golang 1.22.3
```

#### 5. Air 热重载工具找不到

**问题**：`No version is set for command air` 或 `command not found: air`

**原因**：使用 asdf 版本管理时，Go 工具的 shims 没有正确更新

**解决方案**：
```bash
# 方法 1：重新生成 asdf shims（推荐）
asdf reshim golang

# 方法 2：验证 air 安装
which air
# 应该显示：/Users/username/.asdf/shims/air

# 方法 3：重新安装 air（如果上述方法无效）
go install github.com/air-verse/air@latest
asdf reshim golang

# 方法 4：检查 Go 版本一致性
asdf current golang
# 确保显示项目目录的 .tool-versions 文件中指定的版本

# 方法 5：手动验证 air 是否可用
air -v
# 应该显示：v1.62.0, built with Go go1.25.0
```

**详细故障排除步骤**：
1. **检查 asdf 状态**：`asdf current golang`
2. **验证 air 位置**：`which air`
3. **重新生成 shims**：`asdf reshim golang` 
4. **测试 air 命令**：`air -v`
5. **如果仍有问题**：重新安装并重新生成 shims

#### 6. Delve调试器问题

**问题**：`The "dlv" command is not available`

**解决方案**：
```bash
# 方法 1：使用make命令自动安装（推荐）
make debug

# 方法 2：手动安装
go install -v github.com/go-delve/delve/cmd/dlv@latest
asdf reshim golang  # 如果使用asdf

# 验证安装
dlv version
# 应显示：Delve Debugger Version: 1.25.2 (latest)
```

#### 7. F5调试默认选择Django问题

**问题**：按F5时显示Django调试选项而不是Go

**原因**：多项目工作区中调试配置优先级问题

**解决方案**：
1. **在Go项目目录下调试**：打开 `website-gin` 文件夹后按F5
2. **检查项目级配置**：确认 `.vscode/launch.json` 存在
3. **重新加载工作区**：关闭并重新打开Cursor工作区

#### 8. 远程调试连接失败

**问题**：`Failed to connect to remote debugger`

**解决方案**：
```bash
# 确认调试服务器正在运行
make debug
# 等待看到：API server listening at: [::]:2345

# 检查端口占用
lsof -i :2345

# 如果端口被占用，停止进程或修改端口
pkill -f dlv
```

#### 9. 调试时变量显示不完整

**问题**：调试时复杂变量显示为 `<nil>` 或不完整

**解决方案**：
- 项目已优化Delve配置参数
- 支持更深层变量递归查看
- 如仍有问题，在 `.vscode/settings.json` 中调整：
```json
{
  "go.delveConfig": {
    "dlvLoadConfig": {
      "maxVariableRecurse": 5,
      "maxStringLen": 800
    }
  }
}
```

#### 10. Docker调试环境问题

**问题1**：`make docker-debug` 执行失败

**解决方案**：
```bash
# 检查Docker是否运行
docker info

# 检查网络是否存在
docker network ls | grep website-dev-network

# 如果网络不存在，创建网络
docker network create website-dev-network

# 清理可能冲突的容器
docker stop gin-debug 2>/dev/null || true
docker rm gin-debug 2>/dev/null || true

# 重新尝试
make docker-debug
```

**问题2**：Docker调试容器无法连接

**解决方案**：
```bash
# 检查容器状态
docker ps | grep gin-debug

# 查看容器日志
docker logs gin-debug

# 检查端口映射
docker port gin-debug

# 确认调试服务器是否启动
curl -s http://localhost:9899 || echo "应用未响应"
```

**问题3**：端口占用冲突

**解决方案**：
```bash
# 检查端口占用
lsof -i :9899
lsof -i :2345

# 停止占用端口的进程
make docker-stop-debug
make docker-stop

# 或者手动清理
docker stop $(docker ps -q --filter "publish=9899") 2>/dev/null || true
docker stop $(docker ps -q --filter "publish=2345") 2>/dev/null || true
```

#### 11. 热重载不工作

**问题**：文件修改后应用没有自动重启

**解决方案**：
```bash
# 检查 Air 配置文件
cat .air.toml

# 确认监听的文件类型和目录
# 检查 tmp 目录权限
ls -la tmp/

# 重新启动 Air
make watch
```

## 项目特性

### 1. 热重载支持

Docker 开发环境支持代码热重载，修改代码后应用会自动重启。

### 2. 数据库共享

与 Django 项目共享同一个 MySQL 数据库，确保数据一致性。

### 3. 环境配置管理

支持多环境配置，通过 `GO_ENV` 环境变量切换：
- `dev`：开发环境（默认）
- `prod`：生产环境

### 4. 中间件支持

项目包含全局异常处理中间件，确保错误统一处理。

## 部署注意事项

1. **生产环境**：记得设置 `GO_ENV=prod` 环境变量
2. **数据库**：确保生产环境数据库配置正确且安全
3. **端口**：确认生产环境端口配置不冲突
4. **日志**：生产环境建议配置日志轮转和持久化存储
