# IceFireDB-SQLite

IceFireDB SQLite数据库是一种去中心化SQLite数据库。提供一套方便的机制构建全球分布式数据库系统。 支持用户以MySQL协议向IceFireDB进行数据写入。 IceFireDB将数据存储到SQLite数据库，并将数据在P2P自动组网的节点间同步。

## 它是如何工作的?
### 场景介绍
![scene.png](./docs/scene.png)
### 系统架构
![framework](./docs/framework.png)

## 入门

### 安装
```shell
git clone https://github.com/IceFireDB/IceFireDB-SQLite
cd IceFireDB-SQLite
make
```

### 快速开始

https://user-images.githubusercontent.com/21053373/173170247-74b1daeb-7bd5-4dc0-8b93-62b334859ba8.mp4


### 用法
#### 配置文件
```yaml
# mysql-server 配置
server:
  addr: ":23306" # 监听的端口，支持mysql-client直连接

sqlite:
  filename: "sqlite.db"

debug:  # 控制开启debug模式
  enable: true
  port: 17878

# p2p config
p2p:
  enable: true
  # 自定义发现id, topic
  service_discovery_id: "p2p_sqlite_service_test"
  service_command_topic: "p2p_sqlite_service_topic_test"
  service_discover_mode: "advertise" # advertise or announce

# mysql认证用户列表
userlist:
  - user: host1
    password: host1
  - user: test
    password: test
```

#### 连接
支持任意语言通过标准mysql驱动连接使用，命令行mysql使用示例
```shell
$ mysql -h 127.0.0.1 -P 23306 -u host1 -phost1
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10001
Server version: 5.7.0

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> select * from maps;
Empty set (0.00 sec)

mysql>
```

## License
IceFireDB SQLite使用Apache 2.0许可证。请参见 [LICENSE](./LICENSE) directory for details.

## 免责声明
当您使用本软件时，您已同意并声明本软件的作者、维护者和贡献者不对您遇到的任何风险、成本或问题负责。如果发现软件缺陷或BUG，请提交补丁以帮助改进！