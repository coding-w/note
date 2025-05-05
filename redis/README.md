# Redis

1. [系统环境](#系统环境)
2. [Docker安装](#docker-安装)
3. [数据类型](#数据类型)
   - [常用数据类型](#常用数据类型)
   - [常用数据结构](#常用数据结构)
4. [持久化](#持久化)
   - [RDB](#rdb-持久化)
   - [AOF](#aof-持久化)
   - [混合模式](#rdb--aof-混合持久化)
5. [Redis-过期策略](#redis-过期策略)
6. [Redis-内存淘汰](#redis-内存淘汰)
7. [主从同步](#主从同步)
8. [Redis-哨兵](#redis-哨兵)
9. [Redis-集群](#redis-集群)
10. [Redis-性能测试](#redis-性能测试)
11. [Redis-常见问题](#redis-常见问题)

## 系统环境
- 系统：macOS
- 处理器：2.6 GHz 六核Intel Core i7
- 内存：16GB
- 版本：15.0

## docker 安装
redis.config
```
# 60s 内有3个值变动，通过bgsave持久化
save 60 3

# bgsave 失败之后，是否停止持久化数据到磁盘，yes 表示停止持久化，no 表示忽略错误继续写文件。
stop-writes-on-bgsave-error yes



# RDB 文件压缩
rdbcompression yes


# rdb & aof 开启混合持久化
aof-use-rdb-preamble yes


requirepass 123456
appendonly yes
protected-mode no
port 6379
```

```shell
docker pull redis:@latest
docker run -d --name redis-1 \
  -e TZ='Asia/Shanghai' \
  -v /Users/wx/workspace/docker/redis/redis.conf:/usr/local/etc/redis/redis.conf \
  -v /Users/wx/workspace/docker/redis/data:/data \
  -p 6380:6379 \
  redis redis-server /usr/local/etc/redis/redis.conf
```

## 数据类型

常用数据类型  
String（字符串），Hash（哈希），List（列表），Set（集合）、Zset（有序集合）

新增数据类型  
HyperLogLog（2.8 版新增）、GEO（3.2 版新增）、Stream（5.0 版新增）

### 常用数据类型

#### String  
String 是最基本的 key-value 结构，key 是唯一标识，value 是具体的值，value其实不仅是字符串，也可以是数字（整数或浮点数），value 最多可以容纳的数据长度是 512MB  
基本用法
```shell
# 设置一个键值对，以及过期时间
set key value [NX|XX] [GET] [EX seconds|PX millise
mset key1 value1 key2 value2 key3 value3
# 获取 key 对用 value
get key
```  
应用场景  

  - 缓存对象

   ```shell
   SET user:1 '{"name":"xiaolin", "age":18}'
   ```
  - 常规计数
   ```shell
   SET aritcle:readcount:1001 0
   incr aritcle:readcount:1001
   ```
  - 分布式锁
   ```shell
   SET lock_key unique_value NX PX 10000
   ```

#### List  
List(双向列表)列表是简单的字符串列表，按照插入顺序排序，可以从头部或尾部向 List 列表添加元素。  
基本用法

```shell
# 向列表的头/尾部插入 n 个元素
LPUSH/RPUSH key value...
# 获取 列表 头/尾元素
LPOP/RPOP key
# 返回列表key中指定区间内的元素，区间以偏移量start和stop指定，从0开始
LRANGE key start stop
# 从key列表表头弹出一个元素，没有就阻塞timeout秒，如果timeout=0则一直阻塞
BLPOP/BRPOP key timeout
```
应用场景
  - 消息队列  
  生产者使用`LPUSH/RPUSH`往列表中插入消息，消费者使用`BRPOP/BLPOP`消费列表中的数据，保证了有序性，也确保了消费者阻塞式读取数据；生产者使用全局唯一ID确保不会重复消息处理；使用`BRPOPLPUSH`确保消息可靠

#### Hash  
Hash是一个键值对(key-value)，`value=[{field1, value1},{field2, value2},{fieldN, valueN}]`，比较适合存储对象  
基本用法
```shell
# 存储一个哈希表的键值
hset stu1 name wx age 12
# 返回哈希表key中field的数量
hlen stu1
# 返回所有的键值
hgetall stu1
```
应用场景
  - 缓存对象，对于对象中某些频繁变化的属性可以用Hash类型来存储
  ```shell
  hset stu1 name wx age 12
  ```

#### Set  
Set类型是一个无序并唯一的键值集合，它的存储顺序不会按照插入的先后顺序进行存储。  
基本用法
```shell
# 往集合key中存入元素，元素存在则忽略，若key不存在则新建
sadd key member [member ...]
# 从集合key中删除元素
srem key member [member ...]
# 获取所有元素
smembers key
# 获取集合中key中的数量
scard key

# 判断 member元素是否存在集合key中
sismember key member

# 从集合key中随机选出count个元素，元素不从key中删除
srandmember key [count]

# 从集合key中随机选出count个元素，元素从集合中删除
spop key [count]

## 集合 交集、并集、差集运算 ....
# 交集
SINTER key [key ...]
# 交集的结果保存在新的集合中
SINTERSTORE destination key [key ...]
# 并集
SUNION key [key ...]
SUNIONSTORE destination key [key ...]
# 差集
SDIFF key [key ...]
SDIFFSTORE destination key [key ...]
```

应用场景

   - 文章点赞，key是文章id，value是用户id
   
   ```shell
   ## 点赞
   sadd article:1 uid:1
   sadd article:1 uid:2
   sadd article:1 uid:3
   ## 取消点赞
   srem article:1 uid:1
   ## 查看所有点赞用户
   smembers article:1
   ```

   - 抽奖活动，集合中存储用户名

   ```shell
   ## 集合存储用户名
   sadd lockly user1 user2 user3 user4 user5 user6
   ## 抽取一个一等奖
   SRANDMEMBER lockly 1
   # 抽取2个二等奖
   SRANDMEMBER lockly 2
   # 抽取三个三等奖
   SRANDMEMBER lockly 3
   ```

#### Zset  
Zset是有序集合，相比Set类型多了一个排序属性score（分数），对于有序集合ZSet来说，每个存储元素相当于两个值组成的，一个是有序集合的元素值，一个是排序值  
基本用法
```shell
# 往有序集合key中加⼊带分值元素
ZADD key score member [[score member]...]
# 往有序集合key中删除元素
ZREM key member [member...]
# 返回有序集合key中元素member的分值
ZSCORE key member
# 返回有序集合key中元素个数
ZCARD key
# 为有序集合key中元素member的分值加上increment
ZINCRBY key increment member
# 正序获取有序集合key从start下标到stop下标的元素
ZRANGE key start stop [WITHSCORES]
# 倒序获取有序集合key从start下标到stop下标的元素
ZREVRANGE key start stop [WITHSCORES]
# 返回有序集合中指定分数区间内的成员，分数由低到⾼排序。
ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
# 返回指定成员区间内的成员，按字典正序排列, 分数必须相同。
ZRANGEBYLEX key min max [LIMIT offset count]
# 返回指定成员区间内的成员，按字典倒序排列, 分数必须相同
ZREVRANGEBYLEX key max min [LIMIT offset count]
```
应用场景
   
   - 排行榜，文章浏览量排行榜
   
   ```shell
   # 初始化数据
   ZADD article:book 200 article:1 50 article:2 150 article:3 220 article:4 100 article:5
   # article:3 增加浏览量
   ZINCRBY article:book 1 article:3
   # 获取浏览量前三的文章和浏览量
   ZREVRANGE article:book 0 2 withscores
   ```

#### BitMap  
BitMap，位图是⼀串连续的⼆进制数组（0和1），可以通过偏移量（offset）定位元素。BitMap通过最⼩的单位bit
来进⾏ 0|1 的设置，表示某个元素的值或者状态，时间复杂度为O(1)。  
基本用法
```shell
# 设置值，其中value只能是 0 和 1
SETBIT key offset value
# 获取值
GETBIT key offset
# 获取指定范围内值为 1 的个数
# start 和 end 以字节为单位
BITCOUNT key start end

# BitMap间的运算
# operations 位移操作符，枚举值
AND 与运算 &
OR 或运算 |
XOR 异或 ^
NOT 取反 ~
# result 计算的结果，会存储在该key中
# key1 … keyn 参与运算的key，可以有多个，空格分割，not运算只能⼀个key
# 当 BITOP 处理不同⻓度的字符串时，较短的那个字符串所缺少的部分会被看作 0。返回值是保存到 destkey
BITOP [operations] [result] [key1] [keyn…]
# 返回指定key中第⼀次出现指定value(0/1)的位置
BITPOS [key] [value]
```
应用场景

   - 签到统计，在签到打卡的场景中，我们只用记录签到(1)或未签到(0)

   ```shell
   # 用户 a 11月9日 签到
   SETBIT uid:a:202411 8 1
   # 检查该用户是否在 11月9日签到
   GETBIT uid:a:202411 8
   # 统计签到次数
   BITCOUNT uid:a:202411
   ```

#### Geo  
Geo 主要用于存储地理位置信息，并对存储的信息进行操作  
基本用法
```shell
# 存储指定的地理空间位置，可以将⼀个或多个经度(longitude)、纬度(latitude)、位置名称(member)添加
GEOADD key longitude latitude member [longitude latitude member ...]
# 从给定的 key ⾥返回所有指定名称(member)的位置（经度和纬度），不存在的返回 nil。
GEOPOS key member [member ...]
# 返回两个给定位置之间的距离。
GEODIST key member1 member2 [m|km|ft|mi]
# 根据⽤户给定的经纬度坐标来获取指定范围内的地理位置集合。
GEORADIUS key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH
```

应用场景

   - 滴滴叫车

   ```shell
   # 设置 司机(33)车辆位置
   GEOADD cars:locations 116.034567 39.030452 33
   # 用户的经纬度附近 5 公里你的车辆信息
   GEORADIUS cars:locations 116.054579 39.030453 5 km ASC COUNT 10
   ```

#### Stream  
Stream 实现了消息队列，支持消息的持久化、支持自动生成全局唯一ID、支持ack确认消息模式、支持消费组模式等  
基本用法

```
XADD - 添加消息到末尾
XTRIM - 对流进行修剪，限制长度
XDEL - 删除消息
XLEN - 获取流包含的元素数量，即消息长度
XRANGE - 获取消息列表，会自动过滤已经删除的消息
XREVRANGE - 反向获取消息列表，ID 从大到小
XREAD - 以阻塞或非阻塞方式获取消息列表
```

应用场景

   - 消息队列

   ```
   // todo
   ```

### 常用数据结构  
Redis有简单动态字符串、双向链表、压缩列表、哈希表、跳表和整数数组，他们和常用数据类型的对应关系如下
![images-2024-12-02-10-47-29](../images/images-2024-12-02-10-47-29.png)

#### 键和值用什么结构组织？  
Redis使用了哈希表来保存所有键值对，哈希表就是一个数组，数组的每个元素称为一个哈希桶，在每一个哈希桶中保存了键值对数据
![images-2024-12-06-15-47-29](../images/images-2024-12-06-15-47-29.png)
   
哈希表保存了所有的键值对，也称为全局哈希表。哈希表的最大好处很明显，就是可以用 O(1) 的时间复杂度来快速查找到键值对——我们只需要计算键的哈希值，就可以知道它所对应的哈希桶位置，然后就可以访问相应的 entry 元素  
哈希表的 O(1) 复杂度和快速查找键值对，也带来一个问题，**哈希冲突会带来操作阻塞**

#### 哈希表操作变慢  
哈希冲突是不可避免的问题，Redis解决哈希冲突的方式是链式哈希，同一个哈希桶中的多个元素用一个链表来保存，它们之间依次用指针连接。
![images-2024-12-06-16-47-29](../images/images-2024-12-06-16-47-29.png)

这里会引发另一个问题，哈希冲突链上的元素只能通过指针逐一查找再操作。如果哈希表里写入的数据越来越多，哈希冲突可能也会越来越多，这就会导致某些哈希冲突链过长，进而导致这个链上的元素查找耗时长，效率降低  
Redis的解决措施是渐进式rehash，Redis 默认使用了两个全局哈希表：哈希表 1 和哈希表 2，默认使用哈希表 1，当开始进行rehash时，分为三步：
- 给哈希表 2 分配更大的空间，例如是当前哈希表 1 大小的两倍；
- 把哈希表 1 中的数据重新映射并拷贝到哈希表 2 中；
- 释放哈希表 1 的空间。

为了避免第二步拷贝全部数据造成阻塞，每处理一个请求时，从哈希表 1 中的第一个索引位置开始，顺带着将这个索引位置上的所有 entries 拷贝到哈希表 2 中。
避免了一次性大量拷贝的开销，分摊到了多次处理请求的过程中，避免了耗时操作，保证了数据的快速访问

#### 压缩列表  
压缩列表实际上类似于一个数组，数组中的每一个元素都对应保存一个数据。和数组不同的是，压缩列表在表头有三个字段 zlbytes、zltail 和 zllen，分别表示列表长度、列表尾的偏移量和列表中的 entry 个数；压缩列表在表尾还有一个 zlend，表示列表结束。

![images-2024-12-06-17-40-29](../images/images-2024-12-06-17-40-29.png)

如果我们要查找定位第一个元素和最后一个元素，可以通过表头三个字段的长度直接定位，复杂度是 O(1)。而查找其他元素时，就没有这么高效了，只能逐个查找，此时的复杂度就是 O(N)

#### 跳表  
有序链表只能逐一查找元素，导致操作起来非常缓慢，于是就出现了跳表。具体来说，跳表在链表的基础上，增加了多级索引，通过索引位置的几个跳转，实现数据的快速定位  
![images-2024-12-06-17-41-29](../images/images-2024-12-06-17-41-29.png)

不同数据结构查找的时间复杂度  
![images-2024-12-06-17-43-29](../images/images-2024-12-06-17-43-29.png)

## 持久化
三种持久化方式
- **快照方式**（RDB, Redis DataBase）将某一个时刻的内存数据，以二进制的方式写入磁盘；
- **文件追加方式**（AOF, Append Only File），记录所有的操作命令，并以文本的形式追加到文件中；
- **混合持久化方式**，Redis 4.0 之后新增的方式，混合持久化是结合了 RDB 和 AOF 的优点，在写入的时候，先把当前的数据以 RDB 的形式写入文件的开头，再将后续的操作命令以 AOF 的格式存入文件，这样既能保证 Redis 重启时的速度，又能减低数据丢失的风险。

### RDB 持久化
RDB（Redis DataBase）是将某一个时刻的内存快照（Snapshot），以二进制的方式写入磁盘的过程。  
RDB 的持久化触发方式有两类：一类是手动触发，另一类是自动触发。
1. 手动触发  
   手动触发持久化的操作有两个： `save` 和 `bgsave` ，它们主要区别体现在：是否阻塞 Redis 主线程的执行
    - `save`：在客户端执行`save`命令，就会触发持久化，会是Redis进入阻塞状态，慎用
    - `bgsave`：bgsave(background save)，通过fork()一个子进程来执行持久化，只有当创建子进程时会有**短暂阻塞**。
2. 自动触发  
   `save m n` 是指在 m 秒内，如果有 n 个键发生改变，满足设置的触发条件，自动执行一次 `bgsave` 命令  
   `flushall` 命令用于清空Redis数据库，并把RDB文件清空
3. 设置配置，在`redis-cli`操作
   ```shell
   master-slave-config get dbfilename
   master-slave-config get dir
   master-slave-config set dir "" # 设置持久化路径
   master-slave-config set save "" # 禁止持久化
   ```
4. 优点  
   - RDB 的内容为二进制的数据，占用内存小，更紧凑，适合作为备份文件
   - RDB 备份使用的是子进程进行数据持久化至磁盘，不会影响主进程  
   
5. 缺点
   - 只能保存某个时间间隔的数据，中途Redis服务意外终止，数据会丢失
   - fork()子进程会将数据持久化至磁盘，数据集很大，持久化CPU性能不佳，会影响到主进程

### AOF 持久化
AOF（Append Only File）中文是附加到文件，顾名思义 AOF 可以把 Redis 每个键值对操作都记录到文件（appendonly.aof）中。  
1. 持久化配置
   ```shell
   master-slave-config get appendonly #查询 AOF是否启动
   master-slave-config set appendonly yes # 启动 AOF
   master-slave-config set appendonly no # 关闭 AOF
   ```
   `redis.conf`中的配置文件中设置`appendonly yes`即可开启AOF

2. 触发持久化  
   触发条件有两种自动触发和手动触发  
   - 自动触发，满足AOF设置的策略触发，`config get appendonly`，获取AOF策略，主要策略  
     - always：每条Redis操作命令都会写入磁盘，最多丢失一条数据
     - everysec：每秒钟写入一次磁盘，最多丢失一秒的数据
     - no：不设置写入磁盘的规则，根据当前操作系统来决定何时写入磁盘，Linux默认30s写入一次数据至磁盘
   - 手动触发，在客户端执行`bgrewriteaof`命令，可以触发文件重写

3. 文件重写  
   AOF 是通过记录 `Redis` 的**执行命令**来持久化（保存）数据的，不断对`Redis`执行命令`AOF`文件会越来越多，这样不仅增加了服务器的存储压力，也会造成 `Redis` 重启速度变慢，为了解决这个问题 Redis 提供了 AOF 重写的功能。主线程 fork 出后台的 `bgrewriteaof` 子进程,bgrewriteaof 子进程就可以在不影响主线程的情况下，逐一把拷贝的数据写成操作，记入重写日志  
   - `auto-aof-rewrite-min-size`允许 AOF 重写的最小文件容量，默认是 64mb
   - `auto-aof-rewrite-percentage`AOF 文件重写的大小比例，默认值是 100，表示 100%

4. 优点
   - AOF 持久化保存的数据更加完整，最多只会丢失 1s 钟的数据
   - AOF 采用的是命令追加的写入方式，所以不会出现文件损坏的问题
   - AOF 持久化文件，非常容易理解和解析，它是把所有 Redis 键值操作命令，即使使用`flushall`也可以恢复数据  
   
5. 缺点
   - 对于相同的数据集来说，AOF 文件要大于 RDB 文件
   - 在 Redis 负载比较高的情况下，RDB 比 AOF 性能更好

### RDB & AOF 混合持久化

`config set aof-use-rdb-preamble yes`开启混合持久化  
当开启了混合持久化时,在 AOF 重写日志时,`fork`出来的重写子进程会先将与主线程共享的内存数据以 RDB 方式写入到 AOF 文件，然后主线程处理的操作命令会被记录在重写缓冲区里，重写缓冲区里的增量命令会以 AOF 方式写入到 AOF 文件，写入完成后通知主进程将新的含有 RDB 格式和 AOF 格式的 AOF 文件替换旧的的 AOF 文件。  
优点
   - 开头为 RDB 的格式，使得 Redis 可以更快的启动，同时结合 AOF 的优点，有减低了大量数据丢失的风险  

缺点
   - AOF 文件中添加了 RDB 格式的内容，使得 AOF 文件的可读性变得很差
   - 不能兼容redis4.0之前的版本

## Redis 过期策略

### 设置过期时间

1. 使用`pexpire`

   - `expire key seconds`：设置 key 在 n 秒后过期
   - `pexpire key milliseconds`：设置 key 在 n 毫秒后过期
   - `expireat key timestamp`：设置 key 在某个时间戳（精确到秒）之后过期
   - `pexpireat key millisecondsTimestamp`：设置 key 在某个时间戳（精确到毫秒）之后过期

2. 字符串中的过期操作

   - `set key value ex seconds`：设置键值对的同时指定过期时间（精确到秒）
   - `set key value px milliseconds`：设置键值对的同时指定过期时间（精确到毫秒）
   - `setex key seconds valule`：设置键值对的同时指定过期时间（精确到秒）

3. 移除过期时间

   `persist key` 可以移除键值的过期时间

### 持久化中的过期键
RDB，从内存状态持久化成 RDB（文件）的时候，会对 key 进行过期检查，过期的键不会被保存到新的 RDB 文件中，因此 Redis 中的过期键不会对生成新 RDB 文件产生任何影响  
AOF，如果数据库某个过期键还没被删除，那么 AOF 文件会保留此过期键，当此过期键被删除后，Redis 会向 AOF 文件追加一条 DEL 命令来显式地删除该键值。重写之后不会被保存

### 过期策略

Redis 会删除已过期的键值，以此来减少 Redis 的空间占用，但因为 Redis 本身是单线的，如果因为删除操作而影响主业务的执行就得不偿失了，为此 Redis 需要制定多个（过期）删除策略来保证糟糕的事情不会发生  
常见的过期策略：定时删除、惰性删除、定期删除

- 定时删除  
   在设置键值过期时间时，创建一个定时事件，当过期时间到达时，由事件处理器自动执行键的删除操作

   - 优点：保证内存可以被尽快释放
   - 缺点：在Redis高负载的时，会造成Redis服务器卡顿

- 惰性删除  
  不主动删除过期键，每次从数据库获取键值时判断是否过期，如果过期则删除键值，并返回 null

  - 优点：每次访问时，才会判断过期键，只使用了很少的系统资源
  - 缺点：系统占用空间删除不及时，导致空间利用率降低，造成了一定的空间浪费

- 定期删除  
   每隔一段时间检查一次数据库，随机删除一些过期键，配置`hz 10`

  - 优点：通过限制删除操作的时长和频率，来减少操作对Redis主业务的影响，同时也能删除一部份过期的数据
  - 缺点：内存清理方面没有定时删除效果好，同时没有惰性删除使用的系统资源少

总结：Redis 使用的是惰性删除加定期删除的过期策略


## Redis 内存淘汰

当 Redis 运行内存已经超过 Redis 设置的最大内存之后，将采用什么策略来删除符合条件的键值对，以此来保障 Redis 高效的运行，其中64位操作系统的没有对内存大小做限制，`config get maxmemory`查看内存限制。

### 内存淘汰策略
1. 内存淘汰策略分类

   - `noeviction`：不淘汰任何数据，当内存不足时，新增操作会报错，Redis 默认内存淘汰策略；
   - `allkeys-lru`：淘汰整个键值中最久未使用的键值；
   - `allkeys-random`：随机淘汰任意键值;
   - `volatile-lru`：淘汰所有设置了过期时间的键值中最久未使用的键值；
   - `volatile-random`：随机淘汰设置了过期时间的任意键值；
   - `volatile-ttl`：优先淘汰更早过期的键值。
   - `volatile-lfu`：淘汰所有设置了过期时间的键值中，最少使用的键值；
   - `allkeys-lfu`：淘汰整个键值中最少使用的键值。

   其中 allkeys-xxx 表示从所有的键值中淘汰数据，而 volatile-xxx 表示从设置了过期键的键值中淘汰数据。

2. 修改 Redis 内存淘汰策略

   - 通过`config set maxmemory-policy xxx` 命令设置。它的优点是设置之后立即生效，不需要重启 Redis 服务，缺点是重启 Redis 之后，设置就会失效
   - 通过修改 Redis 配置文件修改，设置`maxmemory-policy xxx`，它的优点是重启 Redis 服务后配置不会丢失，缺点是必须重启 Redis 服务，设置才能生效

##### 内存淘汰算法

1. LRU算法，LRU 全称是 Least Recently Used 译为最近最少使用，是一种常用的页面置换算法，选择最近最久未使用的页面予以淘汰

2. LFU算法，LFU 全称是 Least Frequently Used 翻译为最不常用的，最不常用的算法是根据总访问次数来淘汰数据的，它的核心思想是“如果数据过去被访问多次，那么将来被访问的频率也更高”

## 主从同步

主从模式（Master-Slave）是Redis实现读写分离和提高容灾能力的基础。一个主节点（Master）可以有多个从节点（Slave）。

1. 特点

   - 读写分离：主节点负责写入，从节点负责读取，减轻主节点的读压力。
   - 故障转移：如果主节点故障，可以手动或通过其他工具将从节点提升为主节点。
   - 数据同步：从节点通过复制（Replication）获取主节点数据。

2. 配置主从

   - 主节点，无需额外配置
   - 从节点，可以修改`redis.conf`文件，添加设置从节点连接主节点`replicaof <master_ip> <master_port>`和配置主节点密码`masterauth <master_password>`，进行持久化配置。
   也可以在`redis-cli`适用命令进行配置，如下
   
     ```shell
     127.0.0.1:6379> replicaof 192.168.7.130 6379
     OK
     127.0.0.1:6379> master-slave-config set masterauth 123456
     OK
     ```

3. 功能测试  
   首先我们先在主服务器上执行保存数据操作，再去从服务器查询，从服务器只能查询数据

   ![images_2024-11-26_10-21-02](../images/images-2024-11-26-10-21-02.png)

4. 主从数据同步  
   新的从服务器连接时，为了保障多个数据库的一致性，主服务器会执行一次 bgsave 命令生成一个 RDB 文件，然后再以 Socket 的方式发送给从服务器，从
服务器收到 RDB 文件之后再把所有的数据加载到自己的程序中，就完成了一次全量的数据同步。

   ![images_2024-11-26_10-21-02](../images/images-2024-11-26-10-57-40.png)

   主从数据同步，还会出现：主从库间网络断了怎么办？主从同步的细节。主从之间的`buffer`等等，暂且作为遗留问题 `todo`

5. 查询服务器角色  
   使用 `role` 命令，来查询当前服务器的主从角色信息

   ![images-2024-11-26-11-25-08](../images/images-2024-11-26-11-25-08.png)

   关闭主从同步，使用 `replicaof no one` 命令来停止从服务器的复制

## Redis 哨兵

### 哨兵是做什么的

在Redis的主从架构中，由于主从模式时读写分离的，如果主节点挂了，那么没有主节点来服务客户端的写操作，也没有主节点给从节点进行数据同步。
哨兵的作用是实现主从节点故障转移。它会检测主节点是否存活，如果发现主节点挂了，他就会选举一个从节点切换位主节点。

### 哨兵的工作机制

哨兵是一个运行在特殊模式下的Redis进程，所以它是一个节点。哨兵主要负责三件事：监控、选主、通知，哨兵是以集群的方式存在的，降低误差，最好是技术个哨兵(3,5,7)

1. 监控主节点，判断主节点是否故障  
   哨兵会每隔1秒给所有主从节点发送PING命令，当主从节点收到PING命令后，会发送一个响应给哨兵，这样判断出聪是否在正常运行。

   ![images-2024-11-28-10-49-44](../images/images-2024-11-28-10-49-44.png)
   
   主从节点没有在规定的时间(`down-after-milliseconds`)内响应哨兵，哨兵就会将其标记位**主观下线(SDOWN)**，单个哨兵的判断可能会误判，通过多个哨兵节点一起判断，就可以降低误判的情况。
   当一个哨兵判读主节点为**主观下线**，就会向其他哨兵发起判断主节点是否下线的命令，当其他哨兵认同主观下线的数量超过设置的`quorum`时，主节点就会判定为**客观下线**。  
   哨兵判断主节点客观下线后，哨兵要在多个从节点中，选出一个从节点来做新的主节点。

2. 选举哨兵进行主从故障转移  
   主从故障转移是由一个被选举为“领导者”（Leader）的哨兵负责执行的。这种机制通过哨兵间的选举流程确定具体负责执行故障转移的哨兵。当哨兵检测到主节点进入客观下线（ODOWN）状态后，任何哨兵都可以成为“候选人”。候选人会自动投票给自己，同时发起投票请求，争取其他哨兵的支持。每个哨兵只能在每轮选举中投票一次。
   当满足：
   - 候选人获得超过半数哨兵实例（N/2 + 1）的支持票
   
   就可以成为主从故障转移的领导者

### 哨兵主从故障迁移
1. 选择新的主节点，新主节点满足以下条件：

   - 节点必须是在线状态
   - 复制优先级(replica-priority)值越高，优先级越高，主节点的选择概率高
   - 复制偏移量(数据同步状态)，优先级相同的条件下，会选择复制偏移量接近主节点的从节点
   - 从节点ID，上述条件均相同，则根据从节点的Run ID，选择ID最小的节点

   将选定的从节点提升为新的主节点，Sentinel 会向该从节点发送 `SLAVEOF NO ONE` 命令，使其成为主节点

2. 重新配置其他从节点  
   将剩余的从节点重新配置为从属于新的主节点，Sentinel 会向这些从节点发送 `SLAVEOF <new-master-ip> <new-master-port>` 命令

3. 将旧主节点重新配置为从节点  
   如果故障的旧主节点重新上线，Sentinel 会将其重新配置为新主节点的从节点，
   Sentinel 向旧主节点发送 `SLAVEOF <new-master-ip> <new-master-port>` 命令

4. 通知客户端更新主节点地址  
   Sentinel 将通过 Pub/Sub 通知机制通知客户端新的主节点地址`todo`

5. 故障转移流程图  
   ```
   Step 1: 主节点发生故障
            ↓
   Step 2: 哨兵检测到主节点 SDOWN
            ↓
   Step 3: 多数哨兵确认主节点 ODOWN
            ↓
   Step 4: 哨兵间选举出 Leader 哨兵
            ↓
   Step 5: Leader 哨兵选定一个从节点作为新主节点
            ↓
   Step 6: Leader 哨兵完成主从切换，并通知集群
   ```

遗留问题`todo`

从节点升为主节点，新的主节点数据和前主节点数据不完全同步，该怎么办？

在故障迁移时，有数据写入，该怎么办？

### 实操

1. 项目目录

   ```text
   .
   ├── docker-compose.yml          # docker-compose 配置文件
   ├── redis-0                     # redis-0 配置目录
   │   ├── data                    # redis-0 数据存储目录
   │   └── config                  # redis-0 配置文件目录
   │       └── redis.conf          # redis 配置文件
   ├── redis-1                     # redis-1 配置目录
   │   ├── data                    # redis-1 数据存储目录
   │   └── config                  # redis-1 配置文件目录
   │       └── redis.conf          # redis 配置文件
   ├── redis-2                     # redis-2 配置目录
   │   ├── data                    # redis-2 数据存储目录
   │   └── config                  # redis-2 配置文件目录
   │       └── redis.conf          # redis 配置文件
   ├── sentinel1                   # sentinel-1 配置目录
   │   ├── data                    # sentinel-1 数据存储目录
   │   └── config                  # sentinel-1 配置文件目录
   │       └── sentinel.conf       # sentinel 配置文件
   ├── sentinel2                   # sentinel-2 配置目录
   │   ├── data                    # sentinel-2 数据存储目录
   │   └── config                  # sentinel-2 配置文件目录
   │       └── sentinel.conf       # sentinel 配置文件
   ├── sentinel3                   # sentinel-3 配置目录
   │   ├── data                    # sentinel-3 数据存储目录
   │   └── config                  # sentinel-3 配置文件目录
   │       └── sentinel.conf       # sentinel 配置文件
   ```

2. [redis.conf](master-slave-config/redis.conf) redis基础配置，注意事项

   - 主节点需要开启持久化配置，[RDB-AOF混合持久化](#rdb--aof-混合持久化)
   - 若需要设置密码，主从库密码要一致
   - 从库需要配置`replicaof 172.30.1.2 6379`从节点连接主节点的地址和端口，以及主节点的密码`masterauth 123456`

3. [sentinel.conf](master-slave-config/sentinel.conf) 哨兵配置，注意事项

   - 必须设置监控的主节点`sentinel monitor  <ip> <port> <quorum>`
   - 哨兵在这个时间内未收到主节点的响应，则认为主节点不可用`sentinel down-after-milliseconds <master-name> 5000`

4. [docker-compose.yaml](master-slave-config/docker-compose.yaml) 文件，详情请见配置文件，注意事项如下

   - 如果在同一台机器上，使用docker做测试，需要新建一个网络通信
   - 将配置文件的目录挂载至容器，而非文件，哨兵需要对文件修改，需要获得目录和文件的权限

5. 测试
   
   - 在目标目录下，使用`docker-compose up -d`启动主从`Redis`和`sentinel`集群，进入容器，连接Redis
   
   ![images-2024-11-28-14-54-20](../images/images-2024-11-28-14-54-20.png)

   - 使用`docker stop redis-0`停止主库，主库挂了，`sentinel`重新开始设置主库

   ![images-2024-11-28-14-59-02](../images/images-2024-11-28-14-57-03.png)

   - 使用`docker start redis-0`，前主库恢复，主库变从库

   ![images-2024-11-28-14-59-02](../images/images-2024-11-28-14-59-02.png)

   - sentinel 日志分析
   
   ```text
   1:X 28 Nov 2024 14:46:04.988 * oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
   1:X 28 Nov 2024 14:46:04.989 * Redis version=7.4.1, bits=64, commit=00000000, modified=0, pid=1, just started
   1:X 28 Nov 2024 14:46:04.990 * Configuration loaded         # 加载配置文件
   1:X 28 Nov 2024 14:46:04.991 * monotonic clock: POSIX clock_gettime     # Redis 使用单调 POSIX 时钟来测量时间，以保证时间戳的精确性
   1:X 28 Nov 2024 14:46:04.993 * Running mode=sentinel, port=26379.
   1:X 28 Nov 2024 14:46:05.000 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:46:05.000 * Sentinel ID is 64d59cb8360bb0deaffb4e178f3f2d28b0cc9b3a      # Sentinel 的唯一标识符
   1:X 28 Nov 2024 14:46:05.001 # +monitor master mymaster 172.30.1.2 6379 quorum 2    # 主库 ip 端口，quorum 客观下线 数量
   1:X 28 Nov 2024 14:46:05.003 * +slave slave 172.30.1.4:6379 172.30.1.4 6379 @ mymaster 172.30.1.2 6379  # 从节点 172.30.1.4:6379 被检测到并注册为主节点 mymaster 的从节点
   1:X 28 Nov 2024 14:46:05.007 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:46:05.008 * +slave slave 172.30.1.3:6379 172.30.1.3 6379 @ mymaster 172.30.1.2 6379  # 从节点 172.30.1.3:6379 被检测到并注册为主节点 mymaster 的从节点
   1:X 28 Nov 2024 14:46:05.013 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:46:07.015 * +sentinel sentinel 4a4da7d251e2cad526a3f69d21f7c44bfd9027db 172.30.1.6 26379 @ mymaster 172.30.1.2 6379  # 新的哨兵加入，监控主节点
   1:X 28 Nov 2024 14:46:07.021 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:46:07.023 * +sentinel sentinel 171327d57c7fc1e1c2fd8fde9726ea9cd1b21628 172.30.1.7 26379 @ mymaster 172.30.1.2 6379  # 新的哨兵加入，监控主节点
   1:X 28 Nov 2024 14:46:07.030 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:55:26.245 # +sdown master mymaster 172.30.1.2 6379   # 主库，主观下线
   1:X 28 Nov 2024 14:55:26.325 # +odown master mymaster 172.30.1.2 6379 #quorum 2/2   # 确认主库  客观下线
   1:X 28 Nov 2024 14:55:26.326 # +new-epoch 1 # 开启了一个新的纪元（epoch），用于协调故障转移
   1:X 28 Nov 2024 14:55:26.328 # +try-failover master mymaster 172.30.1.2 6379    # Sentinel 开始对主节点进行故障转移
   1:X 28 Nov 2024 14:55:26.333 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:55:26.333 # +vote-for-leader 64d59cb8360bb0deaffb4e178f3f2d28b0cc9b3a 1      # 投票给自己，当领导
   1:X 28 Nov 2024 14:55:26.345 * 4a4da7d251e2cad526a3f69d21f7c44bfd9027db voted for 64d59cb8360bb0deaffb4e178f3f2d28b0cc9b3a 1    # 其他哨兵投票
   1:X 28 Nov 2024 14:55:26.347 * 171327d57c7fc1e1c2fd8fde9726ea9cd1b21628 voted for 64d59cb8360bb0deaffb4e178f3f2d28b0cc9b3a 1
   1:X 28 Nov 2024 14:55:26.412 # +elected-leader master mymaster 172.30.1.2 6379  # 选举完成，当前 Sentinel 成为领导者
   1:X 28 Nov 2024 14:55:26.413 # +failover-state-select-slave master mymaster 172.30.1.2 6379     # 领导者开始选择新的主节点
   1:X 28 Nov 2024 14:55:26.468 # +selected-slave slave 172.30.1.3:6379 172.30.1.3 6379 @ mymaster 172.30.1.2 6379     # 从节点 172.30.1.3:6379 被选为新的主节点
   1:X 28 Nov 2024 14:55:26.470 * +failover-state-send-slaveof-noone slave 172.30.1.3:6379 172.30.1.3 6379 @ mymaster 172.30.1.2 6379      # 向选定的新主服务器发送命令使其不再成为任何其他服务器的从属
   1:X 28 Nov 2024 14:55:26.539 * +failover-state-wait-promotion slave 172.30.1.3:6379 172.30.1.3 6379 @ mymaster 172.30.1.2 6379      # 等待新主服务器的晋升完成
   1:X 28 Nov 2024 14:55:26.604 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:55:26.605 # +promoted-slave slave 172.30.1.3:6379 172.30.1.3 6379 @ mymaster 172.30.1.2 6379     # 从服务器172.30.1.3已被提升为新的主服务器
   1:X 28 Nov 2024 14:55:26.605 # +failover-state-reconf-slaves master mymaster 172.30.1.2 6379        # 开始重新配置剩余的从服务器以跟随新的主服务器
   1:X 28 Nov 2024 14:55:26.700 * +slave-reconf-sent slave 172.30.1.4:6379 172.30.1.4 6379 @ mymaster 172.30.1.2 6379      # 向从服务器172.30.1.4发送重新配置指令
   1:X 28 Nov 2024 14:55:27.455 # -odown master mymaster 172.30.1.2 6379           # 原来的主服务器172.30.1.2现在被认为是可用的（不再odown
   1:X 28 Nov 2024 14:55:27.669 * +slave-reconf-inprog slave 172.30.1.4:6379 172.30.1.4 6379 @ mymaster 172.30.1.2 6379        # 从服务器172.30.1.4的重新配置正在进行中
   1:X 28 Nov 2024 14:55:27.671 * +slave-reconf-done slave 172.30.1.4:6379 172.30.1.4 6379 @ mymaster 172.30.1.2 6379      # 从服务器172.30.1.4的重新配置已完成
   1:X 28 Nov 2024 14:55:27.723 # +failover-end master mymaster 172.30.1.2 6379            # 故障转移结束
   1:X 28 Nov 2024 14:55:27.725 # +switch-master mymaster 172.30.1.2 6379 172.30.1.3 6379      # 主服务器已切换至172.30.1.3
   1:X 28 Nov 2024 14:55:27.726 * +slave slave 172.30.1.4:6379 172.30.1.4 6379 @ mymaster 172.30.1.3 6379  # 从服务器172.30.1.4现在跟随新的主服务器172.30.1.3
   1:X 28 Nov 2024 14:55:27.727 * +slave slave 172.30.1.2:6379 172.30.1.2 6379 @ mymaster 172.30.1.3 6379  # 原来的主服务器172.30.1.2现在也是从服务器，跟随新的主服务器172.30.1.3
   1:X 28 Nov 2024 14:55:27.732 * Sentinel new configuration saved on disk
   1:X 28 Nov 2024 14:55:32.787 # +sdown slave 172.30.1.2:6379 172.30.1.2 6379 @ mymaster 172.30.1.3 6379  # 原来的主服务器172.30.1.2（现在作为从服务器）被认为主观下线
   1:X 28 Nov 2024 14:57:23.806 # -sdown slave 172.30.1.2:6379 172.30.1.2 6379 @ mymaster 172.30.1.3 6379 # 原来的主服务器172.30.1.2（现在作为从服务器）再次被认为是可用的（不再sdown）
   ```

## Redis 集群

[Redis scales horizontally with a deployment topology called Redis Cluster.](https://redis.io/docs/latest/operate/oss_and_stack/management/scaling/)

Redis采用集群的主要目的是为了解决单机性能和容量的限制问题，增加了高可用性

### 集群模式
#### Redis Cluster  
Redis 自带的原生集群模式，支持数据分片  
特点：

- 分片算法：由客户端决定分片规则，通常采用一致性哈希
- 灵活性搞：客户端完全掌控分片逻辑
- 无高可用支持：需要客户端额外管理主从关系和故障管理

使用场景：
- 数据分布策略需要高度定制化
- 不需要复杂的高可用管理，或已有外部的高可用方案

数据分片(data sharding)  
`Redis`集群不使用一致哈希，而是使用一种不同形式的分片，其中每个键在概念上都是我们称之为哈希槽的一部分。
`Redis`集群中有16384个哈希槽，为了计算给定键的哈希槽，我们只需将键进行CRC16哈希然后对16384取模。
`Redis`集群中的每个节点负责哈希槽的一个子集，例如，有一个有3个节点(A,B,C)的集群，其中：

- 节点A包含0 ~ 5500个哈希槽
- 节点B包含从5501到11000的哈希槽
- 节点C包含从11001到16383的哈希槽


集群端口，集群的每一个节点需要打开两个端口，一个用于为客户端提供服务的端口，另一个为集群的总线端口，默认情况下，集群总线端口是10000+客户端端口。可以通过`cluster-port`设置


Redis Cluster 主从模型，为了保证节点发送故障后节点通信是保持可用，Redis Cluster 使用主从的模式，每一个主节点都有多个副节点。
三个主节点三个从节点(A,B,C,A1,B1,C1)，如果主节点B发送故障，集群将提升节点B1为新的主节点继续正常运行，如果B和B1同时发生故障，则`Redis Cluster`将无法继续运行。

Redis集群不能保证强一致性，集群可能会丢失客户端的确认写入

#### Codis

Codis 是由国人开发的开源Redis分布式代理解决方案，旨在帮助用户管理多个Redis实例，**不在更新**  
特点：
- 隐藏了分片逻辑，客户端只需要连接Proxy，不需要感知底层的分布式结构
- 支持横向扩展，多个Proxy可共同分担流量
- 支持高可用，Proxy故障时，客户端可以切换到其他Proxy

使用场景：
- 数据规模大
- 请求并发量高

#### Redis Cluster 与 Codis 的对比

| 特性     | Redis Cluster           | Codis                      |
|--------|-------------------------|----------------------------|
| 分片实现   | 数据分布到16384个哈希槽，有节点负责    | 基于 Proxy，由 Codis 管理分片规则    |
| 客户端透明性 | 需要支持Cluster协议的客户端       | 完全透明，客户端无需感知               |
| 动态扩容   | 数据迁移较复杂，需手动操作           | 支持在线迁移分片                   |
| 高可用性   | 原生支持高可用，自主选举主节点         | 	依赖 Redis 主从复制和 Proxy 故障切换 |
| 使用复杂度  | 相对简单，原生功能               | 相对复杂，需要运维 ZooKeeper/Etcd	  |
| 场景适用性  | 新系统开发，支持Redis Cluster协议 | 适合老系统改造，无需更改业务代码           |

### Cluster 实操

1. 集群目录

   ```text
   .
   ├── 7000
   │   ├── conf
   │   │   └── redis.conf
   │   └── data
   │
   ├── 7001
   │   ├── conf
   │   │   └── redis.conf
   │   └── data
   │
   ├── 7002
   │   ├── conf
   │   │   └── redis.conf
   │   └── data
   │
   ├── 7003
   │   ├── conf
   │   │   └── redis.conf
   │   └── data
   │
   ├── 7004
   │   ├── conf
   │   │   └── redis.conf
   │   └── data
   │
   ├── 7005
   │   ├── conf
   │   │   └── redis.conf
   │   └── data
   │
   └── docker-compose.yaml
   ```

2. [redis.conf](./cluster/redis.conf) redis基础配置，注意事项

   - `cluster-enabled yes` 启用 Redis 的集群模式
   - `cluster-config-file /data/nodes.conf` 指定存储 Redis 集群状态的配置文件路径
   - `cluster-node-timeout 5000` 设置 Redis 集群节点间的超时时间
   - `cluster-require-full-coverage yes` 默认yes，某主从节点都挂掉了，整个集群都无法提供服务

3. [docker-compose.yaml](./cluster/docker-compose.yaml)  
   这次实操，是在一个服务器上启动了 6 个节点，并且在同一主机使用docker部署，直接使用的是主机网络，避免出现网络问题

4. 创建集群

   ```shell
   # 先连接任意一个节点
   docker exec -it redis-0 /bin/bash
   # 创建一个新的集群 --cluster-replicas 1 每一个主节点配置一个从节点
   redis-cli --cluster create 192.168.3.222:7000 192.168.3.222:7001 \
   192.168.3.222:7002 192.168.3.222:7003 192.168.3.222:7004 192.168.3.222:7005 \
   --cluster-replicas 1
   ```

5. 日志分析

   ```text
   >>> Performing hash slots allocation on 6 nodes...    # 给三个主库分配哈希槽
   Master[0] -> Slots 0 - 5460
   Master[1] -> Slots 5461 - 10922
   Master[2] -> Slots 10923 - 16383
   Adding replica 192.168.3.222:7004 to 192.168.3.222:7000         # 给每一个主节点分配一个从节点
   Adding replica 192.168.3.222:7005 to 192.168.3.222:7001
   Adding replica 192.168.3.222:7003 to 192.168.3.222:7002
   >>> Trying to optimize slaves allocation for anti-affinity          # 避免亲和性（anti-affinity），尽量不在同一台物理机上放置主节点及其对应的从节点，下面发出警告
   [WARNING] Some slaves are in the same host as their master
   M: d886b043376653f642bbe5c45149a2f77e83b7aa 192.168.3.222:7000      # 三个主节点的唯一id，以及负责哈希槽的范围
      slots:[0-5460] (5461 slots) master
   M: c2aa2c6a1aea895b17c80e470cc8a3805f257672 192.168.3.222:7001
      slots:[5461-10922] (5462 slots) master
   M: 41c5e7238fdf873506ba32bd21f26c6aeffdd3d1 192.168.3.222:7002
      slots:[10923-16383] (5461 slots) master
   S: 950f242794e4b26afe5e2b82d0198c79ff32ef2f 192.168.3.222:7003      # 三个从节点的信息，以及他们的主节点
      replicates c2aa2c6a1aea895b17c80e470cc8a3805f257672
   S: 01212bd1dc1bed4bccc0fea21ad5dcbfe384fa5b 192.168.3.222:7004
      replicates 41c5e7238fdf873506ba32bd21f26c6aeffdd3d1
   S: 25940bbff999b1dd1f143f8af5fddf9dd65a285e 192.168.3.222:7005
      replicates d886b043376653f642bbe5c45149a2f77e83b7aa
   Can I set the above configuration? (type 'yes' to accept): yes
   >>> Nodes configuration updated
   >>> Assign a different config epoch to each node
   >>> Sending CLUSTER MEET messages to join the cluster
   Waiting for the cluster to join
   .
   >>> Performing Cluster Check (using node 192.168.3.222:7000)        # 等待所有节点加入集群
   M: d886b043376653f642bbe5c45149a2f77e83b7aa 192.168.3.222:7000
      slots:[0-5460] (5461 slots) master
      1 additional replica(s)
   M: 41c5e7238fdf873506ba32bd21f26c6aeffdd3d1 192.168.3.222:7002
      slots:[10923-16383] (5461 slots) master
      1 additional replica(s)
   S: 25940bbff999b1dd1f143f8af5fddf9dd65a285e 192.168.3.222:7005
      slots: (0 slots) slave
      replicates d886b043376653f642bbe5c45149a2f77e83b7aa
   S: 01212bd1dc1bed4bccc0fea21ad5dcbfe384fa5b 192.168.3.222:7004
      slots: (0 slots) slave
      replicates 41c5e7238fdf873506ba32bd21f26c6aeffdd3d1
   M: c2aa2c6a1aea895b17c80e470cc8a3805f257672 192.168.3.222:7001
      slots:[5461-10922] (5462 slots) master
      1 additional replica(s)
   S: 950f242794e4b26afe5e2b82d0198c79ff32ef2f 192.168.3.222:7003
      slots: (0 slots) slave
      replicates c2aa2c6a1aea895b17c80e470cc8a3805f257672
   [OK] All nodes agree about slots configuration.
   >>> Check for open slots...
   >>> Check slots coverage...
   [OK] All 16384 slots covered.
   ```
   
   nodes.conf，中主要记录着集群的主从，以及其分配的哈希槽
   ```text
   # 格式 提供了每个节点的状态、角色、连接情况以及其他相关信息
   <node_id> <ip:port>@<cluster_bus_port>,<optional_fields> <flags> <master_id> <last_ping_sent> <last_pong_rcvd> <config_epoch> <link_state> <slots>
   # 详细如下
   41c5e7238fdf873506ba32bd21f26c6aeffdd3d1 192.168.3.222:7002@17002,,tls-port=0,shard-id=895b2a0394527b0cd73aa5434be44130a1fc3254 master - 0 1733302987000 3 connected 10923-16383
   d886b043376653f642bbe5c45149a2f77e83b7aa 192.168.3.222:7000@17000,,tls-port=0,shard-id=dfb8b8992365af8b487dc1f3078b5c07c70d8379 myself,master - 0 0 1 connected 0-5460
   25940bbff999b1dd1f143f8af5fddf9dd65a285e 192.168.3.222:7005@17005,,tls-port=0,shard-id=dfb8b8992365af8b487dc1f3078b5c07c70d8379 slave d886b043376653f642bbe5c45149a2f77e83b7aa 0 1733302987520 1 connected
   01212bd1dc1bed4bccc0fea21ad5dcbfe384fa5b 192.168.3.222:7004@17004,,tls-port=0,shard-id=895b2a0394527b0cd73aa5434be44130a1fc3254 slave 41c5e7238fdf873506ba32bd21f26c6aeffdd3d1 0 1733302987000 3 connected
   c2aa2c6a1aea895b17c80e470cc8a3805f257672 192.168.3.222:7001@17001,,tls-port=0,shard-id=426b50aaddaa31109ca57eb0b50853bbe58fd9c7 master - 0 1733302987520 2 connected 5461-10922
   950f242794e4b26afe5e2b82d0198c79ff32ef2f 192.168.3.222:7003@17003,,tls-port=0,shard-id=426b50aaddaa31109ca57eb0b50853bbe58fd9c7 slave c2aa2c6a1aea895b17c80e470cc8a3805f257672 0 1733302987722 2 connected
   vars currentEpoch 6 lastVoteEpoch 0
   ```
   
   连接和使用
   ```shell
   root@localhost:/data# redis-cli -c -p 7000
   127.0.0.1:7000> set foo bar
   -> Redirected to slot [12182] located at 192.168.3.222:7002
   OK
   192.168.3.222:7002> set name wang
   -> Redirected to slot [5798] located at 192.168.3.222:7001
   OK
   192.168.3.222:7001> get name
   "wang"
   192.168.3.222:7001> get foo
   -> Redirected to slot [12182] located at 192.168.3.222:7002
   "bar"
   192.168.3.222:7002>
   ```
   
6. 对集群重新分片  
   详细文章请查阅[Redis官方文档](https://redis.io/docs/latest/operate/oss_and_stack/management/scaling/)  
   重新分片可以自动执行，无需以交互方式手动输入参数
   `redis-cli --cluster reshard <host>:<port> --cluster-from <node-id> --cluster-to <node-id> --cluster-slots <number of slots> --cluster-yes`

   - `<host>:<port>` 指定一个集群中的任意节点地址和端口，这个节点将被用来连接到集群并发起重新分片操作
   - `--cluster-from <node-id>` 指定源节点的 ID，从此节点中移出哈希槽的主节点 ID
   - `--cluster-to <node-id>` 指定目标节点的 ID，将哈希槽移入的目标主节点 ID
   - `--cluster-slots <number of slots>` 指定要移动的哈希槽数量，这个值必须在1到16384之间
   - `--cluster-yes` 自动确认所有提示，这使得命令可以在没有用户交互的情况下完成整个重新分片过程

## Redis 性能测试

对Redis进行性能测试，评估Redis缓存能不能顶主系统的吞吐量，对Redis的性能进行测试，有两种测试方式：

   - 编写脚本模拟并发测试
   - 使用`redis-benchmark`进行测试，Redis自带的性能测试工具

目前的水平有限，也不是专攻Redis中间件，就不编写脚本实现测试`todo`，学习和记录一下`redis-benchmark`工具，[官方原文](https://redis.io/docs/latest/operate/oss_and_stack/management/optimization/benchmarks/)

`redis-benchmark -h`查看redis最常用的命令
```text
Usage: redis-benchmark [OPTIONS] [COMMAND ARGS...]

Options:
 -h <hostname>      Server hostname (default 127.0.0.1)
 -p <port>          Server port (default 6379)
 -s <socket>        Server socket (overrides host and port)
 -a <password>      Password for Redis Auth
 --user <username>  Used to send ACL style 'AUTH username pass'. Needs -a.
 -u <uri>           Server URI on format redis://user:password@host:port/dbnum
                    User, password and dbnum are optional. For authentication
                    without a username, use username 'default'. For TLS, use
                    the scheme 'rediss'.
 -c <clients>       Number of parallel connections (default 50).
                    Note: If --cluster is used then number of clients has to be
                    the same or higher than the number of nodes.
 -n <requests>      Total number of requests (default 100000)
 -d <size>          Data size of SET/GET value in bytes (default 3)
 --dbnum <db>       SELECT the specified db number (default 0)
 -3                 Start session in RESP3 protocol mode.
 --threads <num>    Enable multi-thread mode.
 --cluster          Enable cluster mode.
                    If the command is supplied on the command line in cluster
                    mode, the key must contain "{tag}". Otherwise, the
                    command will not be sent to the right cluster node.
 --enable-tracking  Send CLIENT TRACKING on before starting benchmark.
 -k <boolean>       1=keep alive 0=reconnect (default 1)
 -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD,
                    random members and scores for ZADD.
                    Using this option the benchmark will expand the string
                    __rand_int__ inside an argument with a 12 digits number in
                    the specified range from 0 to keyspacelen-1. The
                    substitution changes every time a command is executed.
                    Default tests use this to hit random keys in the specified
                    range.
                    Note: If -r is omitted, all commands in a benchmark will
                    use the same key.
 -P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
 -q                 Quiet. Just show query/sec values
 --precision        Number of decimal places to display in latency output (default 0)
 --csv              Output in CSV format
 -l                 Loop. Run the tests forever
 -t <tests>         Only run the comma separated list of tests. The test
                    names are the same as the ones produced as output.
                    The -t option is ignored if a specific command is supplied
                    on the command line.
 -I                 Idle mode. Just open N idle connections and wait.
 -x                 Read last argument from STDIN.
```

- `-c <clients>`指定并发的连接数量，默认50
- `-n <requests>`发出的请求总数，默认100000
- `-d <size> `对于 SET/GET 操作的数据大小（单位：字节），默认是 3 字节
- `--dbnum <db>`选择使用的数据库编号，默认是 0
- `--threads <num>`启用多线程模式，并指定使用的线程数
- `--cluster`启用集群模式
- `-k <boolean>`是否保持长连接（keep alive），默认 1 表示保持，0 表示每次请求后断开并重新连接
- `-r <keyspacelen>`用于生成随机键的空间长度。如 -r 1000，那么键将从 0 到 999 中随机选取
- `-P <numreq>`通过管道一次发送的请求数量，默认是 1
- `-q`静默测试，只显示 QPS 的值
- `--precision`设定输出延迟结果的小数点精度，默认是 0，意味着不保留小数部分
- `--csv`将测试结果输出为 CSV 格式的文件
- `-l`循环测试
- `-t <tests>`仅运行逗号分隔的测试列表
- `-I`空闲模式


```shell
redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 100000 -d 256 -t set,get -r 100 -a 123456
====== SET ======                                                   
100000 requests completed in 1.59 seconds
50 parallel clients
256 bytes payload
keep alive: 1
host configuration "save": 60 3
host configuration "appendonly": yes
multi-thread: no
......
Summary:
throughput summary: 62932.66 requests per second
latency summary (msec):
avg       min       p50       p95       p99       max
0.709     0.160     0.695     0.943     1.135     8.031

====== GET ======                                                     
100000 requests completed in 0.89 seconds
50 parallel clients
256 bytes payload
keep alive: 1
host configuration "save": 60 3
host configuration "appendonly": yes
multi-thread: no
Summary:
throughput summary: 111731.84 requests per second
latency summary (msec):
avg       min       p50       p95       p99       max
0.236     0.096     0.231     0.327     0.407     0.863
```

|                | SET            | GET             |
|----------------|----------------|-----------------|
| 吞吐量            | 62,932.66 请求/秒 | 111,731.84 请求/秒 |
| 平均延迟 (avg)     | 0.709 ms       | 0.236 ms        |
| 最小延迟 (min)     | 0.160 ms       | 0.096 ms        |
| 中位数延迟 (p50)    | 0.695 ms       | 0.231 ms        |
| 95 百分位延迟 (p95) | 0.943 ms       | 0.327 ms        |
| 99 百分位延迟 (p99) | 1.135 ms       | 0.407 ms        |
| 最大延迟 (max)     | 8.031 ms       | 0.863 ms        |

## Redis 常见问题

### 主从同步原理

1. 主从全量同步，主从第一次同步  
   - 主节点通过`bgsave`生成RDB文件，将文件传递到从节点
   - 此时主节点新操作的数据会存入

### 缓存中的数据和数据库不一致

数据的一致性
   a. 缓存中有数据，那么，缓存的数据值需要和数据库中的值相同
   b. 缓存中本身没有数据，那么，数据库中的值必须是最新值

解决措施：  
1. 先删除缓存，再更新数据库  
   可能会出现的问题，假设线程 A 删除缓存值后，还没有来得及更新数据库（比如说有网络延迟），线程 B 就开始读取数据了，那么这个时候，线程 B 会发现缓存缺失，就只能去数据库读取。这会带来两个问题：
   a. 线程 B 读取到了旧值；
   b. 线程 B 是在缓存缺失的情况下读取的数据库，所以，它还会把旧值写入缓存，这可能会导致其他线程从缓存中读到旧值。

   ![images-2024-12-09-14-51-46](../images/images-2024-12-09-14-51-46.png)

   在线程 A 更新完数据库值以后，我们可以让它先 sleep 一小段时间，再进行一次缓存删除操作。

2. 先更新数据库值，再删除缓存值  

   可能会出现的问题，如果线程 A 删除了数据库中的值，但还没来得及删除缓存值，线程 B 就开始读取数据了，那么此时，线程 B 查询缓存时，发现缓存命中，就会直接从缓存中读取旧值。
   不过，在这种情况下，如果其他线程并发读缓存的请求不多，那么，就不会有很多请求读取到旧值。而且，线程 A 一般也会很快删除缓存值，这样一来，其他线程再次读取时，就会发生缓存缺失，进而从数据库中读取最新值。所以，这种情况对业务的影响较小。

   ![](../images/images-2024-12-09-14-52-01.png)

在大多数业务场景下，把 Redis 作为只读缓存使用。针对只读缓存来说，我可以**先删除缓存值再更新数据库**，也可以**先更新数据库再删除缓存**。建议是，优先使用**先更新数据库再删除缓存**的方法，原因主要有两个：

   - 先删除缓存值再更新数据库，有可能导致请求因缓存缺失而访问数据库，给数据库带来压力；
   - 如果业务应用中读取数据库和写缓存的时间不好估算，那么，延迟双删中的等待时间就不好设置。

不过，当使用先更新数据库再删除缓存时，也有个地方需要注意，如果业务层要求必须读取一致的数据，那么，我们就需要在更新数据库时，先在 Redis 缓存客户端暂存并发读请求，等数据库更新完、缓存值删除后，再读取数据，从而保证数据一致性。

### 缓存雪崩

大量数据不在缓存中，大量请求发送到数据库层，导致数据库层的压力激增，出现原因：

1. 大量数据同事过期，解决措施

   - 避免给大量数据设置相同的过期时间
   - 服务降级，指发生缓存雪崩时，针对不同的数据采取不同的处理

2. Redis突然宕机

   - 在业务系统中实现服务熔断或请求限流机制
   - 事前预防，构建可靠集群

### 缓存击穿

热点数据请求，没有在缓存中进行处理  
解决措施：超长过期时间，或者不设置过期时间

### 缓存穿透

访问的数据既不在 Redis 缓存中，也不在数据库中，导致请求在访问缓存时，发生缓存缺失，再去访问数据库时，发现数据库中也没有要访问的数据
出现的两种情况：

   - 业务层误操作
   - 恶意攻击

解决措施

   - 缓存空值
   - 使用布隆过滤器快速判断数据是否存在，避免从数据库中查询数据是否存在，减轻数据库压力
   - 前端进行请求检测，黑名单

![images-2024-12-09-15-08-03](../images/images-2024-12-09-15-08-03.png)


