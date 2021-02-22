package dto

import "github.com/cornelk/hashmap"

//CMBGateways global gateways for cmb collector
var CMBGateways *hashmap.HashMap

//DeviceCoreSourceMap CMB设备和coresource的对应关系
var DeviceCoreSourceMap map[string]int

//SignalCorePointMap 信号和CorePoint的对应关系
var SignalCorePointMap map[string]int

func init() {
	CMBGateways = &hashmap.HashMap{}
	DeviceCoreSourceMap = make(map[string]int)
	SignalCorePointMap = make(map[string]int)
}
