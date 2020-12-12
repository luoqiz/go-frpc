package cmd

import (
	"go-frpc/src/utils/cmd/linux"
	"go-frpc/src/utils/cmd/windows"
	"runtime"
)

type CMDInterface interface {
	//系统执行命令并获取返回值
	RunCommand(cmd string) (string, error)

	//根据进程名判断进程是否运行
	CheckProRunning(serviceName string) (bool, error)

	//根据进程名称获取进程ID
	GetPID(threadName string) (int, error)

	// 后台运行
	RunCommandBg(cmd string)
}

// 创建工厂结构体并实现工厂接口
type CMDFactory struct {
}

func (i CMDFactory) Generate() CMDInterface {
	switch runtime.GOOS {
	case "windows":
		return windows.Windows{}
	case "linux":
		return linux.Linux{}
	default:
		return windows.Windows{}
	}
}

//func ExecCmd(bash, args string) (string, error) {
//	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
//	//cmd := exec.Command("/bin/bash", "-c", s)
//	//cmd := exec.Command("cmd.exe", "/c", "tasklist", "|", "findstr", "cmd")
//	cmd := exec.Command("cmd.exe", strings.Split(args, " ")...)
//	if runtime.GOOS == "windows" {
//		//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
//	}
//	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
//	var out bytes.Buffer
//	cmd.Stdout = &out
//
//	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
//	err := cmd.Run()
//
//	return out.String(), err
//}

//func RunCommand(cmd string) (string, error) {
//	if runtime.GOOS == "windows" {
//		return runInWindows(cmd)
//	} else {
//		return runInLinux(cmd)
//	}
//}
