package server

import (
	"errors"
	"fmt"
	"myBrookServer/models"
	"myBrookServer/utils"
	"sync"
	"time"

	"github.com/txthinking/brook"
)

//Socks5ServerMap 用来保存所有Socks5服务
type Socks5ServerMap struct {
	sync.RWMutex
	Map map[string]*brook.Socks5Server
}

var (
	//Socks5ServerMapV 全局变量
	Socks5ServerMapV Socks5ServerMap
)

func init() {
	Socks5ServerMapV.Map = make(map[string]*brook.Socks5Server)
}

//InitSocks5ServerList 初始化Socks5服务
func InitSocks5ServerList() {
	fmt.Println("初始化Socks5Server服务～～～")
	//获取所有用户
	lpBrookUserusers, err := models.GetLpBrookUserAll()
	if err == nil {
		for _, v := range lpBrookUserusers {
			if v.Flow > 0 && v.ExpireTime.Unix() <= time.Now().Unix() && v.IsAdmin != -1 {
				OpenSocks5Server(v.Id)
			}
		}
	} else {
		fmt.Println("初始化Socks5Server服务失败，请检查数据库:(")
	}
}

//ShutdownSocks5ByProt 根据端口关闭服务
func ShutdownSocks5ByProt(port string) error {
	socks5Server := Socks5ServerMapV.Map[port]
	if socks5Server == nil {
		return errors.New("未找到服务器:(")
	}
	err := socks5Server.Shutdown()
	if err == nil {
		Socks5ServerMapV.RLock()
		delete(Socks5ServerMapV.Map, port) //删除
		Socks5ServerMapV.RUnlock()
	}
	return err
}

// //ListenAndServeSocks5ByProt 根据端口开启服务
// func ListenAndServeSocks5ByProt(port string) error {
// 	socks5Server := Socks5ServerMapV.Map[port]
// 	if socks5Server == nil {
// 		return errors.New("未找到服务器:(")
// 	}
// 	err := socks5Server.ListenAndServe()
// 	if err == nil {
// 		Socks5ServerMapV.RLock()
// 		Socks5ServerMapV.Map[port] = socks5Server //监听服务
// 		Socks5ServerMapV.RUnlock()
// 	}
// 	return err
// }

//OpenSocks5Server 根据用户id开启一个服务 保存到map中
func OpenSocks5Server(uid int) error {
	lpBrookUser, err := models.GetLpBrookUserById(uid)
	if err != nil {
		panic("Socks5服务器开启时失败-" + err.Error())
	}
	if lpBrookUser == nil {
		panic("未知的用户di-" + fmt.Sprintf("%v", uid))
	}
	portStr := fmt.Sprintf("%v", lpBrookUser.Port)
	code := utils.CheckPort(portStr) //判断linux端口是否被占用
	if code == 0 {
		fmt.Println("Open:", portStr)
		socks5Server, err := brook.NewSocks5Server(":"+portStr, "127.0.0.1", lpBrookUser.Name, lpBrookUser.Passwd, 0, 0) //创建服务
		if err != nil {
			panic("Socks5服务器开启时失败-" + err.Error())
		}
		Socks5ServerMapV.RLock()
		Socks5ServerMapV.Map[portStr] = socks5Server //监听服务
		Socks5ServerMapV.RUnlock()
		go socks5Server.ListenAndServe()
		return nil
	} else {
		return errors.New("当前端口被占用")
	}

}
