package controllers

import (
	"fmt"
	"myBrookServer/models"
	"myBrookServer/server"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RemoteStartService struct {
	beego.Controller
}

//远程开启服务api 注册对接用（弃用）
func (c *RemoteStartService) StartService() {
	// UProt := c.GetString("uprot")
	// UName := c.GetString("uname")
	// UPasswd := c.GetString("upasswd")
	uid, _ := c.GetInt("uid")

	remote_u := c.GetString("remote_u")
	remote_p := c.GetString("remote_p")

	o := orm.NewOrm()

	sysMap := make(orm.Params)
	o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

	//验证 远程帐号密码
	if remote_u == sysMap["remote_u"] && remote_p == sysMap["remote_p"] {

		//查询当前服务器类型
		lpBrookServer := models.GetThisServerInfor()

		switch lpBrookServer.Type {
		case 1:
			//Brook服务
			fmt.Println("远程开启Brook")
			server.OpenBrookServer(uid)

		case 2:
			//Socks5服务
			fmt.Println("远程开启Socks5")
			server.OpenSocks5Server(uid)

		case 3, 4:
			//Ws服务
			fmt.Println("远程开启Ws")
			server.OpenWsServer(uid)
		default:
			fmt.Println("服务器关闭！")
		}

		c.Data["json"] = map[string]interface{}{"success": 0, "message": "ok," + strconv.Itoa(lpBrookServer.Type)}

	} else {
		c.Data["json"] = map[string]interface{}{"success": 1, "message": "remote_u and remote_p error"}
	}

	c.ServeJSON()

}

//修改密码api
func (c *RemoteStartService) UpdataServicePasswd() {
	// UProt := c.GetString("uprot")
	// UName := c.GetString("uname")
	// UPasswd := c.GetString("upasswd")
	userId, _ := c.GetInt("user_id")

	remote_u := c.GetString("remote_u")
	remote_p := c.GetString("remote_p")

	o := orm.NewOrm()

	sysMap := make(orm.Params)
	o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

	//验证 远程帐号密码
	if remote_u == sysMap["remote_u"] && remote_p == sysMap["remote_p"] {

		//查询当前服务器类型
		lpBrookServer := models.GetThisServerInfor()
		//获取用户信息
		userInfo, err := models.GetLpBrookUserById(userId)
		if err != nil {
			panic("数据库查询失败-" + err.Error())
		}

		portStr := fmt.Sprintf("%v", userInfo.Port)
		switch lpBrookServer.Type {
		case 1:
			//Brook服务
			fmt.Println("远程修改Brook密码", portStr)
			server.ShutdownBrookByProt(portStr) //关闭服务
			server.OpenBrookServer(userId)      //开启服务

		case 2:
			//Socks5服务
			fmt.Println("远程修改Socks5密码", portStr)
			server.ShutdownSocks5ByProt(portStr) //关闭服务
			server.OpenSocks5Server(userId)      //开启服务

		case 3, 4:
			//Ws服务
			fmt.Println("远程修改Ws密码", portStr)
			server.ShutdownWsByProt(portStr) //关闭服务
			server.OpenWsServer(userId)      //开启服务
		default:
			panic("服务器关闭！")
		}

		c.Data["json"] = map[string]interface{}{"success": 0, "message": "ok," + strconv.Itoa(lpBrookServer.Type)}

	} else {
		c.Data["json"] = map[string]interface{}{"success": 1, "message": "remote_u and remote_p error"}
	}

	c.ServeJSON()

}

//根据uid关闭服务
func (c *RemoteStartService) CloseServer() {
	userId, _ := c.GetInt("user_id")

	remote_u := c.GetString("remote_u")
	remote_p := c.GetString("remote_p")

	o := orm.NewOrm()

	sysMap := make(orm.Params)
	o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

	//验证 远程帐号密码
	if remote_u == sysMap["remote_u"] && remote_p == sysMap["remote_p"] {

		//查询当前服务器类型
		lpBrookServer := models.GetThisServerInfor()
		//获取用户信息
		userInfo, err := models.GetLpBrookUserById(userId)
		if err != nil {
			panic("数据库查询失败-" + err.Error())
		}

		portStr := fmt.Sprintf("%v", userInfo.Port)
		switch lpBrookServer.Type {
		case 1:
			//Brook服务
			fmt.Println("远程关闭Brook密码", portStr)
			server.ShutdownBrookByProt(portStr) //关闭服务
		case 2:
			//Socks5服务
			fmt.Println("远程修改Socks5密码", portStr)
			server.ShutdownSocks5ByProt(portStr) //关闭服务

		case 3, 4:
			//Ws服务
			fmt.Println("远程修改Ws密码", portStr)
			server.ShutdownWsByProt(portStr) //关闭服务
		default:
			panic("服务器关闭！")
		}

		c.Data["json"] = map[string]interface{}{"success": 0, "message": "ok," + strconv.Itoa(lpBrookServer.Type)}

	} else {
		c.Data["json"] = map[string]interface{}{"success": 1, "message": "remote_u and remote_p error"}
	}

	c.ServeJSON()

}

//根据uid开启服务
func (c *RemoteStartService) OpenServer() {
	userId, _ := c.GetInt("user_id")

	remote_u := c.GetString("remote_u")
	remote_p := c.GetString("remote_p")

	o := orm.NewOrm()

	sysMap := make(orm.Params)
	o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

	//验证 远程帐号密码
	if remote_u == sysMap["remote_u"] && remote_p == sysMap["remote_p"] {

		//查询当前服务器类型
		lpBrookServer := models.GetThisServerInfor()
		//获取用户信息
		userInfo, err := models.GetLpBrookUserById(userId)
		if err != nil {
			panic("数据库查询失败-" + err.Error())
		}

		portStr := fmt.Sprintf("%v", userInfo.Port)
		switch lpBrookServer.Type {
		case 1:
			//Brook服务
			fmt.Println("远程修改Brook密码", portStr)
			server.OpenBrookServer(userId) //开启服务

		case 2:
			//Socks5服务
			fmt.Println("远程修改Socks5密码", portStr)
			server.OpenSocks5Server(userId) //开启服务

		case 3, 4:
			//Ws服务
			fmt.Println("远程修改Ws密码", portStr)
			server.OpenWsServer(userId) //开启服务
		default:
			panic("服务器关闭！")
		}

		c.Data["json"] = map[string]interface{}{"success": 0, "message": "ok," + strconv.Itoa(lpBrookServer.Type)}

	} else {
		c.Data["json"] = map[string]interface{}{"success": 1, "message": "remote_u and remote_p error"}
	}

	c.ServeJSON()

}
