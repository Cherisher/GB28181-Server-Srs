FROM golang:alpine as builder
LABEL maintainer="Cherisher cool0616@163.com"

ENV GOPROXY https://goproxy.cn,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /data

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache gcc musl-dev
RUN go env && go list && go build -v -a -mod=vendor -tags='srs' -ldflags "-extldflags \"-static\" "  -o gb28181-proxy-srs .

#RUN go build -tags='srs' -ldflags '-s -w -linkmode "external" -extldflags "-static" ' -v -o gb28181-proxy-srs .

FROM nginx:alpine
LABEL maintainer="Cherisher cool0616@163.com"

RUN apk add --no-cache  gettext tzdata   && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone && \
    date && \
    apk del tzdata

WORKDIR /data


COPY docker-start.sh ./
COPY nginx.conf.tpl nginx.conf.tpl

COPY --from=builder /data/view /var/www/
COPY --from=builder /data/conf/app.json conf/
COPY --from=builder /data/gb28181-proxy-srs .


ENV API_SERVER="http://127.0.0.1:8888"

EXPOSE 80

ENTRYPOINT ["./docker-start.sh"]

