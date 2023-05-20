package src

import (
	"github.com/fatih/color"
	"github.com/gogf/gf/v2/os/gfile"
	"x-project/build/types"
)

func ProjectReName(baseData *types.BaseData) {
	_, _ = color.New(color.FgYellow).Println("开始处理 项目名称替换 " + baseData.ProjectName)
	files, err := gfile.ScanDirFile(baseData.OutPath, "*.go,go.mod,config.yaml", true)
	if err != nil {
		_, _ = color.New(color.FgRed, color.Bold, color.BgBlack).Println("❌ 项目 module name 初始化失败!!!")
	}
	_, _ = color.New(color.FgBlue).Printf("扫描到需要处理 %d 个文件\n", len(files))
	// 循环处理文件
	var errFilePathList []string

	for _, file := range files {
		if err := gfile.ReplaceFile("x-project", baseData.ProjectName, file); err != nil {
			errFilePathList = append(errFilePathList, file)
		}
	}
	_, _ = color.New(color.FgHiGreen).Printf("处理成功 %d 个文件\n", len(files)-len(errFilePathList))
	if len(errFilePathList) > 0 {
		_, _ = color.New(color.FgRed).Printf("处理失败 %d 个文件\n", len(errFilePathList))
		// 输出处理失败的文件目录
		for _, s := range errFilePathList {
			_, _ = color.New(color.FgRed).Printf("处理失败文件路径 %s\n", s)
		}
	}
	_, _ = color.New(color.FgYellow).Println("处理 项目名称完成")
}
