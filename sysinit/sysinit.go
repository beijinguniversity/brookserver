package sysinit

import (
	"fmt"
	"myBrookServer/models"
	"myBrookServer/utils"

	"github.com/astaxie/beego"
)

func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志
	utils.InitLogs()
	//初始化缓存
	utils.InitCache()
	//初始化数据库
	InitDatabase()

	//初始化服务器信息
	initServerInfo()

	//查询当前服务器类型
	lpBrookServer := models.GetThisServerInfor()

	switch lpBrookServer.Type {
	case 1:
		//初始化Brook服务
		InitBrook()
	case 2:
		//初始化Socks5服务
		InitSocks5()
	case 3, 4:
		//初始化Ws服务
		InitWs()
	default:
		fmt.Println("服务器关闭！", lpBrookServer.Type)

	}

	fmt.Println("初始化定时任务")
	//初始化定时任务
	initTask()
}
