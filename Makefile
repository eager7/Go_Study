TARGET:=$(shell pwd | sed 's/\/.*.\///g')

default:all

all:
	export GOPROXY="https://goproxy.io" && GO111MODULE=on go build main.go

install:
	mv src/$(TARGET) ./

clean:
	rm $(TARGET)