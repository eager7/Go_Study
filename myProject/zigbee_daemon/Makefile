TARGET:=$(shell pwd | sed 's/\/.*.\///g')

export GOPATH=$(shell pwd)

default:all install

all:
	echo $(GOPATH)
	cd src;go build $(TARGET).go

install:
	mv src/$(TARGET) ./

clean:
	rm $(TARGET)