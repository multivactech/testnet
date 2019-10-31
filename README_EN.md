# MultiVAC Testnet 3.0 Beta Test Guide

**This time’s CBT ends now.**
This asset provides you the tutorial of MultiVAC testnet's mining client and relevant tools.

- [中文版 Chinese version](README.md)

[![link](https://img.shields.io/badge/link-browser-red)](http://test.e.mtv.ac)
[![version](https://img.shields.io/badge/version-testnet3.0-blue)](http://test.e.mtv.ac)
[![contact-tel](https://img.shields.io/badge/contact-telegram-blue)](https://t.me/joinchat/I-io_BT_CZjznBGo90vdRA)
[![contact-wechat](https://img.shields.io/badge/contact-wechat-brightgreen)](https://s2.ax1x.com/2019/10/30/K4Rqne.jpg)

## Files

- Tools：Address swap tool. You can use this tool to change the MultiVAC testnet 2.0's address to the one in testnet 3.0
- Client：Download links for our mining client in different OS

## To register your mining address

### Create a new address

If you don't have a MultiVAC address yet, please click [Address Application](http://test.e.mtv.ac/#/wallet/create) to apply for a new address.

### Swap an old-version address

If you use the old-version address, please use the [Address Swap Tool](tools/addressconversion/README.md) to convert it to the new address suitable for testnet 3.0.

### Fill in the form

After creating the new address, please fill it in the form. We will help you deposit 1 million MTV for the beta test.

- [中文表单](http://mtvmining.va.mikecrm.com/rZqHF3o)
- [English form](http://mtvmining.va.mikecrm.com/yIMw0Jn)

If you want to join our beta test after the start time at 14:00 (UTC+8) on Oct 31, you cannot join through your own address. Please contact our admins to get a private key in order to join the test.

## Configuration

To ensure the client can run steadily, please make sure that your computer is not lower than the following configuration.

Project | The Minimum Requirement
---|---
CPU | 2 Core
ROM | 2 GB
Disk Space| 8 GB
Network Bandwidth | 10Mbps/s


## Download

You can go to the sub-directory of the [client](client/README.md) and find the download link there.

Or download through command-line.

- Mac version
```bash
$ wget https://multivac-hk.s3.ap-east-1.amazonaws.com/download/mac/MultiVAC
```

- Linux version
```bash
$ wget https://multivac-hk.s3.ap-east-1.amazonaws.com/download/linux/MultiVAC
```

## Run the Mining Client

Please first open the file where the MultiVAC mining client is located in the terminal (command line), and use the following command line to run the client (here we take Linux version as an example).

- Modify execution permissions
```bash
$ chmod +x MultiVAC
```

- Start running
```bash
$ ./MultiVAC --listen="ip:port" --sk="your_sk"
```

### Program command parameter
Parameter name | Definition 
---|---
listen | IP address and Port No. monitored by the program (To check your LAN IP address, please use the command line ifconfig).
sk | MultiVAC deposit account private key.

### Notes
The MultiVAC mining client won't consume a lot of storage and computing resources. However, it's dependent on your network environment. Please make sure your network environment and resources are sufficient to run the client.

## FAQ

### How to know that my node has joined the network successfully ?
Keywords in the Log:
- Accept a new block, shard: xx, height: xxx
- I'm potential leader for round xxx

When the terminal keeps outputting the above info, it means the node is serving for a Shard, which indicates it has joined the network successfully.

### How to know that my node failed to join the network?

If you keep receiving the information displayed below, your node might lag behind because of poor network environment. Please try to restart the node first. If multiple attempts still fail to solve the problem, please contact our tech support team.

![image](https://note.youdao.com/yws/public/resource/7fdcb5cc8b2e8eb70072e13d14205b1d/xmlnote/51CE068C618D4F6B99ED59C9F51EE2F1/11915)

### How to know a node has generated a block?

If the node has joined the network successfully, there exists a probability that a block will be generated. Please be patient. Once a block is generated, the below info will show.
\==============================================

💰💰 Mined a new block, shard: x, height: xxx

\==============================================


### No deposit proof? 
If the configured node doesn't have deposit proof, the client will log out after a while. Please contact our support team to confirm the deposit information.