package types

import "github.com/gogf/gf/v2/container/garray"

type BaseData struct {
	Debug        bool
	ProjectName  string
	OutPath      string
	BasePath     string
	ContinueList *garray.StrArray
}
