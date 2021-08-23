FROM golang:latest

# golang官方镜像的gopath为/go/src
RUN mkdir -p /go/src/beego-hananoq-blog
# 进入工作目录
WORKDIR /go/src/beego-hananoq-blog
# 将工程全部复制到工作目录
COPY . /go/src/beego-hananoq-blog
# 使用代理，解决go get下载超时或很慢
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN export GOPROXY=https://goproxy.cn
# 下载依赖
# 官方说go get -d 可以自动下载所有命令？待试试，还有go install都是待测, 下次build时候可以试试
RUN go get github.com/astaxie/beego && \
go get github.com/beego/bee && \
go get github.com/garyburd/redigo && \
go get github.com/go-redis/redis && \
go get github.com/go-redis/redis/v8 && \
go get github.com/gorilla/sessions && \
go get github.com/lib/pq && \
go get github.com/nacos-group/nacos-sdk-go && \
go get github.com/parkingwang/go-sign && \
go get github.com/shiena/ansicolor && \
go get github.com/smartystreets/goconvey && \
go get github.com/ulule/limiter/v3 && \
go get gopkg.in/boj/redistore.v1 && \
go get gopkg.in/go-oauth2/redis.v3 && \
go get gopkg.in/oauth2.v3 && \
go get gopkg.in/yaml.v2
EXPOSE 9696

CMD ["bee", "run"]