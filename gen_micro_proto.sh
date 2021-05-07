#!/bin/bash

#go get -u -v github.com/micro/go-micro/v2
#go get -v github.com/micro/micro #此处需要go 1.12版本，或者不需要这个工具，纯手工写;v3版本的开源协议存在争议

#go get -u -v github.com/micro/protoc-gen-micro/v2
#go get -u -v github.com/golang/protobuf

#brew install protobuf //安装protoc工具

protoc --go_out=./ --micro_out=./ proto/*.proto

