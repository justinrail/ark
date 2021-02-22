package bus

import "ark/hub/domain"

const (
	//QueueSize 队列大小
	QueueSize = 1024
)

//COGBus bus channel for COG
var COGBus chan *domain.COG

//CORBus bus channel for COR
var CORBus chan *domain.COR

//COVBus bus channel for COV
var COVBus chan []*domain.COV

func init() {
	COGBus = make(chan *domain.COG, QueueSize)
	CORBus = make(chan *domain.COR, QueueSize)
	COVBus = make(chan []*domain.COV, QueueSize)
}
