package exe

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

//ExecuteInfo  可执行程序信息
type ExecuteInfo struct {
	AppPath string
}

var executeInfo *ExecuteInfo

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	executeInfo = &ExecuteInfo{pwd}
}

func getExeutablePath() string {
	file, _ := exec.LookPath(os.Args[0])
	ph, _ := filepath.Abs(file)
	index := strings.LastIndex(ph, string(os.PathSeparator))

	return path.Dir(ph[:index])
}

//Info 获取可执行程序信息
func Info() *ExecuteInfo {
	return executeInfo
}


