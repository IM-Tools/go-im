## GO-IM


<p align="center">
<img src="https://img.shields.io/badge/license-MIT-green" />
</p>
<br/>
<br/>

<a  href="https://github.com/IM-Tools/Im-Services">旧版本停止维护⚠️ 新版本地址：https://github.com/IM-Tools/Im-Services</a>

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



![](https://img-blog.csdnimg.cn/622dd0dcb8de42d3bd2abb6e4a583c92.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5r2c6Juw,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center)


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

![软图](https://img-blog.csdnimg.cn/280378527ad34e0caac69f6e696c21c0.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5r2c6Juw,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center)
![软图](https://img-blog.csdnimg.cn/b04c1fb7ec244a0ea5923199ef4743c5.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5r2c6Juw,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center)
![软图](https://img-blog.csdnimg.cn/446c7f1cb8384d2a8f34ff053e43f201.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5r2c6Juw,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center)

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