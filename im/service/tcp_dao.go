/**
  @author:panliang
  @data:2021/9/25
  @note
**/
package service

import (
	userModel "go_im/im/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"net"
)

type TcpDao struct {}

/**
tcp用户登录认证
 */
func (*TcpDao)Login(conn net.Conn,username string,password string) (user userModel.Users,err error)  {
	var users userModel.Users
	model.DB.Model(&userModel.Users{}).Where("name = ?",username).Find(&users)
	if users.ID == 0 {
		conn.Write([]byte(`用户不存在`))
		conn.Close()
		return
	}
	if !helpler.ComparePasswords(users.Password, password) {
		conn.Write([]byte(`账号或者密码错误`))
		conn.Close()
		return
	}
	return users,nil
}
