# 使用官方 Go 镜像作为构建基础
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装 git（GoFrame 需要 git 来下载依赖）
RUN apk add --no-cache git

# 下载 GoFrame 工具
RUN go install github.com/gogf/gf/v2@latest

# 创建最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将 GoFrame 工具从构建阶段复制到最终镜像中
COPY --from=builder /go/bin/gf /usr/local/bin/gf

# 设置环境变量（可选）
ENV GF_HOME=/root/.gf

# 启动容器时的默认命令（可以根据需要修改）
CMD ["/bin/sh"]