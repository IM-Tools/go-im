/**
  @author:panliang
  @data:2021/11/16
  @note
**/
package message

type Messages struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DeviceType int    `json:"device_type"`
}
