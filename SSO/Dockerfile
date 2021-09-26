FROM golang:latest AS builder

WORKDIR /app/

COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY="https://goproxy.cn,direct" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main main.go

FROM xylonx/cn-ubuntu:latest AS prod

LABEL maintainer="xylonx"

ARG PROJECT_NAME="uniquesso-backend"
WORKDIR /opt/${PROJECT_NAME}

COPY --from=0 /app/main ./${PROJECT_NAME}
RUN mkdir config && \
    echo "./${PROJECT_NAME} server -c=settings.yaml" >>start.sh && \
    chmod 755 -R /opt/$PROJECT_NAME

CMD ./start.sh

EXPOSE 5000