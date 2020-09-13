package linux

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// 创建产品1并实现产品接口
type Linux struct {
}

func (l Linux) RunCommand(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:" + cmd)
	command := exec.Command("/bin/sh", "-c", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := command.StdoutPipe()
	command.Stderr = command.Stdout

	if err != nil {
		return "", err
	}

	if err = command.Start(); err != nil {
		return "", err
	}

	var sb strings.Builder
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		sb.Write(tmp)
		if err != nil {
			break
		}
	}

	if err = command.Wait(); err != nil {
		return sb.String(), err
	}
	return sb.String(), nil
}

func (l Linux) CheckProRunning(serviceName string) (bool, error) {
	a := `ps aux | awk '/` + serviceName + `/ && !/awk/ {print $2}'`
	pid, err := Linux.RunCommand(Linux{}, a)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}

func (l Linux) GetPID(threadName string) (int, error) {
	a := `ps aux | awk '/` + threadName + `/ && !/awk/ {print $2}'`
	pid, err := Linux.RunCommand(Linux{}, a)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(pid)
}