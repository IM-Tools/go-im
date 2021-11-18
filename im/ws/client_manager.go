/**
 @author:panliang
 @data:2021/11/18
 @note
**/
package ws

import (
	"strconv"
	"sync"
)

type ImClientManager struct {
	ImClientMap map[string]*ImClient // 存放在线用户连接
	Broadcast   chan []byte          // 收集消息分发到客户端
	Register    chan *ImClient       // 新注册的长连接
	Unregister  chan *ImClient       // 已注销的长连接
}

var mutexKey sync.Mutex

// 客户端管理器

var ImManager = ImClientManager{
	ImClientMap: make(map[string]*ImClient),
	Broadcast:   make(chan []byte),
	Register:    make(chan *ImClient),
	Unregister:  make(chan *ImClient),
}

type ClientHandler interface {
	SetClientInfo(conn *ImClient)            // 设置客户端信息
	DelClient(conn *ImClient)                // 删除客户端信息
	Start()                                  // 启动服务
	ImSend(message []byte, ignore *ImClient) // 给指定客户端投递消息 该方法可能用不着了..
}

func (manager *ImClientManager) SetClientInfo(conn *ImClient) {

	// 关于加锁的问题 可能有更好的方法 以后学会了在优化 先这样吧

	mutexKey.Lock()
	manager.ImClientMap[conn.ID] = &ImClient{ID: conn.ID, Socket: conn.Socket, Send: conn.Send}
	mutexKey.Unlock()

}

func (manager *ImClientManager) DelClient(conn *ImClient) {

	close(conn.Send)
	// 清理map应该要加锁
	mutexKey.Lock()
	delete(manager.ImClientMap, conn.ID)
	mutexKey.Unlock()

}

func (manager *ImClientManager) Start() {
	for {
		select {
		case conn := <-ImManager.Register:

			// 设置客户端信息
			manager.SetClientInfo(conn)

			id, _ := strconv.ParseInt(conn.ID, 10, 64)

			// 用户在线消息下发
			LaunchOnlineMsg(conn.ID, manager)

			// 更新用户在线状态
			PushUserOnlineNotification(conn, id)

		case conn := <-ImManager.Unregister:

			// 设置用户离线
			PushUserOfflineNotification(manager, conn)

		case message := <-ImManager.Broadcast:

			// 消息投递

			LaunchMessage(message, manager)

			// 关于离线消息是否需要存到 mq
		}
	}
}

func (manager *ImClientManager) ImSend(message []byte, ignore *ImClient) {
	data, ok := manager.ImClientMap[ignore.ID]
	if ok {
		data.Send <- message
	}
}
