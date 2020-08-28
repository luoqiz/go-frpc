package frp

import (
	"os/exec"
)

func CheckStatus() {
	exec.Command("tasklist", "|", "findstr", "frp")
}

func Start() {

}
