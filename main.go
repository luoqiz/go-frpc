package main

import (
	"fmt"
	"go-frpc/ui"
	"os/exec"
	"runtime"
	"syscall"
)

func main()  {
	// 查看frp进程
	//frpStatus := exec.Command("D:\\software\\frp_0.33.0_windows_amd64\\frpc.exe", "-c","D:\\software\\frp_0.33.0_windows_amd64\\frpc.ini")

	//RunCommand("D:\\software\\frp_0.33.0_windows_amd64\\frpc.exe", "-c","D:\\software\\frp_0.33.0_windows_amd64\\frpc.ini")
	cmd := exec.Command("cmd.exe", "/c", "tasklist", "|", "findstr", "frp")
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	//Start执行不会等待命令完成，Run会阻塞等待命令完成。
	//err := cmd.Start()
	//err := cmd.Run()
	//cmd.Output()
	//函数的功能是运行命令并返回其标准输出。
	buf, err := cmd.Output()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			status := exitErr.Sys().(syscall.WaitStatus)
			switch {
			case status.Exited():
				fmt.Printf("Return exit error: exit code=%d\n", status.ExitStatus())
			case status.Signaled():
				fmt.Printf("Return exit error: signal code=%d\n", status.Signal())
			}
		} else {
			fmt.Printf("Return other error: %s\n", err)
		}
	} else {
		fmt.Printf("Return OK\n")
		fmt.Println(string(buf))
	}

	//frpStatus.Start()
	//
	//cmd := exec.Command("D:\\software\\frp_0.33.0_windows_amd64\\frpc.exe", "-c","D:\\software\\frp_0.33.0_windows_amd64\\frpc.ini")
	//cmd.Start()
	//StartOtherApp("D:\\software\\frp_0.33.0_windows_amd64\\frpc.exe","-c D:\\software\\frp_0.33.0_windows_amd64\\frpc.ini",true)
	ui.MainWindow()
}