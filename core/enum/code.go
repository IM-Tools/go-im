/**
  @author:panliang
  @data:2021/12/14
  @note
**/
package enum

const (
	StatusSuccess = 200 //接口正常响应
	StatusError   = 500 //服务器内部错误

	//ws消息状态码
	MESSAGE_OK         = 200  //单聊消息
	DELETE_FRIEND      = 1001 //删除好友
	ADD_FRIEND_REQUEST = 1002 //添加好友
	ADD_FRIEDN_ERROR   = 1003 //拒绝好友
	ADD_FRIEDN_SUCCESS = 1004 //添加好友成功
)
