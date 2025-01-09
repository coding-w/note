# Kafka

## 名词介绍

![](../images/kafka/Kafka.png)

- Topic: 消息队列，生产者和消费者面向的都是一个Topic
- Broker: 一个Kafka服务器就是一个Broker，一个集群由多个Broker组成。一个Broker可以容纳多个Topic
- Producer: 消息生产者，向Kafka Broker发生消息的客户端
- Consumer: 消息消费者，向Kafka Broker取消息的客户端
- Consumer Group(CG): 消费者组，有多个Consumer组成。消费者组内每个消费者负责不同分区的数据，一个分区只能有一个组内消费者消费；消费者组之间互不影响。所有消费者都属于某一个消费者组
- Partition: 为了实现扩展性，一个非常大的Topic可以分布到多个Broker上，一个topic可以分为多个partition
- Replica: 副本，每一个分区都有若干个副本，保证系统的稳定性
- Leader: 分区副本的“领导者”，生产者发送数据的对象，消费者消费数据的对象
- Follower: 分区副本的“从”，保持和Leader数据的同步。Leader发送故障，会从Follower中选举新的Leader

## Kafka集群安装

使用docker对Kafka进行集群安装，[docker-compose.yaml](./cluster/docker-compose.yaml)

```shell
# 切换到目标目录
# 目录结构
#.
#├── docker-compose.yaml
#│
#├── kafka_data1
#│
#├── kafka_data2
#│
#├── kafka_data3
#│
#├── zookeeper_data1
#│
#├── zookeeper_data2
#│
#└── zookeeper_data3
# 启动
docker-compose up -d
```

### 配置
详情见[server.properties](./singleton/server.poperties)配置

## 操作命令

进入容器`docker exec -it kafka1 /bin/bash`

### Topic

可以使用`kafka-topics.sh`查看操作主题的命令，常用命令如下

| 参数                                                   | 描述                         |
|------------------------------------------------------|----------------------------|
| `--bootstrap-server <String: server to connect to>`  | 连接的 Kafka Broker 主机名称和端口号。 |
| `--topic <String: topic>`                            | 操作的主题名称。                   |
| `--create`                                           | 创建主题。                      |
| `--delete`                                           | 删除主题。                      |
| `--alter`                                            | 修改主题配置。                    |
| `--list`                                             | 列出所有主题。                    |
| `--describe`                                         | 查看指定主题的详细信息。               |
| `--partitions <Integer: # of partitions>`            | 设置主题的分区数量（仅在创建主题时使用）。      |
| `--replication-factor <Integer: replication factor>` | 设置分区的副本因子（仅在创建主题时使用）。      |
| `--config <String: name=value>`                      | 更新或设置主题的配置参数。              |

```shell
# 创建一个主题
kafka-topics.sh --bootstrap-server kafka1:9093 --create --partitions 1 --replication-factor 3 --topic first
## --partitions 分区数
## --replication-factor 副本数
## --topic 主题名称

# 查看主题列表
kafka-topics.sh --bootstrap-server kafka1:9093 --list

# 查看主题详情
kafka-topics.sh --bootstrap-server kafka1:9093 --describe --topic first
Topic: first	TopicId: nmM8zHmmRsqm5WjhTXuYHw	PartitionCount: 1	ReplicationFactor: 3	Configs: 
	Topic: first	Partition: 0	Leader: 2	Replicas: 2,3,1	Isr: 2,3,1	Elr: N/A	LastKnownElr: N/A
## TopicId Topic 唯一标识符（UUID）
## Partition: 0 Topic 中的第一个分区（编号从 0 开始） 
## Leader: 2 分区 0 的 Leader 是 Broker 2

# 修改分区数，分区数只能增加，不能减少
kafka-topics.sh --bootstrap-server kafka1:9093 --alter --topic first --partitions 3
# 删除分区
kafka-topics.sh --bootstrap-server kafka1:9093 --topic first --delete
```

### Producer

可以使用`kafka-console-producer.sh`查看生产者的命令，常用命令如下

| 参数                                                   | 描述                         |
|------------------------------------------------------|----------------------------|
| `--bootstrap-server <String: server to connect to>`  | 连接的 Kafka Broker 主机名称和端口号。 |
| `--topic <String: topic>`                            | 操作的主题名称。                   |

```shell
# 生产消息
kafka-console-producer.sh --bootstrap-server kafka2:9094 --topic first
```

#### 发送消息流程

[生产者代码](./src/main/java/com/coding/w/producer/MyProducer.java)

1. `kafkaProducer`需要配置一些基本的参数，比如 Kafka 集群的地址、序列化方式、重试策略等
2. 发送消息，当`batch.size=16k || linger.ms = 0ms`时会调用`send()`方法将消息发送到指定的 topic 和 partition
3. 发送请求到`Broker Kafka`生产者将序列化后的消息发送给目标broker。生产者与Kafka broker之间的通信通常基于TCP协议，生产者发送的消息会经过 Kafka 的网络层，最终到达目标 broker
4. 确认机制`acks`  
   - `acks=0`：生产者不等待确认，发送后就认为消息发送成功。
   - `acks=1`：生产者等待 leader 节点的确认，消息在 leader 节点写入成功后返回确认。
   - `acks=all`（或 `acks=-1`）：生产者等待所有副本的确认，只有所有副本都成功写入后才返回确认，保证消息的持久性。
5. 错误处理与重试，如果消息发送失败，Kafka 生产者会根据 retries 配置进行重试

#### 生产数据可靠性

主要设置`acks`应答级别可以解决
- `acks=0`：生产者不等待确认，发送后就认为消息发送成功。可靠性差，效率高
- `acks=1`：生产者等待 leader 节点的确认，消息在 leader 节点写入成功后返回确认。可靠性中等，效率中等
- `acks=all`（或 `acks=-1`）：生产者等待所有副本的确认，只有所有副本都成功写入后才返回确认，保证消息的持久性。可靠性高，效率低，数据可能会重复

#### 幂等性
幂等性是指Producer不论向Broker发送多少次重复数据，Broker只会持久化一条
重复数据的判断标准：具有<PID,Partition,SeqNumber>相同主键的消息提交时，Broker只会持久化一条。其中PID是Kafka每次重启都会分配一个新的；Partition表示分区号；SeqNumber是单调递增的
开启参数 `enable.idempotence` 默认为 true，false关闭


### Broker

#### 工作流程

1. Kafka 启动成功后，会在zk中注册
2. 注册的先后顺序可以争夺leader,broker 获取Controller权限
3. 选举出Controller之后，监听Brokers节点变化
4. Controller决定Leader选举，选举规则：在isr中存活为前提。按照AR中排在前面的优先
5. Controller将节点信息上传到ZK上
6. 其他Controller同步Leader Controller 信息

如果Brokers中的Leader挂了，根据第4点的选举规则，选举新的Leader，更新Leader，以及isr

### Kafka副本

Kafka副本主要作用，提高数据可靠性  
Kafka默认1个副本，生产环境一般配置2个  
Kafka分区中的所有副本统称为AR(Assigned Repllicas)  
AR = isr + osr

isr: 表示和Leader保持同步的Follower集合。如果Follower长时间未向Leader发送通信请求或同步数据，则该Follower将被踢出isr

osr: 表示Follower与Leader副本同步时，延迟过多的副本

#### Kafka 文件存储机制
Topic是逻辑上的概念，而partition是物理上的概念，每一个partition对应一个log文件，该log文件中存储的就是Producer生产的数据。Producer生产的数据会被不断追加到该log文件末端，为防止log文件过大导致数据定位效率低下，Kafka采取了分片和索引机制，将每个partition分为多个segment。每一个segment包含：".index"文件、".log"文件、".timeindex"文件。

Kafka存储日志的索引使用的是稀疏索引，大约每往log文件写入4kb数据，会往index文件写入一条索引

Kafka集群日志，默认7天后被清除，清楚策略有两种：删除、压缩策略

- delete 删除日志，`log.cleanup.policy=delete`，基于时间或者基于大小
- compact 日志压缩，对于相同key的不同的value值，保留最后一个版本

#### Kafka高效读写数据

1. Kafka本身是分布式集群，可以采用分区技术，并行度高
2. 读数据采用稀疏索引，可以快速定位要消费的数据
3. 顺序写磁盘，Kafka的producer生产数据，要写入到log文件中，写的过程是一只追加到文件末端，为顺序写。



### Consumer

#### 消费方式
- pull（拉）模式，消费只主动拉去消息，Kafka采用这一个
- push（推）模式，集群推送消息到消费者，网络消耗比较大

#### 常用操作

可以使用`kafka-console-consumer.sh`查看生产者的命令，常用命令如下

| 参数                                                  | 描述                         |
|-----------------------------------------------------|----------------------------|
| `--bootstrap-server <String: server to connect to>` | 连接的 Kafka Broker 主机名称和端口号。 |
| `--topic <String: topic>`                           | 操作的主题名称。                   |
| `--from-beginning`                                  | 从头开始消费                     |
| `--group <String: consumer group id>`               | 指定消费者组名称                   |


```shell
# 接收消息
## 注意如果没有--from-beginning 参数，不会获取到订阅之前的消息
## 先订阅再有消息
kafka-console-consumer.sh --bootstrap-server kafka2:9094 --topic first --from-beginning
```

#### Consumer Group 消费者组
由多个consumer组成，形成一个消费者组的条件，消费者组的groupid相同  
- 消费者组内每个消费者负责消费不同分区的数据，一个分区只能由一个组内消费者消费
- 消费者组之间互不影响。所有消费者都属于某个消费者组，即**消费者组是逻辑上的一个订阅者**

##### 消费者组初始化流程
使用`coordinator`辅助实现消费者组的初始化和分区，每一个broker都有`coordinator`。选择一个`coordinator`作为集群的leader

1. 每一个消费者都会发送一个JoinGroup请求，发送至`coordinator` leader
2. 随机选一个consumer作为leader
3. 把要消费的topic情况发送给消费者leader
4. 消费者leader制定消费方案
5. 把消费方案发给`coordinator`
6. `coordinator`把消费方案发送给各个consumer
7. 每一个消费者都会和`coordinator`保持心跳(3ms)，一旦超时(45s)，该消费者会被移除，并触发再平衡（[重新分配任务](#消费者分区分配)），或者消费者处理消息的时间过长（超过5分钟）

##### 消费者组详细消费流程

1. 消费者发送消费请求到`ConsumerNetworkClient`
2. `ConsumerNetworkClient`会每批次抓数据，当数据达到`fetch.min.bytes`大小或者`fetch.max.wat.ms`时间
3. 抓去数据成功回调，进过序列化到达消费者

##### 消费者分区分配