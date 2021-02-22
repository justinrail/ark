package domain

import (
	"ark/hub/enum"
	"container/list"

	"github.com/cornelk/hashmap"
)

const (
	//CorePacketLogQueueSize coreSource级别日志队列
	CorePacketLogQueueSize = 10
)

//CoreSource 采集单元的领域对象
type CoreSource struct {
	CoreSourceID int
	GatewayID    int
	//UniqueID 低端采集设备的对应ID
	UniqueID   string
	SourceName string
	State      int

	CorePoints *hashmap.HashMap

	PacketLogs *list.List
}

//NewCoreSource create domain gateway
func NewCoreSource() CoreSource {
	coreSource := CoreSource{}
	coreSource.CorePoints = &hashmap.HashMap{}
	coreSource.PacketLogs = list.New()
	return coreSource
}

//UpdateState update state
func (coresource *CoreSource) UpdateState(cor *COR) {
	switch cor.Flag {
	case enum.CoreSourceFlagConDown:
		coresource.State = enum.CoreSourceConStateOffline
	default:
		coresource.State = enum.CoreSourceConStateOnline
	}

	coresource.PacketLogs.PushBack(cor) // Enqueue

	if coresource.PacketLogs.Len() > CorePacketLogQueueSize {
		e := coresource.PacketLogs.Front()

		if e != nil {
			coresource.PacketLogs.Remove(e)
		}

	}
}
