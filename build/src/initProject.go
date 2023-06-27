package src

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"x-project/build/types"
)

func InitProject(baseData *types.BaseData) {
	_, _ = color.New(color.FgYellow).Println("开始执行项目 go mod 初始化操作 自动升级版本库")
	cmd := exec.Command("go", "env")
	output, err := cmd.Output()
	if err != nil {
		_, _ = color.New(color.FgRed).Println("没有检测到可执行的 go 指令")
		return
	}
	_, _ = color.New(color.FgBlue).Printf("检测到可用环境----------------------------------\n%s\n", output)
	wd, _ := os.Getwd()
	var tidyCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		tidyCmd = exec.Command("cmd.exe", "/C", fmt.Sprintf("cd %s/%s && go get -u && go mod tidy", wd, baseData.ProjectName))
	} else {
		tidyCmd = exec.Command("bash", "-c", fmt.Sprintf("cd %s/%s && go get -u && go mod tidy", wd, baseData.ProjectName))
	}
	output, err = tidyCmd.Output()
	if err != nil {
		_, _ = color.New(color.FgRed).Println("执行失败", err)
	}
	_, _ = color.New(color.FgYellow).Println("处理 go mod 初始化操作完成")
}
