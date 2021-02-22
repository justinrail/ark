package repo

import (
	"ark/store/entity"
	"ark/store/mysql"
	"ark/util/log"
)

//GetCorePointByByCoreSource 根据coresource加载对应的corepoint
func GetCorePointByByCoreSource(coreSourceID int) []entity.CorePoint {
	corePoints := make([]entity.CorePoint, 0)

	err := mysql.Engine().Where("CoreSourceId = ?", coreSourceID).Find(&corePoints)
	checkErr(err)

	return corePoints
}

//GetCorePointHasCron 获取有存储设置的测点
func GetCorePointHasCron() []entity.CorePoint {
	corePoints := make([]entity.CorePoint, 0)

	err := mysql.Engine().Where(" LENGTH(Cron)>1 ").Find(&corePoints)
	checkErr(err)

	return corePoints
}

//AddCorePoint 根据配置增加CorePoint到数据库
func AddCorePoint(coreSourceID int, name string, uniqueID string, dataType int, eventSeverity int, standardID string, readable bool, writable bool, stateRule int) int {
	cp := new(entity.CorePoint)
	cp.CoreSourceID = coreSourceID
	cp.CoreDataTypeID = dataType //1:整数，2：浮点数
	cp.Accuracy = "00.00"
	cp.EventSeverity = eventSeverity
	cp.PointName = name
	cp.Readable = readable
	cp.Writable = writable
	cp.StateRuleID = stateRule
	cp.OriginStandardID = standardID
	cp.UniqueID = uniqueID
	cp.Step = 0.1

	_, err := mysql.Engine().Insert(cp)

	if err != nil {
		log.Error(err)
	}
	return cp.CorePointID
}
