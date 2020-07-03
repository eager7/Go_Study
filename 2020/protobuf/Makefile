# Makefile for protobuf, if protoc: Command not found, then install by: sudo apt-get install protobuf-compiler
# 参考文档: https://blog.csdn.net/zhonglinzhang/article/details/53118854 https://segmentfault.com/a/1190000020393528
# export GOPROXY=https://goproxy.io && GO111MODULE=on go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# export GOPROXY=https://goproxy.io && GO111MODULE=on go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
# export GOPROXY=https://goproxy.io && GO111MODULE=on go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger


default: all
SRC_ROOT := $(shell pwd)/
CSOURCE += $(wildcard $(SRC_ROOT)/*.proto)

COBJS := $(patsubst %.proto,%.proto,$(notdir $(CSOURCE)))

.PHONY: core api # 强制编译
all: core

core:
	protoc -I=. -I=$(GOPATH)/src --go_out=plugins=grpc,paths=source_relative:. $(COBJS)

.PHONY: clean
clean:
	-rm *.pb.go*

test:
	@echo $(CSOURCE)
	@echo $(COBJS)