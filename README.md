# website-gin

拉取项目
```shell
git clone 

my-gin-app/
├── cmd/                         # 应用程序入口
│   └── server/
│       └── main.go              # 主程序入口
├── internal/                    # 私有代码，不能被外部导入
│   ├── handler/                 # HTTP 处理器层
│   │   ├── user_handler.go
│   │   ├── auth_handler.go
│   │   ├── product_handler.go
│   │   └── health_handler.go
│   ├── service/                 # 业务逻辑层
│   │   ├── interface/           # 服务接口定义
│   │   │   ├── user_service.go
│   │   │   ├── auth_service.go
│   │   │   └── product_service.go
│   │   ├── impl/                # 服务接口实现
│   │   │   ├── user_service_impl.go
│   │   │   ├── auth_service_impl.go
│   │   │   └── product_service_impl.go
│   │   └── service.go           # 服务管理器
│   ├── repository/              # 数据访问层
│   │   ├── interface/           # 仓储接口定义
│   │   │   ├── user_repository.go
│   │   │   └── product_repository.go
│   │   ├── impl/                # 仓储接口实现
│   │   │   ├── user_repository_impl.go
│   │   │   └── product_repository_impl.go
│   │   └── repository.go        # 仓储管理器
│   ├── model/                   # 数据库模型（实体）
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── order.go
│   │   └── base.go              # 公共字段
│   ├── dto/                     # 数据传输对象
│   │   ├── request/             # 请求 DTO
│   │   │   ├── user_request.go
│   │   │   ├── auth_request.go
│   │   │   ├── product_request.go
│   │   │   └── pagination_request.go
│   │   └── response/            # 响应 DTO
│   │       ├── user_response.go
│   │       ├── auth_response.go
│   │       ├── product_response.go
│   │       ├── pagination_response.go
│   │       └── common_response.go
│   ├── vo/                      # 值对象
│   │   ├── user_vo.go
│   │   ├── product_vo.go
│   │   ├── order_vo.go
│   │   └── password_vo.go
│   ├── middleware/              # 中间件
│   │   ├── auth_middleware.go
│   │   ├── cors_middleware.go
│   │   ├── logger_middleware.go
│   │   ├── rate_limit_middleware.go
│   │   ├── recovery_middleware.go
│   │   └── validator_middleware.go
│   ├── config/                  # 配置相关
│   │   ├── config.go
│   │   ├── database.go
│   │   ├── redis.go
│   │   └── jwt.go
│   └── wire/                    # 依赖注入（使用 Wire）
│       ├── wire.go
│       ├── wire_gen.go
│       └── provider.go
├── router/                      # 路由配置
│   ├── router.go                # 主路由配置
│   ├── api.go                   # API 路由组
│   ├── user_routes.go           # 用户相关路由
│   ├── auth_routes.go           # 认证相关路由
│   ├── product_routes.go        # 产品相关路由
│   └── middleware.go            # 路由级中间件
├── pkg/                         # 公共库，可被外部导入
│   ├── utils/                   # 工具函数
│   │   ├── jwt.go
│   │   ├── hash.go
│   │   ├── validator.go
│   │   ├── response.go
│   │   ├── pagination.go
│   │   └── file.go
│   ├── database/                # 数据库相关
│   │   ├── mysql.go
│   │   ├── redis.go
│   │   └── migration.go
│   ├── logger/                  # 日志相关
│   │   ├── logger.go
│   │   └── zap.go
│   ├── errors/                  # 错误处理
│   │   ├── errors.go
│   │   ├── codes.go
│   │   └── handler.go
│   └── constants/               # 常量定义
│       ├── constants.go
│       ├── status.go
│       └── message.go
├── api/                         # API 文档
│   └── docs/                    # Swagger 文档
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
├── web/                         # 静态文件（可选）
│   ├── static/
│   │   ├── css/
│   │   ├── js/
│   │   └── images/
│   └── templates/
│       ├── index.html
│       └── admin.html
├── configs/                     # 配置文件
│   ├── config.yaml              # 默认配置
│   ├── config.dev.yaml          # 开发环境配置
│   ├── config.test.yaml         # 测试环境配置
│   ├── config.prod.yaml         # 生产环境配置
│   └── locales/                 # 国际化文件
│       ├── en.yaml
│       └── zh.yaml
├── storage/                     # 存储目录
│   ├── logs/                    # 日志文件
│   │   ├── app.log
│   │   └── error.log
│   ├── uploads/                 # 上传文件
│   └── cache/                   # 缓存文件
├── scripts/                     # 脚本文件
│   ├── build.sh                 # 构建脚本
│   ├── deploy.sh                # 部署脚本
│   ├── migrate.sh               # 数据库迁移脚本
│   └── start.sh                 # 启动脚本
├── migrations/                  # 数据库迁移文件
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_products_table.up.sql
│   ├── 000002_create_products_table.down.sql
│   └── migrate.go               # 迁移管理程序
├── tests/                       # 测试文件
│   ├── unit/                    # 单元测试
│   │   ├── handler/
│   │   │   ├── user_handler_test.go
│   │   │   └── auth_handler_test.go
│   │   ├── service/
│   │   │   ├── user_service_test.go
│   │   │   └── auth_service_test.go
│   │   └── repository/
│   │       └── user_repository_test.go
│   ├── integration/             # 集成测试
│   │   ├── api/
│   │   │   ├── user_api_test.go
│   │   │   └── auth_api_test.go
│   │   └── database/
│   │       └── user_db_test.go
│   ├── fixtures/                # 测试数据
│   │   ├── users.json
│   │   └── products.json
│   ├── mocks/                   # Mock 文件
│   │   ├── user_service_mock.go
│   │   └── user_repository_mock.go
│   └── testdata/                # 测试资源文件
│       ├── test_config.yaml
│       └── test_data.sql
├── docker/                      # Docker 相关文件
│   ├── Dockerfile               # 应用 Dockerfile
│   ├── Dockerfile.dev           # 开发环境 Dockerfile  
│   ├── docker-compose.yml       # Docker Compose 配置
│   ├── docker-compose.dev.yml   # 开发环境 Docker Compose
│   └── nginx/                   # Nginx 配置
│       ├── nginx.conf
│       └── default.conf
├── deployments/                 # 部署配置
│   ├── kubernetes/              # K8s 部署文件
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   ├── configmap.yaml
│   │   └── ingress.yaml
│   ├── helm/                    # Helm Charts
│   │   ├── Chart.yaml
│   │   ├── values.yaml
│   │   └── templates/
│   └── docker-swarm/            # Docker Swarm 配置
│       └── stack.yml
├── docs/                        # 项目文档
│   ├── README.md
│   ├── API.md                   # API 文档
│   ├── DEPLOYMENT.md            # 部署文档
│   ├── DEVELOPMENT.md           # 开发文档
│   └── CHANGELOG.md             # 变更日志
├── tools/                       # 开发工具
│   ├── gen/                     # 代码生成工具
│   │   ├── model.go
│   │   └── api.go
│   └── lint/                    # 代码检查工具
│       └── golangci.yml
├── .github/                     # GitHub 相关
│   └── workflows/               # GitHub Actions
│       ├── ci.yml
│       ├── cd.yml
│       └── test.yml
├── .gitignore                   # Git 忽略文件
├── .env                         # 环境变量文件
├── .env.example                 # 环境变量示例
├── .air.toml                    # Air 热重载配置
├── go.mod                       # Go 模块文件
├── go.sum                       # Go 模块校验文件
├── Makefile                     # Make 构建文件
├── README.md                    # 项目说明
├── LICENSE                      # 许可证文件
└── CHANGELOG.md                 # 变更日志


```


安装asdf
```shell
brew install asdf
```

在 macOS 上通过 asdf 使用 Go 多版本管理 非常方便。以下是完整、实用的指南，适合你在开发中快速切换 Go 版本。

🧰 一、前置准备
确保已经安装好 asdf：

bash
复制
编辑
brew install asdf
echo -e '\n. $(brew --prefix asdf)/libexec/asdf.sh' >> ~/.zshrc
source ~/.zshrc
🪛 二、添加 Go 插件
bash
复制
编辑
asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
⬇️ 三、安装多个 Go 版本
bash
复制
编辑
asdf install golang 1.22.1
asdf install golang 1.21.7
安装完后，可以查看已安装版本：

bash
复制
编辑
asdf list golang
⚙️ 四、切换版本
1. 设置全局 Go 版本（系统默认）
   bash
   复制
   编辑
   asdf global golang 1.22.1
2. 设置项目本地 Go 版本（目录默认）
   进入你的项目目录，然后：

bash
复制
编辑
asdf local golang 1.21.7
会生成 .tool-versions 文件：

txt
复制
编辑
golang 1.21.7
只要你进入该目录，asdf 会自动切换到这个版本。

✅ 五、验证是否生效
bash
复制
编辑
go version
输出应该是你刚设置的版本，如：

bash
复制
编辑
go version go1.21.7 darwin/arm64
🛠️ 六、管理多个版本的小技巧
1. 卸载某个版本
   bash
   复制
   编辑
   asdf uninstall golang 1.21.7
2. 查看当前激活版本
   bash
   复制
   编辑
   asdf current golang
3. 使用 .tool-versions 快速复用配置
   将其复制到其他项目目录，即可复用相同版本。

🚀 七、Go 项目建议
如果你在不同项目中使用不同版本的 Go：

在每个项目根目录使用 asdf local golang <version>

把 .tool-versions 添加到版本控制中

配合 go.mod 可以更方便地维护依赖和构建兼容性





