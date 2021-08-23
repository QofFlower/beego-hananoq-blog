# beego-hananoq-blog

#### 介绍
参考答案博客java后端，本项目使用的是golang的beego框架实现后端业务逻辑

#### 软件架构
软件架构说明

|    框架    |  版本  |
| :--------: | :----: |
|   Beego    | 1.12.1 |
| PostgreSQL |   13   |
|   Oauth2   | 3.12.0 |
|   Redis    | 3.2.1  |
|   Nacos    | 1.4.0  |
由于我莫得money，服务器太拉跨，跑不起ElasticSearch，所以就不用了

#### 安装教程

1.  `git clone ...`
2.  `go get` / `go mod tidy` 安装依赖
3.  `bee run`

#### 使用说明

1. 安装并配置好redis，用于Oauth2 token信息的存储

2. OSS为阿里云对象存储，也可以改为本机存储

3. 放行的请求路由在配置文件中release_router配置，也可以直接插入过滤器

4. 我使用的postgreSQL是部署在docker里的，docker容器postgres出现远程链接错误问题，将容器里的`/var/lib/postgresql/data/pg_hba.conf`最后添加

   允许IPV4的所有地址链接即可

   ```properties
   # TYPE  DATABASE        USER            ADDRESS                 METHOD
   
   # "local" is for Unix domain socket connections only
   local   all             all                                     trust
   # IPv4 local connections:
   host    all             all             127.0.0.1/32            trust
   host    all             all             0.0.0.0/0               trust
   # IPv6 local connections:
   host    all             all             ::1/128                 trust
   # Allow replication connections from localhost, by a user with the
   # replication privilege.
   local   replication     all                                     trust
   host    replication     all             127.0.0.1/32            trust
   host    replication     all             ::1/128                 trust
   ```

   使用挂载和进入容器修改均可

5.这里配置文件默认加载的是nacos内的配置文件，可以修改为本地的config_xxx.yaml文件