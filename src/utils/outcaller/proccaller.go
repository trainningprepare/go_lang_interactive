package outcaller

import (
	"os/exec"
)

func exeCall(exePath string,args...string) error  {
	cmd:=exec.Command(exePath,args...)
	return cmd.Run()
}