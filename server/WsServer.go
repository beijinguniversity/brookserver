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

//WSServerMap 用来保存所有Ws服务
type WSServerMap struct {
	sync.RWMutex
	Map map[string]*brook.WSServer
}

var (
	//WSServerMapV 全局变量
	WSServerMapV WSServerMap
)

func init() {
	WSServerMapV.Map = make(map[string]*brook.WSServer)
}

//InitWsServerList 初始化Ws服务
func InitWsServerList() {
	fmt.Println("初始化WsServer服务～～～")
	//获取所有用户
	lpBrookUserusers, err := models.GetLpBrookUserAll()
	if err == nil {
		for _, v := range lpBrookUserusers {
			if v.Flow > 0 && v.ExpireTime.Unix() <= time.Now().Unix() && v.IsAdmin != -1 {
				OpenWsServer(v.Id)
			}
		}
	} else {
		panic("初始化Brook服务失败，请检查数据库:(")
	}
}

//ShutdownWsByProt 根据端口关闭服务
func ShutdownWsByProt(port string) error {
	wsServer := WSServerMapV.Map[port]
	if wsServer == nil {
		return errors.New("未找到服务器:(")
	}
	err := wsServer.Shutdown()
	if err == nil {
		WSServerMapV.RLock()
		delete(WSServerMapV.Map, port) //删除
		WSServerMapV.RUnlock()
	}
	return err
}

//OpenWsServer 根据用户id开启一个服务 保存到map中
func OpenWsServer(uid int) error {
	lpBrookUser, err := models.GetLpBrookUserById(uid)
	if err != nil {
		panic("Ws服务器开启时失败-" + err.Error())
	}
	if lpBrookUser == nil {
		panic("未知的用户di-" + fmt.Sprintf("%v", uid))
	}
	serverInfo := models.GetThisServerInfor() //获取当前服务器配置

	path := ""
	if serverInfo.Type == 4 { //wss模式 空字符串默认为ws
		path = "/wss"
	}

	portStr := fmt.Sprintf("%v", lpBrookUser.Port)
	code := utils.CheckPort(portStr) //判断linux端口是否被占用
	if code == 0 {
		fmt.Println("Open:", portStr)
		// addr, password, domain, path string, tcpTimeout, udpTimeout int
		wsServer, err := brook.NewWSServer(":"+portStr, lpBrookUser.ProxyPasswd, serverInfo.Domain, path, 0, 0) //创建服务
		if err != nil {
			panic("Ws服务器开启时失败-" + err.Error())
		}
		WSServerMapV.RLock()
		WSServerMapV.Map[portStr] = wsServer //监听服务
		WSServerMapV.RUnlock()
		go wsServer.ListenAndServe()
		return nil
	} else {
		return errors.New("当前端口被占用")
	}

}
