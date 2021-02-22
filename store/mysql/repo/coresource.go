package repo

import (
	"ark/store/entity"
	"ark/store/mysql"
	"ark/util/log"
)

//GetCoreSourceByGateway 根据gateway加载对应的coresource
func GetCoreSourceByGateway(gatewayID int) []entity.CoreSource {
	coreSources := make([]entity.CoreSource, 0)

	//err := mysql.Engine().SQL("SELECT * from gateway where collector = ?", "stub").Find(&gateways)
	err := mysql.Engine().Where("GatewayId = ?", gatewayID).Find(&coreSources)
	checkErr(err)

	return coreSources
}

//GetAllCoreSource 获取所有gateway
func GetAllCoreSource() []entity.CoreSource {
	coreSources := make([]entity.CoreSource, 0)

	err := mysql.Engine().Find(&coreSources)
	checkErr(err)

	return coreSources
}

//AddCoreSource 根据配置增加CoreSource到数据库
func AddCoreSource(gatewayID int, name string, uniqueID string) int {
	cs := new(entity.CoreSource)
	cs.GatewayID = gatewayID
	cs.SourceName = name
	cs.UniqueID = uniqueID

	_, err := mysql.Engine().Insert(cs)

	if err != nil {
		log.Error(err)
	}
	return cs.CoreSourceID
}
