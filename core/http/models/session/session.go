/**
  @author:panliang
  @data:2021/12/20
  @note
**/
package session

// 用户会话表
type ImSessions struct {
	ID          int64  `json:"id"`
	MId         int64  `json:"m_id"`
	FId         int64  `json:"f_id"`
	TopStatus   int    `json:"top_status"`
	TopTime     string `json:"top_time"`
	Note        string `json:"note"`
	ChannelType int    `json:"channel_type"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"created_at"`
}
