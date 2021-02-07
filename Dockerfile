FROM golang:1.14.3 AS builder
LABEL maintainer="Cherisher cool0616@163.com"

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /data

COPY . .

RUN go build -tags='srs' -ldflags '-s -w -linkmode "external" -extldflags "-static" ' -v -o gb28181-proxy-srs

FROM centos:7.4

WORKDIR /data
COPY --from=builder /data/conf/app.json conf/
COPY --from=builder /data/gb28181-proxy-srs .

EXPOSE 28080

ENTRYPOINT ["./gb28181-proxy-srs"]
#CMD ["-c", "conf/app.json"]
