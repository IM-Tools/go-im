/**
 @author:panliang
 @data:2021/11/18
 @note
**/
package ws

import (
	"im_app/im/cache"
	"im_app/pkg/pool"
	"sync"
)

var node = new(cache.ServiceNode)

type ImClientManager struct {
	ImClientMap map[int64]*ImClient // 存放在线用户连接
	Broadcast   chan []byte          // 收集消息分发到客户端
	Register    chan *ImClient       // 新注册的长连接
	Unregister  chan *ImClient       // 已注销的长连接
}

var mutexKey sync.Mutex

// 客户端管理器

var ImManager = ImClientManager{
	ImClientMap: make(map[int64]*ImClient),
	Broadcast:   make(chan []byte),
	Register:    make(chan *ImClient),
	Unregister:  make(chan *ImClient),
}

// 客户端管理 方法集合

type ClientHandler interface {
	SetClientInfo(conn *ImClient)            // 设置客户端信息
	DelClient(conn *ImClient)                // 删除客户端信息
	Start()                                  // 启动服务
	ImSend(message []byte, ignore *ImClient) // 给指定客户端投递消息 该方法可能用不着了..
	LaunchOnlineMsg(id string)               // 用户上线通知方法
	LaunchMessage(msg_byte []byte)           // 消息下发用户
}

// 客户端方法集合

type ImClientHandler interface {
	PullMessageHandler(message []byte) // 拉取消息处理入队
	ImRead()                           // 读消息
	ImWrite()                          // 写消息
}

// 关于加锁的问题 可能有更好的方法 以后学会了在优化 先这样吧
func (manager *ImClientManager) SetClientInfo(conn *ImClient) {

	mutexKey.Lock()
	manager.ImClientMap[conn.ID] = &ImClient{ID: conn.ID, Socket: conn.Socket, Send: conn.Send}
	mutexKey.Unlock()

}

func (manager *ImClientManager) DelClient(conn *ImClient) {
	close(conn.Send)
	mutexKey.Lock()
	delete(manager.ImClientMap, conn.ID)
	mutexKey.Unlock()

}

func (manager *ImClientManager) Start() {


	for {
		select {
		case conn := <-ImManager.Register:
			manager.SetClientInfo(conn)      // 设置客户端信息
			manager.LaunchOnlineMsg(conn.ID) // 用户在线下发通知
			node.SetUserServiceNode(conn.ID) // 设置用户节点
		case conn := <-ImManager.Unregister:
			PushUserOfflineNotification(manager, conn) // 设置用户离线状态
			node.DelUserServiceNode(conn.ID)          // 删除用户节点
		case message := <-ImManager.Broadcast:
			pool.AntsPool.Submit(func() {
				manager.LaunchMessage(message) // 协程池任务调度 抢占式消息下发
			})


		}
	}
}

func (manager *ImClientManager) ImSend(message []byte, ignore *ImClient) {
	data, ok := manager.ImClientMap[ignore.ID]
	if ok {
		data.Send <- message
	}
}
