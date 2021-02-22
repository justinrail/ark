package bolt

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COGStateUpdater 更新Gateway的实际状态，并进行修改和日志
type COGStateUpdater struct {
	flow.Component
	In chan []*domain.COG
}

//OnIn Request handler
func (updater *COGStateUpdater) OnIn(cogs []*domain.COG) {
	metrics.AppMetrics.HubTopoCOGStateUpdaterCounter.Inc(int64(len(cogs)))
	for _, cog := range cogs {
		if cog.Flag >= 0 {
			updater.handleCOGState(cog.Clone())
		}
	}
}

func (updater *COGStateUpdater) handleCOGState(cog *domain.COG) {
	g, ok := domain.Gateways.Get(cog.ID)

	gateway := g.(*domain.Gateway)

	if ok {
		gatewayConState := enum.GatewayConStateUnkown

		switch cog.Flag {
		case enum.GatewayFlagRegister:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagAlarm:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagFTPAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagGetConfigAck:
			//配置同步是否成功，无法直接判断，由具体collector判断更新
		case enum.GatewayFlagGetDataAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagGetInfoAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagGetStorageRuleAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagGetThresholdAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagHeartbeat:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagRebootAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagSendConfig:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagSetConfigAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagSetInfoUpdateIntervalAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagSetPointAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagSetStorageRuleAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagSetThresholdAck:
			gatewayConState = enum.GatewayConStateOnline
		case enum.GatewayFlagTimeCheckAck:
			gatewayConState = enum.GatewayConStateOnline
		}

		gateway.ConState = gatewayConState

		gateway.AppendPacketLogs(cog)

	}
}
