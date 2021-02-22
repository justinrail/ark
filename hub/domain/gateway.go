package domain

import (
	"ark/hub/enum"
	"ark/store/mysql/repo"
	"ark/util/queue"
	"container/list"

	"github.com/cornelk/hashmap"
)

const (
	//GateWayRequestQueueSize gateway级别控制队列
	GateWayRequestQueueSize = 4
	//GateWayLogQueueSize gateway级别日志队列
	GateWayLogQueueSize = 10
)

//Gateway Gateway Domain Object
type Gateway struct {
	ID          int
	UUID        string
	Name        string
	Collector   string
	IP          string
	ConState    int
	SynState    int
	DebugState  bool
	Joined      bool
	CoreSources *hashmap.HashMap

	//RequestQueue Gateway，CoreSource，CorePoint的所有下发请求队列
	RequestQueue *queue.EsQueue
	//Gateway下发队列日志
	RequestLogs *queue.EsQueue
	//ResponseLogs Gateway反馈队列日志，反馈队列由chan进行收集处理，这里只记录最终日志
	ResponseLogs *queue.EsQueue
	//PacketLogs COG历史日志（什么时候Register，什么时候offline等切换）
	PacketLogs *list.List
}

//CascadeDeleteGateway 级联删除gateway
func CascadeDeleteGateway(gatewayID int) {
	repo.CascadeDeleteGateway(gatewayID)
}

//NewGateway create domain gateway
func NewGateway() Gateway {
	gateway := Gateway{}
	gateway.ConState = enum.GatewayConStateUnkown
	gateway.SynState = enum.GatewaySynStateUnkown
	gateway.DebugState = false
	gateway.CoreSources = &hashmap.HashMap{}
	gateway.RequestQueue = queue.NewQueue(GateWayRequestQueueSize)
	gateway.RequestLogs = queue.NewQueue(GateWayLogQueueSize)
	gateway.ResponseLogs = queue.NewQueue(GateWayLogQueueSize)
	gateway.PacketLogs = list.New()
	return gateway
}

//AppendPacketLogs update state
func (gateway *Gateway) AppendPacketLogs(cog *COG) {
	gateway.PacketLogs.PushBack(cog) // Enqueue

	if gateway.PacketLogs.Len() > GateWayLogQueueSize {
		e := gateway.PacketLogs.Front()

		if e != nil {
			gateway.PacketLogs.Remove(e)
		}
	}
}
