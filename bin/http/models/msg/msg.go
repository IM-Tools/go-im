/**
  @author:panliang
  @data:2021/6/30
  @note
**/
package msg

type ImMessage struct {
	ID uint64 `json:"id"`
	Msg string `json:"msg"`
	CreatedAt string `json:"created_at"`
	FromId uint64 `json:"user_id"`
	ToId uint64 `json:"send_id"`
	Channel string `json:"channel"`
}

