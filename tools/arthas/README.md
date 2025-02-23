# Arthas

Arthas 是阿里巴巴开源的一款线上监控诊断产品，实时查看应用 load、内存、gc、线程的状态信息，并能在不修改应用代码的情况下，对业务问题进行诊断，包括查看方法调用的出入参、异常，监测方法执行耗时，类加载信息等。  

这里主要是学习和实践 Arthas，[原文参见](https://arthas.aliyun.com/doc/)

## 安装与启动

```sh
curl -O https://arthas.aliyun.com/arthas-boot.jar
java -jar arthas-boot.jar
```

启动后选择要监控的Java进程

## 常用命令

`dashboard`：实时监控系统性能，包括线程、内存、GC等

`thread`：查看线程信息，如`thread -n 3`显示最忙的3个线程

`jad`：反编译类文件，如`jad com.example.MyClass`

`watch`：监控方法调用，如`watch com.example.MyClass myMethod "{params, returnObj}"`

`trace`：追踪方法调用链路，如`trace com.example.MyClass myMethod`

`stack`：查看方法调用栈，如`stack com.example.MyClass myMethod`

`monitor`：统计方法调用次数和耗时，如`monitor -c 5 com.example.MyClass myMethod`


