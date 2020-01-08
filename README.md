# MultiVAC 测试网挖矿客户端说明(版本：Oracle)

本仓库提供了MultiVAC测试网挖矿客户端程序以及相关工具和挖矿教程

- [English Version 英文版](README_EN.md)

[![link](https://img.shields.io/badge/link-browser-red)](http://test.e.mtv.ac)
[![version](https://img.shields.io/badge/version-testnet3.0-blue)](http://test.e.mtv.ac)
[![contact-tel](https://img.shields.io/badge/contact-telegram-blue)](https://t.me/joinchat/I-io_BT_CZjznBGo90vdRA)
<!-- [![contact-wechat](https://img.shields.io/badge/contact-wechat-brightgreen)](https://s2.ax1x.com/2019/10/30/K4Rqne.jpg) -->

<!-- ## 文件目录

- tools：包含地址转换工具，用于将MultiVAC testnet2.0地址转换成testnet3.0地址
- client：包含不同系统平台的挖矿程序下载链接

## 登记挖矿地址

### 新地址生成

如果您没有MultiVAC地址，点击[地址申请](http://test.e.mtv.ac/#/wallet/create)申请新的地址

### 旧地址转换方法

如果你使用的是旧版本地址，请先使用[地址转换工具](tools/addressconversion/README.md)将旧地址转换成适配`testnet3.0`版本的新地址

### 填写表单

生成地址后，请将地址填入表单，便于我们提前为大家抵押100万MTV用于内测：
- [中文表单](http://mtvmining.va.mikecrm.com/rZqHF3o)
- [English form](http://mtvmining.va.mikecrm.com/yIMw0Jn)

活动开启后（10月31日14：00 UTC+8），不能使用自己的地址参与测试，请联系客服人员获取私钥参与测试 -->

## 获取测试专属私钥

您可以通过联系MultiVAC运营同学获取专属私钥

## 下载客户端

进入本仓库[client](client/README.md)文件夹对应的系统平台子目录中可以找到下载链接

<!-- 或者也可以通过命令行下载：

- Mac 版本
```bash
$ wget https://multivac-hk.s3.ap-east-1.amazonaws.com/download/mac/MultiVAC
```

- Linux 版本
```bash
$ wget https://multivac-hk.s3.ap-east-1.amazonaws.com/download/linux/MultiVAC
``` -->

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

1. 删除遗留文件
如果你曾经运行过以往的客户端版本我们建议首先删除遗留文件：

mac os：
```bash
$ rm -rf /Users/xxx/Library/Application\ Support/Multivac
```

linux：
```bash
$ rm -rf /home/mac/.multivac
```

注：`xxx`是你系统的用户名

2. 修改执行权限
```bash
$ chmod +x MultiVAC
```

3. 启动运行
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
- Accept a new block, shard: xx, height: xxx
- I'm potential leader for round xxx

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

### MAC OS无法验证开发者?
请在终端输入：sudo xattr -r -d com.apple.quarantine  “MultiVAC路径”
