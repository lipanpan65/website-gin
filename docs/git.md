

重构分支名称：
```shell
git checkout -b refactor/complete-rewrite develop
```


# 从主分支（如 main 或 master）创建新分支
git checkout -b refactor/directory-structure

# 重构完成后，将功能分支合并到 develop 分支
git checkout develop
git merge refactor/directory-structure


```
your-project/
├── app/                # 应用核心逻辑
│   ├── config/         # 配置管理（如环境变量、数据库配置）
│   ├── middleware/     # 中间件
│   ├── router/         # 路由定义
│   ├── service/        # 业务逻辑层
│   ├── model/          # 数据模型（结构体、数据库映射）
│   └── controller/     # 控制器（处理请求参数、调用服务层）
├── pkg/                # 自定义工具包（可复用的模块，如日志、数据库连接、JWT 等）
├── cmd/                # 入口文件（main 函数，启动服务）
├── internal/           # 内部工具或私有代码（不建议对外暴露）
├── docs/               # 文档（如 Swagger 接口文档）
├── test/               # 测试文件
├── go.mod              # Go 模块依赖管理
└── main.go             # 主入口（可移至 cmd 目录下）
```
