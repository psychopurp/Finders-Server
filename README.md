# 项目文档

## 基本介绍

> 兴趣社交平台 Finders 的服务端

## 技术选型

- 后台管理系统：用基于`vue`的`Element-UI`构建基础页面。
- 后端：用`Gin`快速搭建基础 RESTFUL 风格 API，`Gin`是一个 go 语言编写的 Web 框架。
- 数据库：采用`MySql`，使用`gorm`实现对数据库的基本操作,已添加对 sqlite 数据库的支持。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API 文档：使用`Swagger`构建自动化文档。
- 配置文件：使用`fsnotify`和`viper`实现`yaml`格式的配置文件。
- 日志：使用`go-logging`实现日志记录。
- 热重载：使用 fresh 来进行 web 热重启，实现快速调试。

## 运行说明

### 运行环境

- golang 版本 > 1.14
- IDE 推荐 > vscode

### 运行 Server

```bash
git clone git@github.com:psychopurp/Finders-Server.git

cd Finders-Server
#此目录下有go.mod
# 安装go依赖包
go list (go mod tidy)
# 编译
go run main.go

```

## Swagger 自动化 API 文档

### 安装 Swagger

```bash
go get -u github.com/swaggo/swag/cmd/swag
cd Finders-Server
swag init
#执行上面的命令后，FInders-Server目录下会出现docs文件夹，登录http://localhost:port/swagger/index.html，即可查看swagger文档
```

## fresh 进行后端热重启 (避免每次更改又得重新手动编译)

### 安装 fresh

```bash
go get github.com/pilu/fresh
cd Finders-Server
fresh
#执行上面的命令后， fresh 将会自动运行项目的 main.go
```

## 项目架构

```
Finders-Server        (后端文件夹)
    │─api             (API)
    │─config          (配置包)
    │─core  	      (內核)
    │─docs  	      (swagger文档目录)
    │─global          (全局对象)
    │─initialiaze     (初始化)
    │─middleware      (中间件)
    │─model           (结构体层)
    │─resource        (资源)
    │─router          (路由)
    │─service         (服务)
    │─utils           (公共功能)


```
