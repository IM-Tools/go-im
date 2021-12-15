/**
  @author:panliang
  @data:2021/12/15
  @note
**/
package service

import "im_app/core/ws"

// 系统单独发送
func SendMessage(code int, f_id int, t_id int, message string) {
	ws.ImManager.SystemMessageDelivery(int64(f_id), &ws.Msg{Code: code, FromId: f_id, Msg: message, ToId: t_id, Status: 0, MsgType: 1,
		ChannelType: 3})
}
