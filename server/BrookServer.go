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

//BrookServerMap 用来保存所有Brook服务
type BrookServerMap struct {
	sync.RWMutex
	Map map[string]*brook.Server
}

var (
	//BrookServerMapV 全局变量
	BrookServerMapV BrookServerMap
)

func init() {
	BrookServerMapV.Map = make(map[string]*brook.Server)
}

//InitBrookServer 初始化Brook服务
func InitBrookServer() {
	fmt.Println("初始化Brook服务～～～")
	//获取所有用户
	lpBrookUserusers, err := models.GetLpBrookUserAll()
	if err == nil {
		for _, v := range lpBrookUserusers {
			if v.Flow > 0 && v.ExpireTime.Unix() <= time.Now().Unix() && v.IsAdmin != -1 {
				OpenBrookServer(v.Id)
			}
		}
	} else {
		panic("初始化Brook服务失败，请检查数据库:(")
	}
}

//ShutdownBrookByProt 根据端口关闭服务
func ShutdownBrookByProt(port string) error {
	brookServer := BrookServerMapV.Map[port]
	if brookServer == nil {
		return errors.New("未找到服务器:(")
	}
	err := brookServer.Shutdown()
	if err == nil {
		BrookServerMapV.RLock()
		delete(BrookServerMapV.Map, port) //删除
		BrookServerMapV.RUnlock()
	}
	return err
}

//OpenBrookServer 根据用户id开启一个服务 保存到map中
func OpenBrookServer(uid int) error {
	lpBrookUser, err := models.GetLpBrookUserById(uid)
	if err != nil {
		panic("brook服务器开启时失败-" + err.Error())
	}
	if lpBrookUser == nil {
		panic("未知的用户di-" + fmt.Sprintf("%v", uid))
	}
	portStr := fmt.Sprintf("%v", lpBrookUser.Port)
	code := utils.CheckPort(portStr) //判断linux端口是否被占用
	if code == 0 {
		fmt.Println("Open:", portStr)
		brookServer, err := brook.NewServer(":"+portStr, lpBrookUser.Passwd, 0, 0) //创建服务
		if err != nil {
			panic("brook服务器开启时失败-" + err.Error())
		}
		BrookServerMapV.RLock()
		BrookServerMapV.Map[portStr] = brookServer //监听服务
		BrookServerMapV.RUnlock()
		go brookServer.ListenAndServe()
		return nil
	} else {
		return errors.New("当前端口被占用")
	}

}
