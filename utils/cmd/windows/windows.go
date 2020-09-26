package windows

import (
	"fmt"
	"go-frpc/utils"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// 创建 windows 并实现产品接口
type Windows struct {
}

func (l Windows) RunCommand(cmd string) (string, error) {
	fmt.Println("Running Windows cmd:" + cmd)
	command := exec.Command("cmd", "/c", cmd)

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

func (l Windows) CheckProRunning(serviceName string) (bool, error) {
	a := `ps aux | awk '/` + serviceName + `/ && !/awk/ {print $2}'`
	pid, err := Windows.RunCommand(Windows{}, a)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}

func (l Windows) GetPID(threadName string) (int, error) {
	pid := 0
	a := `tasklist | findstr ` + threadName
	out, err := Windows{}.RunCommand(a)
	if err != nil {
		return pid, err
	}
	utils.Log.Info(out)
	re, _ := regexp.Compile(`\d+`)
	//查找符合正则的第一个
	all := re.FindAll([]byte(out), -1)
	for index, item := range all {
		if index == 0 {
			pid, _ = strconv.Atoi(string(item))
			break
		}

	}
	return pid, nil
}
