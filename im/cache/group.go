/**
  @author:panliang
  @data:2021/9/16
  @note
**/
package cache

import "strconv"

func getGroupIdsStr(group_id int) string {
	return "group_ids_"+strconv.Itoa(group_id)
}

func delGroupId()  {
	
}