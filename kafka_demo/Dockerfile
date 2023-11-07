FROM golang:1.20 as build

WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -o app main.go

FROM ubuntu:latest as prod
WORKDIR /app
COPY --from=build /app/app .
ENTRYPOINT /app/app