/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package group

type ImGroupUsers struct {
	ID uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	CreatedAt string `json:"created_at"`
	GroupId uint64 `json:"group_id"`
	Remark string `json:"remark"`
}

