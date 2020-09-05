package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

//判断端口是否被占用 0未被占用 1被占用 -1错误
func CheckPort(prot string) int {
	sysType := runtime.GOOS
	if sysType == "linux" {
		//执行命令
		cmd := exec.Command("bash", "-c", "lsof -i :"+prot+" | grep -v ESTABLISHED | tail -n +2")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(output))
			if string(output) == "" {
				return 0
			} else {
				return -1 //错误
			}
		} else {
			if string(output) != "" {
				return 1
			} else {
				return 0
			}
		}

	}

	if sysType == "windows" {
		// windows系统
		return -1
	}

	return -1
}

//判断 输入端口流量控制 是否打开 -1err 0关闭状态 1开启状态
func IsPortInputFlowControl(prot string) int {
	sysType := runtime.GOOS

	if sysType == "linux" {
		f, err := exec.Command("bash", "-c", "iptables -L -v -n -x | grep :"+prot+" | awk '{print $2}'").Output()

		if err != nil {
			return -1
		}

		fmt.Println(string(f))
		if string(f) == "" {
			return 0
		} else {
			return 1
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return -1
	// }
	return -1
}

//判断 输出端口流量控制 是否打开 -1err 0关闭状态 1开启状态
func IsPortOutputFlowControl(prot string) int {
	sysType := runtime.GOOS

	if sysType == "linux" {
		f, err := exec.Command("bash", "-c", "iptables -L -v -n -x | grep spt:"+prot+" | awk '{print $2}'").Output()

		if err != nil {
			return -1
		}

		if string(f) == "" {
			return 0
		} else {
			return 1
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return -1
	// }
	return -1
}

//获取input监听的流量
func GetPortInputFlowControl(prot string) string {
	sysType := runtime.GOOS

	if sysType == "linux" {
		f, err := exec.Command("bash", "-c", "iptables -L -v -n -x | grep dpt:"+prot+" | awk 'END {print}' | awk '{print $2}'").Output()
		if err != nil {
			return "0"
		}

		inputStr := string(f)
		inputStr = strings.Replace(inputStr, "\n", "", -1)
		inputStr = strings.Replace(inputStr, "\t", "", -1)
		return inputStr
	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	fmt.Println("windows 待开发~~~")
	// }
	return "0"
}

//获取output监听的流量
func GetPortOutputFlowControl(prot string) string {
	sysType := runtime.GOOS

	if sysType == "linux" {
		f, err := exec.Command("bash", "-c", "iptables -L -v -n -x | grep spt:"+prot+" | awk 'END {print}' | awk '{print $2}'").Output()
		if err != nil {
			return "0"
		}
		outputStr := string(f)
		outputStr = strings.Replace(outputStr, "\n", "", -1)
		outputStr = strings.Replace(outputStr, "\t", "", -1)
		return outputStr

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	fmt.Println("windows 待开发~~~")
	// }
	return "0"
}

//打开端口输入流量控制
func OpenPortInputFlowControl(prot string) error {
	sysType := runtime.GOOS

	if sysType == "linux" {

		//执行命令 输入监控
		//iptables -A INPUT -p tcp --dport 8080
		if _, err := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", prot).Output(); err != nil {
			return err
		} else {
			return nil
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return errors.New("windows 待开发~~~")
	// }

	return errors.New("未知系统？？？")
}

//打开端口输出流量控制
func OpenPortOutputFlowControl(prot string) error {
	sysType := runtime.GOOS

	if sysType == "linux" {

		//输出监控 输出监控
		//iptables -A OUTPUT -p tcp --sport 8080
		if _, err := exec.Command("iptables", "-A", "OUTPUT", "-p", "tcp", "--sport", prot).Output(); err != nil {
			return err
		} else {
			return nil
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return errors.New("windows 待开发~~~")
	// }

	return errors.New("未知系统？？？")
}

//重置所有输入流量
func ResetInputFlowControl() error {
	sysType := runtime.GOOS

	if sysType == "linux" {

		if _, err := exec.Command("iptables", "-Z", "INPUT").Output(); err != nil {
			return err
		} else {
			return nil
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return errors.New("windows 待开发~~~")
	// }

	return errors.New("未知系统？？？")
}

//重置所有输出流量
func ResetOutputFlowControl() error {
	sysType := runtime.GOOS

	if sysType == "linux" {

		if _, err := exec.Command("iptables", "-Z", "OUTPUT").Output(); err != nil {
			return err
		} else {
			return nil
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return errors.New("windows 待开发~~~")
	// }

	return errors.New("未知系统？？？")
}

//重置所有端口流量
func ResetAllFlowControl() error {
	sysType := runtime.GOOS

	if sysType == "linux" {

		if _, err := exec.Command("iptables", "-Z").Output(); err != nil {
			return err
		} else {
			return nil
		}

	} else {
		panic("不支持的系统:(")
	}

	// if sysType == "windows" {
	// 	// windows系统
	// 	return errors.New("windows 待开发~~~")
	// }
	return nil
}
