package govern

import "github.com/cornelk/hashmap"

const (
	//FuseGatewaylifeSupervisor switch for gateway life supervisor processor
	FuseGatewaylifeSupervisor = iota
	//FuseCoreSourceLifeSupervisor switch for coresource life supervisor processor
	FuseCoreSourceLifeSupervisor
)

//Fuses all fuses
var Fuses hashmap.HashMap

func init() {
	Fuses = hashmap.HashMap{}
	Fuses.Insert(FuseGatewaylifeSupervisor, true)
	Fuses.Insert(FuseCoreSourceLifeSupervisor, true)
}

// //ProcessorFuse 流处理器熔断器
// type ProcessorFuse struct {
// }

//CheckFuse 检测熔断器状态，如果Good/Broken(true/false)=Broken则吃掉所有的消息
func CheckFuse(key int) bool {
	fuseStatus, ok := Fuses.Get(key)
	if ok {
		return fuseStatus.(bool)
	}

	return true
}
