# ServerStatus

一个简单的服务器探针，Go实现，低占用



## 使用
### 服务端
配置文件`config.yml`
```
title: 服务器监控
port: 8080
token: passwd
nodes:
  - id: node1
  - id: node2
```
运行
```
status -s -f config.yml
```
### 客户端
配置文件`config.yml`
```
server_url: https://example.com
id: node1
token: passwd
```
运行
```
status -c -f config.yml
```

## 构建
### 拉取源码
```
git clone --recurse-submodules https://github.com/kzw200015/ServerStatus.git
```
### 构建前端
```
cd ServerStatus/assets/frontend
yarn && yarn build
```
### 打包构建前后端
```
cd ServerStatus
go build -o status
```