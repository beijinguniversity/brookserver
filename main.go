package main

import (
	_ "myBrookServer/routers"
	_ "myBrookServer/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()

}
