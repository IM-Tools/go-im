/**
  @author:panliang
  @data:2021/8/19
  @note
**/
package validates

type CreateGroupParams struct {
	UserId map[string]string  `valid:"user_id"`
	GroupName string `valid:"group_name"`
}

