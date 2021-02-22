package sysevent

import (
	evbus "github.com/asaskevich/EventBus"
)

//EventBus EventBus
var EventBus evbus.Bus

func init() {
	EventBus = evbus.New()
}
