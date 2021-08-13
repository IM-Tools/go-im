/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package utils

import "log"

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
