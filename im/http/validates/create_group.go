/**
  @author:panliang
  @data:2021/8/19
  @note
**/
package validates

type CreateGroupParams struct {
	UserId map[string]string  `json:"user_id"`
	GroupName string `json:"group_name"`
}

