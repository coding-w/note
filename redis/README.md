### Redis

#### 系统环境
- 系统：macOS
- 处理器：2.6 GHz 六核Intel Core i7
- 内存：16GB
- 版本：15.0

#### docker 安装
redis.config
```
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

#### 持久化

三种持久化方式

- **快照方式**（RDB, Redis DataBase）将某一个时刻的内存数据，以二进制的方式写入磁盘；
- **文件追加方式**（AOF, Append Only File），记录所有的操作命令，并以文本的形式追加到文件中；
- **混合持久化方式**，Redis 4.0 之后新增的方式，混合持久化是结合了 RDB 和 AOF 的优点，在写入的时候，先把当前的数据以 RDB 的形式写入文件的开头，再将后续的操作命令以 AOF 的格式存入文件，这样既能保证 Redis 重启时的速度，又能减低数据丢失的风险。

##### RDB 持久化
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
   config get dbfilename
   config get dir
   config set dir "" # 设置持久化路径
   config set save "" # 禁止持久化
   ```
   
4. RDB 优缺点

   优点
   
   - RDB 的内容为二进制的数据，占用内存小，更紧凑，适合作为备份文件
   - RDB 备份使用的是子进程进行数据持久化至磁盘，不会影响住进程

   缺点

   - 只能保存某个时间间隔的数据，中途Redis服务意外终止，数据会丢失
   - fork()子进程会将数据持久化至磁盘，数据集很大，持久化CPU性能不佳，会影响到住进程



