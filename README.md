## GoIm


<p align="center">
<img src="https://img.shields.io/badge/license-MIT-green" />
</p>
<br/>
<br/>

> 这是一个由golang+vue编写的web IM应用服务端 📦📦📦

#### 简单的功能

   - [x] 支持账号密码、微博登录
   - [x] 端对端消息推送、图片发送、表情包、语音发送
   - [ ] 视频功能
   - [x] rabbitmq 离线消息推送
   - [x] 创建群聊
   - [x] 群聊消息
   - [x] 响应式的前端界面支持pc与h5
   - [x] 严禁网络不良用语、过滤敏感词汇
   - [x] 支持tcp命令行登录
   - [ ] 数据限流
   - [ ] 支持tcp&websocket数据交互(这点不难 解决方案待调整 用户登录时更新设备端的状态)
   - [ ] 服务集群(不同节点用户使用grpc协议通讯、redis存储服务节点数据 )

   


> 一些库的使用。

 * 支持cors跨域
 * 集成mysql、redis、协程池
 * jwt签名认证
 * zap日志收集
 * viper配置文件解析
 * swag接口文档生成
 * rabbitmq存储离线消息
 

   
   

#### 架构梳理
![](docs/WechatIMG533.png)


#### web登录 效果图
![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/aajXTvR3GF.png!large)

![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/2tVT1ndyTS.png!large)

![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/3Gg8G6wca9.png!large)

 ![](https://cdn.learnku.com/uploads/images/202108/14/32593/XnIO6j3QEr.jpg!large)
 
![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/8p1uALKM18.png!large)

#### [前端源码](https://github.com/pl1998/web-im-app)



#### 
  * [应用部署](/docs/1.部署文档.md)

