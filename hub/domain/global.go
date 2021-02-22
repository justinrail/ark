package domain

import (
	"ark/util/slide"

	"github.com/cornelk/hashmap"
)

//Gateways global gateways
var Gateways *hashmap.HashMap

func init() {
	Gateways = &hashmap.HashMap{}
	CoreLiveEvents = &hashmap.HashMap{}
	CorePoints = &hashmap.HashMap{}
	CoreSources = &hashmap.HashMap{}
	ComplexIndexs = &hashmap.HashMap{}
	MessageSinks = make(map[string]*slide.LengthWindow)
}

//CoreLiveEvents  global gateways
var CoreLiveEvents *hashmap.HashMap

//CorePoints global points
var CorePoints *hashmap.HashMap

//CoreSources global coresources
var CoreSources *hashmap.HashMap

//ComplexIndexs global complex index cache
var ComplexIndexs *hashmap.HashMap

//MessageSinks 通知发送消息的有限长度缓冲hash
var MessageSinks map[string]*slide.LengthWindow
