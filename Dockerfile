FROM golang:1.10-alpine3.7

MAINTAINER 18852852150@163.com

# 项目名 && 模块名
ENV PROJECT myservices
ENV MODULE books

# workspace 工作目录
ENV WORKSPACE /go/src/${PROJECT}/${MODULE}

# build main.go所在的目录
ENV BUILDPATH ${WORKSPACE}/cmd/${MODULE}-server/main.go

# 进入workspace 并copy代码进入容器
WORKDIR ${WORKSPACE}
COPY . .

# 编译
RUN go build -o ${MODULE} ${BUILDPATH}

# 端口
EXPOSE 5001
EXPOSE 5002
EXPOSE 5003
EXPOSE 5004
EXPOSE 5005

# 启动服务
ENTRYPOINT ./$MODULE
