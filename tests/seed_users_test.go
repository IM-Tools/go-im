/**
  @author:panliang
  @data:2021/11/20
  @note
**/
package tests

import (
	"crypto/rand"
	"fmt"
	"im_app/config"
	"im_app/im"
	user2 "im_app/im/http/models/user"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"log"
	"math/big"
	"sync"
	"testing"
	"time"
)

func init()  {
	config.Initialize()
}

var wg sync.WaitGroup
func TestSeedUsers(T *testing.T)  {
	//设置池
	im.SetupPool()
	wg.Add(7)
	go install(6205,10000)
	go install(10001,20000)
	go install(20001,30000)
	go install(30001,40000)
	go install(40001,50000)
	go install(50001,60000)
	go install(60001,70000)
	wg.Wait() //阻塞直到所有任务完成
	fmt.Println("over")

}

func install(start int ,end int)  {
	create_time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	for i:=start;i<end;i++ {
		name := fmt.Sprintf("测试%d",i)
		age := random1()
		user := user2.Users{ID: uint64(i),
			Name:name,
			Avatar:"https://cdn.learnku.com/uploads/avatars/27407_1531668878.png!/both/100x100",
			Password: helpler.HashAndSalt("123456"),
			CreatedAt: create_time,
			Sex: 1,
			Status: 0,
			ClientType: 1,
			Age:age,
			LastLoginTime: create_time,
		}
		model.DB.Create(&user)
	}
}

func random1() int  {
	max := big.NewInt(100)
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal("rand:", err)
	}
	return i.BitLen()
}

