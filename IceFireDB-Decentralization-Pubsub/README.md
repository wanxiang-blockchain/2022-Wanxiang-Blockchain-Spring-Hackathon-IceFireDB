# [IceFireDB-Decentralization-Pubsub]

IceFireDB Decentralization Pubsub是一个高性能、高可用性和去中心化订阅系统。

它可以无缝地将使用redis发布和订阅的web2应用程序迁移到去中心化p2p订阅网络中

## How does it work?

### 场景介绍
<img width="1070" alt="image" src="https://user-images.githubusercontent.com/52234994/173174080-95775add-c287-4f43-943b-686b903e67b7.png">

### 系统架构
<img width="1061" alt="image" src="https://user-images.githubusercontent.com/52234994/173174090-7c2f084b-4ccc-4c56-983b-210486f13a60.png">

该应用程序适用于同一网络或不同网络上的多个节点。专用网络上NAT后面的节点可以相互通信。使用Kademlia DHT和IPFs网络发现对等发现和路由。通过支持redis pub-sub协议，构建了一个全局分布式Web3发布-订阅系统。

您可以像使用redis发布订阅功能一样使用它。

## Getting Started

### Configuration

在config目录中，用户存放项目配置文件，文件名：config.yaml，可根据自身需要进行修改

```yaml
# 项目代理配置
proxy:
  local_port: 16379
  enable_mtls: false

# p2p 配置
p2p:
  enable: true

...
```

### Quickstart

https://user-images.githubusercontent.com/52234994/173171008-8c73ce17-4ba7-42ec-8257-025e98d2e647.mp4

直接运行二进制文件，如果需要在后台运行可以添加到systemd系统管理中
```shell
$ make
$ ./bin/Icefiredb-proxy -c ./config/config.yaml
```

### Usage
IceFireDB 订阅系统主要用法为两个命令：SUBSCRIBE 和 PUBLISH，主要实现在[pubsub](./pkg/router/redisNode/ppubsub.go)

可以根据需求，进行二次开发，或者增加其他指令。

SUBSCRIBE
```shell
$ redis-cli
127.0.0.1:16379> SUBSCRIBE name
...
...
```
PUBLISH
```shell
$ redis-cli
127.0.0.1:16379> PUBLISH name hello
...
...
```

## License
Icefiredb代理使用Apache 2.0许可证。请参见 [LICENSE](./LICENSE) directory for details.

## 免责声明
当您使用本软件时，您已同意并声明本软件的作者、维护者和贡献者不对您遇到的任何风险、成本或问题负责。如果发现软件缺陷或BUG，请提交补丁以帮助改进！
