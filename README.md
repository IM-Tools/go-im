## GO-IM

<p align="center">
<img src="https://img.shields.io/badge/license-MIT-green" />
</p>
<br/>
<br/>

> 这是一个由golang编写的高性能IM服务器 📦📦📦

#### 支持以下功能

   - [x] 支持账号密码、微博登录
   - [x] 单聊消息、群聊消息、离线消息同步
   - [x] 支持单机部署、集群部署
   - [ ] 多设备登录
   - [x] 客户端：web端、桌面应用

   
> 一些库的使用。

 * 支持cors跨域
 * 集成mysql、redis、协程池
 * jwt签名认证
 * zap日志收集
 * viper配置文件解析
 * swag接口文档生成
 * rabbitmq存储离线消息
 * 集群服务使用grpc向不同服务节点投递消息
 

   
  
#### 架构梳理



![](docs/架构实例图.png)


#### 安装使用

#### 安装redis
```shell
docker pull redis

docker run -p 6379:6379 --name redis
-v /data/redis/redis.conf:/etc/redis/redis.conf
-v /data/redis/data:/data
-d redis redis-server /etc/redis/redis.conf --appendonly yes
```

#### 安装mysql
```shell
docker pull mysql
docker run --name mysqlserver -v $PWD/conf:/etc/mysql/conf.d -v $PWD/logs:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d -i -p 3306:3306 mysql:latest
```
#### 安装rabbitmq
```shell
docker pull rabbitmq
docker run -d --hostname my-rabbit --name rabbit -p 15672:15672 -p 5672:5672 rabbitmq

```
#### 安装项目
```shell
git close https://github.com/IM-Tools/go-im.git
cd go-im
```

#### 配置.env 调整.env文件配置
```shell
cp .env.example .env
```
#### 启动
```shell
go run main.go
```

#### 桌面端

![软图](docs/WechatIMG670.png)
![软图](docs/WechatIMG671.png)
![软图](docs/WechatIMG672.png)

#### web登录 效果图
![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/aajXTvR3GF.png!large)

![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/2tVT1ndyTS.png!large)

![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/3Gg8G6wca9.png!large)

 ![](https://cdn.learnku.com/uploads/images/202108/14/32593/XnIO6j3QEr.jpg!large)
 
![golang+vue3开发的一个im应用](https://cdn.learnku.com/uploads/images/202108/14/32593/8p1uALKM18.png!large)

#### [前端源码](https://github.com/pl1998/web-im-app)
#### [桌面端](暂未开源)



#### 
  * [应用部署](/docs/1.部署文档.md)

#### 学习交流

QQ:2540463097