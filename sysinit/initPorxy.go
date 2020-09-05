package sysinit

import "myBrookServer/server"

//初始化Brook服务
func InitBrook() {
	server.InitBrookServer()
}

//初始化socks5服务
func InitSocks5() {
	server.InitSocks5ServerList()

}

//初始化Ws服务
func InitWs() {
	server.InitWsServerList()

}
