package bolt

import (
	"ark/hub/domain"

	flow "github.com/trustmaster/goflow"
)

//CORAshbin 垃圾桶，用来进行metrics统计
type CORAshbin struct {
	flow.Component
	In chan []*domain.COR
}

//OnIn Request handler
func (ashbin *CORAshbin) OnIn(cors []*domain.COR) {
	//fmt.Printf("ashbin cog:%p val=%v\n", ashbin, cogs)
	//fmt.Print(cog)
}
