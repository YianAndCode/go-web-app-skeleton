FROM golang:1.14

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
