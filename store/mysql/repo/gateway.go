package repo

import (
	"ark/store/entity"
	"ark/store/mysql"
	"ark/util/log"
)

//GetGatewaysByCollector 根据采集器加载对应的gateway
func GetGatewaysByCollector(collectorName string) []entity.Gateway {
	gateways := make([]entity.Gateway, 0)

	//err := mysql.Engine().SQL("SELECT * from gateway where collector = ?", "stub").Find(&gateways)
	err := mysql.Engine().Where("collector = ?", collectorName).Find(&gateways)
	checkErr(err)

	return gateways
}

//FindGatewayByUUID 查找gatway
func FindGatewayByUUID(uuid string) entity.Gateway {
	gateways := make([]entity.Gateway, 0)
	err := mysql.Engine().Where("UUID = ?", uuid).Find(&gateways)
	checkErr(err)
	if len(gateways) == 0 {
		return entity.Gateway{}
	}

	return gateways[0]
}

//UpdateGatewayJoined UpdateGatewayJoined
func UpdateGatewayJoined(gatewayID int, joined bool) bool {
	gateway := entity.Gateway{Joined: joined}
	_, err := mysql.Engine().Id(gatewayID).Cols("Joined").Update(&gateway)

	if err == nil {
		return true
	}

	return false
}

//CascadeDeleteGateway 删除Gateway的所有采集配置
func CascadeDeleteGateway(gatewayID int) {
	sql := "delete from gateway where id = ?"
	_, err := mysql.Engine().Exec(sql, gatewayID)

	if err != nil {
		log.Error(err)
		return
	}

	css := GetCoreSourceByGateway(gatewayID)

	for _, cs := range css {

		sql = "delete from CorePoint where CoreSourceId = ?"
		_, err2 := mysql.Engine().Exec(sql, cs.CoreSourceID)

		if err2 == nil {
			sql = "delete from CoreSource where CoreSourceId = ?"
			_, err3 := mysql.Engine().Exec(sql, cs.CoreSourceID)

			if err3 != nil {
				log.Error(err3)
			}
		}
	}
}

//AddGateway 根据配置增加Gateway到数据库
func AddGateway(uuid string, name string, collector string, ip string) *entity.Gateway {
	gw := new(entity.Gateway)
	gw.UUID = uuid
	gw.Name = name
	gw.Collector = collector
	gw.IP = ip

	_, err := mysql.Engine().Insert(gw)

	if err != nil {
		log.Error(err)
	}
	return gw
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
