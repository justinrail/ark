package processor

import (
	"ark/hub/bus"
	"ark/hub/collector/cmb/dto"
	"ark/hub/collector/cmb/ws"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"
	"time"

	flow "github.com/trustmaster/goflow"
)

//LoginProcessor LoginProcessor
type LoginProcessor struct {
	flow.Component
	In chan *ws.Conversation
}

//OnIn Request handler
func (loginProcessor *LoginProcessor) OnIn(cvs *ws.Conversation) {
	metrics.AppMetrics.HubCollectorCMBLoginProcessorCounter.Inc(1)

	req := &dto.LoginRequest{}
	cvs.GetRequest(req)

	//从现有的Gateways中查找这个gateway的UUID==FSUID，如果存在，则返回成功，生成COG事件入BUS，事件状态为Registerd，流程结束
	//如果不存在，增加Gateway状态为配置同步中，等FSU发送配置数据，等收到配置数据写库和更新FSU信息后，修改状态为时钟同步中
	//如果时钟同步完毕，则Gateway更改为正常模式，如果收到心跳则标志心跳COG
	//上述工作在Gateway Life Supervisor完成，生命周期管理这个功能负责维护数据的一致性

	//本处负责异常处理，不允许接入的，这里判断，不合格的这里判断。已屏蔽的这里判断。如果是正常的
	//在目前的例子中，暂不考虑异常，所以默认返回成功即可，返回成功后，配置同步才可能进行（配置同步后，才存gateway）

	loginRep := &dto.LoginResponse{}
	loginRep.FSUID = req.FSUID
	loginRep.Result = 1
	loginRep.FailureCause = "NULL"

	cvs.SendResponse("LOGIN_ACK", loginRep)

	//login的意义在于，让fsu确认中心存在并且可以处理业务了。
	//如果之前不存在gateway，则直接成功，等配置上送
	//如果之前存在gateway，则把register数据给后台
	gs := domain.Gateways.Iter()
	for kv := range gs {
		g := kv.Value.(*domain.Gateway)
		if g.UUID == req.FSUID {

			cog := &domain.COG{}
			cog.ID = g.ID
			cog.Flag = enum.GatewayFlagRegister
			cog.Timestamp = time.Now().Unix()
			cog.Name = g.Name
			cog.Address = cvs.GetClientIP()

			metrics.AppMetrics.HubCollectorCMBCOGCounter.Inc(1)
			bus.COGBus <- cog
			break
		}

	}
}
