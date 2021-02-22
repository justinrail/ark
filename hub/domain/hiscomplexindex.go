package domain

//HisComplexIndex 复杂指标历史记录
type HisComplexIndex struct {
	ComplexIndexID   int
	ComplexIndexName string
	Category         string
	Title            string
	Label            string
	BusinessID       string
	ObjectTypeID     int
	GlobalResourceID int
	CalcType         int
	Unit             string
	//上次的值
	LastValue float64
	//上次采集时间
	LastTimestamp int64
	//当前数据采集时间
	Timestamp int64
	//当前值
	CurrentValue float64
}
