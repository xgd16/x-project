package main

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
	"os"
	"x-project/build/src"
	"x-project/build/types"
	_ "x-project/packed"
)

var version = 1
var debug = false

var basePath = "x-project"

var continueList = garray.NewStrArrayFrom([]string{
	"bin",
	"log",
}, true)

var fileList = func() []string {
	var arr []string

	paths := gres.ScanDir(fmt.Sprintf("%s", basePath), "*")

	for _, f := range paths {
		name := f.FileInfo().Name()
		if continueList.Contains(name) {
			continue
		}
		arr = append(arr, name)
	}

	return arr
}()

func main() {
	// 获取输入参数
	var p string
	flag.StringVar(&p, "p", "./x-project", "输出路径")
	flag.BoolVar(&debug, "debug", false, "调试模式")
	flag.Parse()
	// 如果是调试模式直接删除输入目录
	if debug {
		_ = gfile.Remove(p)
	}
	// 异常退出时删除生成的目录
	defer func() {
		if err := recover(); err != nil {
			_ = gfile.Remove(p)
			panic(err)
		}
	}()
	// 判断目录是否已经存在
	if gfile.IsDir(p) {
		fmt.Println(fmt.Sprintf("当前目录已存在 %s 文件夹", p))
		os.Exit(0)
	} else {
		if err := gfile.Mkdir(p); err != nil {
			panic("初始化项目失败错误: " + err.Error())
		}
	}

	fmt.Println("输出位置: " + p)
	// 输出目录数据
	var err error
	for _, s := range fileList {
		fmt.Println(fmt.Sprintf("已导出: %s", s))
		err = gres.Export(fmt.Sprintf("%s/%s", basePath, s), p)
		if err != nil {
			_ = gfile.Remove(p)
			panic("初始化项目失败错误: " + err.Error())
		}
		_ = gfile.Move(fmt.Sprintf("%s/%s/%s", p, basePath, s), fmt.Sprintf("%s/%s", p, s))
		_ = gfile.Remove(fmt.Sprintf("%s/%s", p, basePath))
	}
	// 创建基础 信息数据
	baseData := &types.BaseData{
		Debug:        debug,
		ProjectName:  gstr.Replace(p, "./", ""),
		OutPath:      p,
		BasePath:     basePath,
		ContinueList: continueList,
	}
	// 调用处理项目内容
	if basePath != baseData.OutPath {
		src.ProjectReName(baseData)
	}
	// 自动初始化项目需求扩展
	src.InitProject(baseData)
}
