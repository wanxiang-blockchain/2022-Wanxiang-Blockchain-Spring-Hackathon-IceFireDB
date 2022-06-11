![banner](./docs/logo.png)
# IceFireDB-Proxy
[![Build](https://github.com/IceFireDB/IceFireDB-Proxy/actions/workflows/main.yml/badge.svg)](https://github.com/IceFireDB/IceFireDB-Proxy/actions/workflows/main.yml) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IceFireDB/IceFireDB-Proxy)](https://github.com/IceFireDB/IceFireDB-Proxy/blob/master/go.mod)

IceFireDB Redis Proxy是一种高性能、高可用性和用户友好的Resp协议群集代理解决方案。它支持P2P网络，是IceFireDB生态系统中的一个网络组件。
1. 完整的数据源模式支持：单机、集群模式
2. 丰富的命令支持
3. 出色的群集状态管理和故障切换
4. 优秀的流量控制策略：流量读写分离和多租户数据隔离
5. 出色的命令遥测功能
6. 使用接近云原生7的思维和基础能力打底。支持P2P自动组网，代理帮助传统Redis数据库实现数据分散。

## How does it work?

![comp-archotecture](./docs/comp-archotecture.png)

![usage-architecture](./docs/usage-architecture.png)


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
  service_discovery_id: "p2p_redis_proxy_service_test"
  # pubsub topic
  service_command_topic: "p2p_redis_proxy_service_topic_test"
  # 节点发现模式
  service_discover_mode: "advertise" # advertise or announce
# 代理redis配置
redisdb:
# prometheus metrics 配置
prometheus_exporter:
...
```

### Quickstart

https://user-images.githubusercontent.com/52234994/173170991-08713e52-291c-4fae-bf46-ce87b959ce90.mp4

直接运行二进制文件，如果需要在后台运行可以添加到systemd系统管理中
```shell
$ make
$ ./bin/Icefiredb-proxy -c ./config/config.yaml
```

### Usage

#### String
* APPEND
* BITCOUNT
* BITPOS
* DECR
* DECRBY
* DEL
* EXISTS
* GET
* GETBIT
* SETBIT
* GETRANGE
* GETSET
* INCR
* INCRBY
* MGET
* MSET
* SET
* SETEX
* SETEXAT
* SETRANGE
* EXPIRE
* EXPIREAT
* TTL


#### Set
* SADD
* SCARD
* SETBIT
* SISMEMBER
* SMEMBERS
* SPOP
* SRANDMEMBER
* SREM
* SSCAN

#### List
* LINDEX
* LINSERT
* LLEN
* LPOP
* LPUSH
* LPUSHX
* LRANGE
* LREM
* LSET
* LTRIM
* RPOP
* RPUSH
* RPUSHX

#### Hash
* HDEL
* HEXISTS
* HGET
* HGETALL
* HINCRBY
* HINCRBYFLOAT
* HKEYS
* HLEN
* HMGET
* HMSET
* HSCAN
* HSET
* HSETNX
* HSTRLEN
* HVALS

#### Sorted Sets
* ZADD
* ZCARD
* ZCOUNT
* ZINCRBY
* ZLEXCOUNT
* ZPOPMAX
* ZPOPMIN
* ZLEXCOUNT
* ZRANGE
* ZRANGEBYLEX
* ZRANGEBYSCORE
* ZRANK
* ZREM
* ZREMRANGEBYLEX
* ZREMRANGEBYRANK
* ZREMRANGEBYSCORE
* ZREVRANGE
* ZREVRANGEBYLEX
* ZREVRANGEBYSCORE
* ZREVRANK
* ZSCAN
* ZSCORE

#### Stream
* XACK
* XADD
* XCLAIM
* XDEL
* XLEN
* XINFO
* XPENDING
* XRANGE
* XREADGROUP
* XREVRANGE
* XTRIM
* XGROUP


#### Others

* COMMAND
* PING
* QUIT

## License
Icefiredb代理使用Apache 2.0许可证。有关详细信息，请参阅[LICENSE](.LICENSE)。

## 免责声明
当您使用本软件时，您已同意并声明本软件的作者、维护者和贡献者不对您遇到的任何风险、成本或问题负责。如果发现软件缺陷或BUG，请提交补丁以帮助改进！