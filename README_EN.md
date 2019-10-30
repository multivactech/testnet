# MultiVAC Testnet 3.0 Beta Test Guide

This asset provides you the tutorial of MultiVAC testnet's mining client and relevant tools.

- [中文版 Chinese version](README.md)

[![link](https://img.shields.io/badge/link-browser-red)](http://test.e.mtv.ac)
[![version](https://img.shields.io/badge/version-testnet3.0-blue)](http://test.e.mtv.ac)
[![contact-tel](https://img.shields.io/badge/contact-telegram-blue)](https://t.me/joinchat/I-io_BT_CZjznBGo90vdRA)
[![contact-wechat](https://img.shields.io/badge/contact-wechat%3Amtv2018001-brightgreen)](mtv2018001)

## Files

- Tools：Address swap tool. You can use this tool to change the MultiVAC testnet 2.0's address to the one in testnet 3.0
- Client：Download links for our mining client in different OS

## Address Conversion

If you use the old-version address, please use the [address swap tool](tools/README.md) to swap it to the new address suitable for testnet 3.0.

## Configuration

To ensure the client can run steadily, please make sure that your computer is not lower than the following configuration:

Project | The Minimum Requirement
---|---
CPU | 2 Core
ROM | 2 GB
Disk Space| 8 GB
Network Bandwidth | 10Mbps/s


## Download

You can go to the sub-directory of the [client](client/README.md) and find the download link there.

Or download through command-line

- Mac version
```bash
$ wget url
```

- Linux version
```bash
$ wget url
```

## Run the Mining Client

Please first open the file where the MultiVAC mining client is located in the terminal (command line), and use the following command line to run the client (here we take Linux version as an example)

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
listen | IP address and Port No. monitored by the program (To check your LAN IP address, please use the command line ifconfig)
sk | MultiVAC deposit account private key

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