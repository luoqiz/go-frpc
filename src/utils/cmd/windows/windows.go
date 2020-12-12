package windows

import (
	"fmt"
	"go-frpc/src/utils"
	"os/exec"
	"strconv"
	"strings"
)

// 创建 windows 并实现产品接口
type Windows struct {
}

func (l Windows) RunCommandBg(cmd string) {
	fmt.Println("Running windows cmd:" + cmd)
	command := exec.Command("cmd", "/c", cmd)
	//command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	var err error
	if err = command.Start(); err != nil {
		fmt.Printf("%v: Command finished with error: %v\n", "get_time()", err)
		return
	}
	return
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
	a := `wmic process get name,executablepath,processid | findstr ` + threadName
	out, err := Windows{}.RunCommand(a)
	if err != nil {
		return pid, err
	}

	rs := strings.Split(out, "\n")
	utils.Log.Info(rs)
	for _, item := range rs {
		item = utils.DeleteExtraSpace(item)
		item = strings.TrimSpace(item)
		threadInfo := strings.Split(item, " ")
		if len(threadInfo) == 3 {
			tName := threadInfo[1]
			if tName[0:len(tName)-4] == threadName {
				pid, _ = strconv.Atoi(threadInfo[2])
				return pid, nil
			}
		}
	}

	return 0, nil
}
