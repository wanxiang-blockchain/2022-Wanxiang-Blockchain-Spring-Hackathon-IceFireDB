# [IceFireDB-Decentralization-Pubsub]
[![Build](https://github.com/IceFireDB/IceFireDB-Proxy/actions/workflows/main.yml/badge.svg)](https://github.com/IceFireDB/IceFireDB-Proxy/actions/workflows/main.yml) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IceFireDB/IceFireDB-Proxy)](https://github.com/IceFireDB/IceFireDB-Proxy/blob/master/go.mod)    

IceFireDB-Decentralization-Pubsub is a high-performance, highly available and decentralized subscription system.

It can seamlessly migrate applications that publish and subscribe using redis to the icefiredb subscription system

## How does it work?

The application is applicable to multiple nodes on the same network or different networks. The nodes behind NAT on the private network can communicate with each other. Kademlia DHT and IPFs networks are used for peer-to-peer discovery and routing. By supporting the redis publish subscribe protocol, a global distributed Web3 publish subscribe system is built.

You can use it just as you use the redis publish subscribe function.


## Getting Started

### Installation

This project uses Go. Go check them out if you don't have them locally installed.
```text
1. Install Go
2. git clone https://github.com/IceFireDB/IceFireDB-Decentralization-Pubsub.git $GOPATH/src/github.com/IceFireDB/IceFireDB-Decentralization-Pubsub
3. cd $GOPATH/src/github.com/IceFireDB/IceFireDB-Decentralization-Pubsub
4. make
```
### Quickstart

Run a binary file directly, if you need to run in the background can be added to the systemd system management
```shell
./bin/Icefiredb-proxy -c ./config/config.yaml
```

### Usage
Here is our video teaching：[video](./demo.mp4)

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

## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## License
Icefiredb-proxy is under the Apache 2.0 license. See the [LICENSE](./LICENSE) directory for details.

## Disclaimers
When you use this software, you have agreed and stated that the author, maintainer and contributor of this software are not responsible for any risks, costs or problems you encounter. If you find a software defect or BUG, ​​please submit a patch to help improve it!
