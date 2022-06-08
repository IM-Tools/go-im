/**
  @author:panliang
  @data:2022/5/26
  @note
**/
package controller

import "strconv"

func StringToInt(str string) int {
	strInt, _ := strconv.Atoi(str)
	return strInt
}
