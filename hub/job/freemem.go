package job

import "runtime/debug"

func freeSystemMemory() {
	debug.FreeOSMemory()
}
