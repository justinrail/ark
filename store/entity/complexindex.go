package entity

//ComplexIndex 复杂指标计算，类似VDS
type ComplexIndex struct {
	ComplexIndexID   int    `xorm:"not null pk autoincr INT(11)"`
	ComplexIndexName string `xorm:"VARCHAR(128)"`
	Category         string `xorm:"VARCHAR(128)"`
	Title            string `xorm:"VARCHAR(128)"`
	Label            string `xorm:"VARCHAR(128)"`
	ObjectTypeID     int    `xorm:"null INT(11)"`
	BusinessID       string `xorm:"VARCHAR(128)"`
	GlobalResourceID int    `xorm:"null INT(11)"`
	CalcCron         string `xorm:"VARCHAR(128)"`
	CalcType         int    `xorm:"not null INT(11)"`
	//后期处理Expression
	AfterCalc  string `xorm:"VARCHAR(3096)"`
	SaveCron   string `xorm:"VARCHAR(128)"`
	Expression string `xorm:"VARCHAR(3096)"`
	Unit       string `xorm:"VARCHAR(128)"`
	Remark     string `xorm:"VARCHAR(255)"`
}
