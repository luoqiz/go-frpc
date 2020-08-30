package frp

import (
	"fmt"
	"go-frpc/utils"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

//检查frp状态，未启动返回0，运行中返回其pid
func CheckStatus() int {
	out, _ := utils.ExecCmd("cmd.exe", "/c tasklist | findstr frp")

	if out == "" {
		println("frp 未启动")
		return 0
	}
	fmt.Print(out)
	re, _ := regexp.Compile(`\d+`)
	//查找符合正则的第一个
	all := re.FindAll([]byte(out), -1)
	pid := 0
	for index, item := range all {
		if index == 0 {
			pid, _ = strconv.Atoi(string(item))
			break
		}

	}
	return pid
}

//开启frpc
func Start() {
	workspace, _ := filepath.Abs("")
	frpc := utils.GetDirectory(workspace + "/frpc")[0]
	frpStatus := exec.Command(frpc+"/frpc.exe", "-c", frpc+"/frpc.ini")
	frpStatus.Start()
	println("frpc start success...")
}

// 关闭frpc
func Stop() {
	//pid := CheckStatus()
	out, _ := utils.ExecCmd("cmd.exe", "/c taskkill /f /im frpc.exe")
	if out == "" {
		println("frp 未启动")
	}
	fmt.Print(out)
	println("frpc closed...")
}

// 下载frp
func Download() {
	println("frp download start...")
	workDir, _ := filepath.Abs("")
	utils.Download("https://github.com/fatedier/frp/releases/download/v0.33.0/frp_0.33.0_windows_amd64.zip", workDir+"/frp.zip")
	utils.UnZip(workDir+"/frp.zip", workDir+"/frpcc", true)
	println("frp download end...")
}
