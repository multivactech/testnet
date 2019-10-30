# MultiVAC 测试网挖矿客户端说明

本仓库提供了MultiVAC测试网挖矿客户端程序以及相关工具和挖矿教程

- [English Version 英文版](README_EN.md)

[![link](https://img.shields.io/badge/link-browser-red)](http://test.e.mtv.ac)
[![version](https://img.shields.io/badge/version-testnet3.0-blue)](http://test.e.mtv.ac)
[![contact-tel](https://img.shields.io/badge/contact-telegram-blue)](https://t.me/joinchat/I-io_BT_CZjznBGo90vdRA)
[![contact-wechat](https://img.shields.io/badge/contact-wechat%3Amtv2018001-brightgreen)](mtv2018001)

## 文件目录

- tools：包含地址转换工具，用于将MultiVAC testnet2.0地址转换成testnet3.0地址
- client：包含不同系统平台的挖矿程序下载链接

## 地址转换

如果你使用的是旧版本地址，请先使用[地址转换工具](tools/README.md)将旧地址转换成适配`testnet3.0`版本的新地址

## 下载客户端

进入本仓库[client](client/README.md)文件夹对应的系统平台子目录中可以找到下载链接

或者也可以通过命令行下载：

- Mac 版本
```bash
$ wget url
```

- Linux 版本
```bash
$ wget url
```

## 硬件配置
为保证程序正常运行，请确保主机不低于以下配置：

项目 | 最低要求
---|---
CPU | 2核
运行内存 | 2GB
磁盘空间 | 8GB
网络带宽 | 10Mbps/s

## 运行客户端

在终端（命令行）中进入MultiVAC客户端所在的文件夹并输入下面命令运行客户端：

- 修改执行权限
```bash
$ chmod +x MultiVAC
```

- 启动运行
```bash
$ ./MultiVAC --listen="ip:port" --sk="your_sk"
```

### 参数

参数名称 | 释义
---|---
listen | 程序监听的IP地址和端口号（本机局域网IP地址，命令行使用ifconfig查看）
sk | MultiVAC抵押账户私钥

### 注意
MultiVAC客户端不会大量消耗主机的存储和计算资源，但会依赖于网络环境，请确保为主机提供充足的网络环境和资源

## FAQ

### 如何判定节点成功加入了网络？
日志关键词：
- **Accept a new block, shard: xx, height: xxx**
- **I'm potential leader for round xxx**

当终端连续输出上述信息时，代表节点正在为某分片服务，也意味着成功加入了网络

### 如何判定节点没有加入网络？
如果长时间连续出现如下图所示的日志情况，则可能是由于网络环境不好导致节点落后，请尝试重启节点，如果多次重启仍然无法解决，请联系客服

![image](https://note.youdao.com/yws/public/resource/7fdcb5cc8b2e8eb70072e13d14205b1d/xmlnote/51CE068C618D4F6B99ED59C9F51EE2F1/11915)

### 如何判定节点出块了？
在节点加入网络的提前下，会有一定概率出块，请耐心等待。出块提示如下：
\==============================================

💰💰 Mined a new block, shard: x, height: xxx

\==============================================


### 没有抵押证明？
如果配置的节点没有抵押证明，那么一段时间客户端将会退出，请联系客服人员确认抵押信息
