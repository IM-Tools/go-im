/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package group

import "github.com/golang/protobuf/ptypes/timestamp"

type ImGroup struct {
	ID uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	GroupName string `json:"group_name"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
}
