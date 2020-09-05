package routers

import (
	"myBrookServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// beego.Router("/remote/startservice", &controllers.RemoteStartService{}, "Get:StartService")
	beego.Router("/remote/UpdataServicePasswd", &controllers.RemoteStartService{}, "Get:UpdataServicePasswd")
	beego.Router("/remote/CloseServer", &controllers.RemoteStartService{}, "Get:CloseServer")
	beego.Router("/remote/OpenServer", &controllers.RemoteStartService{}, "Get:OpenServer")

}
