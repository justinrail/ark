package vm

//ComplexIndexView complex index view model
type ComplexIndexView struct {
	ComplexIndexID   int
	ComplexIndexName string
	Category         string
	Title            string
	Label            string
	ObjectTypeID     int
	BusinessID       string
	GlobalResourceID int
	CalcCron         string
	CalcType         int
	//后期处理Expression
	AfterCalc  string
	SaveCron   string
	Expression string
	Unit       string
	Remark     string
	//上次的值
	LastValueString string
	//上次采集时间
	LastTimestampString string
	//当前数据采集时间
	TimestampString string
	//当前值
	CurrentValueString string

	//当前值是否有效
	IsValid bool
}
