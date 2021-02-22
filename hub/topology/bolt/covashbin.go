package bolt

import (
	"ark/hub/domain"

	flow "github.com/trustmaster/goflow"
)

//COVAshbin 垃圾桶，用来进行metrics统计
type COVAshbin struct {
	flow.Component
	In chan []*domain.COV
}

//OnIn Request handler
func (ashbin *COVAshbin) OnIn(covs []*domain.COV) {

}
