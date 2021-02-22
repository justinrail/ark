package task

import (
	"ark/hub/collector/cmb/dto"
	"ark/hub/collector/cmb/ws"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"
	"encoding/xml"
	"fmt"

	"github.com/robfig/cron"
)

var cr *cron.Cron

//Ready prepare for jobs
func Ready() {
	cr = cron.New()
	cr.AddFunc("@every 5s", GetFSUData)

}

//Run start task
func Run() {
	cr.Run()
}

//GetFSUData 遍历cmbgateway，调用getData命令//获取实时数据
func GetFSUData() {

	cgs := dto.CMBGateways.Iter()

	for kv := range cgs {
		gateway := kv.Value.(*dto.CMBGateway)
		dg, exist := domain.Gateways.Get(gateway.GatewayID)
		if exist {
			if (dg.(*domain.Gateway)).ConState != enum.GatewayConStateOnline {
				return
			}
		}

		req := &dto.GetDataRequest{}
		req.FSUID = gateway.ID
		req.DeviceList = dto.DeviceList{}
		req.DeviceList.Devices = make([]dto.GetDataDevice, 0)
		req.DeviceList.Devices = append(req.DeviceList.Devices, dto.GetDataDevice{})
		returnXML := ws.SOAPInvoke(gateway.IP, "GET_DATA", req)

		resp := &dto.Response{}

		err := xml.Unmarshal([]byte(returnXML), resp)

		if err == nil {
			//counter
			metrics.AppMetrics.HubCollectorCMBGetDataTaskCounter.Inc(1)
			if resp.PKType.Name == "GET_DATA_ACK" {
				xmlContent := fmt.Sprintf("<Info>%s</Info>", resp.Info.InnerText)
				gdresp := &dto.GetDataResponse{}
				xml.Unmarshal([]byte(xmlContent), gdresp)
				//fmt.Println(xmlContent)
				handleGetDataResponse(gdresp)
			}
		}
	}

}
