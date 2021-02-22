package bolt

import (
	"ark/hub/domain"

	flow "github.com/trustmaster/goflow"
)

//COGAshbin 垃圾桶，用来进行metrics统计
type COGAshbin struct {
	flow.Component
	In chan []*domain.COG
}

//OnIn Request handler
func (ashbin *COGAshbin) OnIn(cogs []*domain.COG) {
	//fmt.Printf("ashbin cog:%p val=%v\n", ashbin, cogs)
	//fmt.Print(cog)
}
