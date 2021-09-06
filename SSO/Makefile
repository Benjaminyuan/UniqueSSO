PROJECT:=unique-sso
CONFIG_FILE:=./conf.yaml
GO_MODULE_STATE:=$(shell go env GO111MODULE)
GO_PROXY:=$(shell go env GOPROXY)

.PHONY: build dep serve clean

.DEFAULT: ${PROJECT}

${PROJECT}: build

build: dep
	CGO_ENABLED=0 go build -o ${PROJECT}

dep:
ifneq ($(GO_MODULE_STATE), on)
	go env -w GO111MODULE="on"
endif
ifeq ($(GO_PROXY), https://proxy.golang.org,direct)
	go env -w GOPROXY="https://goproxy.cn,direct"
endif
	go mod tidy

serve: ${PROJECT}
	./${PROJECT} -c ./${CONFIG_FILE} 

clean:
	rm -rf ${PROJECT}