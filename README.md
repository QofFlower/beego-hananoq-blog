# beego-hananoq-blog

#### 介绍
参考答案博客java后端，本项目使用的是golang的beego框架实现后端业务逻辑

#### 软件架构
软件架构说明

|    框架    |  版本  |
| :--------: | :----: |
|   Beego    | 1.12.1 |
| PostgreSQL |   12   |
|   Oauth2   | 3.12.0 |
|   Redis    | 3.2.1  |


#### 安装教程

1.  `git clone ...`
2.  `go get` / `go mod tidy` 安装依赖
3.  `bee run`

#### 使用说明

1.  安装并配置好redis，用于Oauth2 token信息的存储
2.  OSS为阿里云对象存储，也可以改为本机存储
3.  放行的请求路由在配置文件中release_router配置，也可以直接插入过滤器

