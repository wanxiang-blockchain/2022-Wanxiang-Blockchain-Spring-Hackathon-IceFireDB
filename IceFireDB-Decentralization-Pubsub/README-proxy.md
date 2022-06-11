![banner](./docs/logo.png)
# IceFireDB-Proxy
[![Build](https://github.com/IceFireDB/IceFireDB-Proxy/actions/workflows/main.yml/badge.svg)](https://github.com/IceFireDB/IceFireDB-Proxy/actions/workflows/main.yml) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IceFireDB/IceFireDB-Proxy)](https://github.com/IceFireDB/IceFireDB-Proxy/blob/master/go.mod)

IceFireDB-Proxy is a high-performance, high-availability, and user-friendly Resp protocol cluster proxy solution. It is supporting P2P networking and is a network component in the IceFireDB ecosystem.

1. Complete data source mode support: stand-alone, cluster mode
2. Rich command support
3. Excellent cluster state management and failover
4. Excellent traffic control policies: Traffic read/write separation and multi-tenant data isolation
5. Excellent command telemetry features
6. Bottom-fishing use of mind and base abilities that are closer to cloud native
7. Supports P2P automatic networking, and Proxy helps traditional Redis databases achieve data decentralization.
8. New framework for faster network, will be upgraded soon. [redhub](https://github.com/IceFireDB/redhub)

## How does it work?

![comp-archotecture](./docs/comp-archotecture.png)

![usage-architecture](./docs/usage-architecture.png)

## Getting Started

### Installation

This project uses Go. Go check them out if you don't have them locally installed.
```text
1. Install Go
2. git clone https://github.com/IceFireDB/IceFireDB-Proxy.git $GOPATH/src/github.com/IceFireDB/IceFireDB-Proxy
3. cd $GOPATH/src/github.com/IceFireDB/IceFireDB-Proxy
4. make
```

### Quickstart

Run a binary file directly, if you need to run in the background can be added to the systemd system management
```shell
./bin/Icefiredb-proxy -c ./config/config.yaml
```

### Usage

Command support
```shell
$ redis-cli
127.0.0.1:16379> SUBSCRIBE name
...
...
```

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

## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## License
Icefiredb-proxy is under the Apache 2.0 license. See the [LICENSE](./LICENSE) directory for details.

## Disclaimers
When you use this software, you have agreed and stated that the author, maintainer and contributor of this software are not responsible for any risks, costs or problems you encounter. If you find a software defect or BUG, ​​please submit a patch to help improve it!
