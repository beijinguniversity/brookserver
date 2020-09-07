package sysinit

import (
	"fmt"
	"myBrookServer/models"
	"myBrookServer/server"
	"myBrookServer/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

var lpBrookServer models.LpBrookServer

func initServerInfo() {
	o := orm.NewOrm()
	//查询当前服务器类型
	o.QueryTable(models.LpBrookServerTBName()).Filter("Id", beego.AppConfig.String("lp_brook_server_id")).One(&lpBrookServer)
}

func initTask() {
	//初始化一个任务
	tk1 := toolbox.NewTask("tk1", "0/30 * * * * *", uploadFlow)

	//可以测试开启运行
	// err := tk1.Run()
	// if err != nil {
	// 	panic("定时任务执行错误")
	// }
	//加入全局的计划任务列表
	toolbox.AddTask("tk1", tk1)

	//开始执行全局的任务
	toolbox.StartTask()
	// defer toolbox.StopTask()

}

//上传流量
func uploadFlow() error {
	fmt.Println("上传流量...")
	// 获取所有用户
	lpBrookUserArr, err := models.GetLpBrookUserAll()
	if err == nil {

		for _, userInfo := range lpBrookUserArr {
			portStr := fmt.Sprintf("%v", userInfo.Port)
			//判断 输入端口流量控制 是否打开 -1err/0关闭状态/1开启状态
			inputStateCode := utils.IsPortInputFlowControl(portStr)
			if inputStateCode == 0 {
				//打开端口输入流量控制
				utils.OpenPortInputFlowControl(portStr)
			}
			//判断 输出端口流量控制 是否打开 -1err/0关闭状态/1开启状态
			outputStateCode := utils.IsPortOutputFlowControl(portStr)
			if outputStateCode == 0 {
				//打开端口输出流量控制
				utils.OpenPortOutputFlowControl(portStr)
			}

			if userInfo.Flow <= 0 || userInfo.ExpireTime.Unix() <= time.Now().Unix() || userInfo.IsAdmin == -1 {
				if utils.CheckPort(portStr) == 1 {
					fmt.Printf("用户ID:%v 关闭服务 \t\n", userInfo.Id)
					//用户没流量啦，关闭服务哈～
					switch lpBrookServer.Type {
					case 1:
						server.ShutdownBrookByProt(portStr)
					case 2:
						server.ShutdownSocks5ByProt(portStr)
					case 3, 4:
						server.ShutdownWsByProt(portStr)
					default:
						fmt.Println("服务器关闭！", lpBrookServer.Type)
					}
				}
			} else {
				if utils.CheckPort(portStr) == 0 {
					fmt.Printf("用户ID:%v 开启服务 \t\n", userInfo.Id)
					//用户有流量啦，开启服务哈～
					switch lpBrookServer.Type {
					case 1:
						server.OpenBrookServer(userInfo.Id)

					case 2:
						server.OpenSocks5Server(userInfo.Id)

					case 3, 4:
						server.OpenWsServer(userInfo.Id)
					default:
						panic("服务器类型未知的或关闭！")
					}
				}

				flowinputStr := utils.GetPortInputFlowControl(portStr)   // 获取输入流量
				flowoutputStr := utils.GetPortOutputFlowControl(portStr) // 获取输出流量

				flowinputF64, _ := strconv.ParseFloat(flowinputStr, 64)
				flowoutputF64, _ := strconv.ParseFloat(flowoutputStr, 64)

				num := ((flowinputF64 + flowoutputF64) / 1048576) * lpBrookServer.FlowRatio //b -> mb * 流量比例

				//更新用户流量
				models.UpdateUserFlowById(userInfo.Id, num)

				//获取今天的 月_日
				month := int(time.Now().Month())
				day := time.Now().Day()
				month_day := fmt.Sprintf("%v_%v", month, day)

				flowLog := make([]map[string][]float64, 0)
				if err := utils.GetCache("flow_log_"+fmt.Sprintf("%v", userInfo.Id), &flowLog); err != nil { //获取ip数组
					utils.SetCache("flow_log_"+fmt.Sprintf("%v", userInfo.Id), flowLog, 9999999) // 设置缓存
				}

				if len(flowLog) == 0 || len(flowLog) < 7 { //没有数据 或者 数据少于7天(不等于7天)
					//追加今天的空数组
					flowLog = append(flowLog, make(map[string][]float64))
				} else if len(flowLog) == 7 && flowLog[6][month_day] == nil { //有7天的数据 并且 没有今天的数据
					//删除切片的第一个数据
					flowLog = append(flowLog[0:0], flowLog[1:]...)
					//追加今天的空数组
					flowLog = append(flowLog, make(map[string][]float64))
				}

				//今天的数组赋值
				flowLog[len(flowLog)-1][month_day] = append(flowLog[len(flowLog)-1][month_day], num)

				utils.SetCache("flow_log_"+fmt.Sprintf("%v", userInfo.Id), flowLog, 9999999) // 设置缓存

			}

		}

		//重置所有监听流量
		utils.ResetAllFlowControl()

	} else {
		fmt.Println("流量服务开启失败:(    请检查数据库")
	}

	return nil
}
