package entity

//CorePoint 测点
type CorePoint struct {
	CorePointID      int     `xorm:"not null pk autoincr INT(11)"`
	PointName        string  `xorm:"VARCHAR(128)"`
	Accuracy         string  `xorm:"VARCHAR(32)"`
	Unit             string  `xorm:"VARCHAR(32)"`
	Max              string  `xorm:"VARCHAR(32)"`
	Min              string  `xorm:"VARCHAR(32)"`
	CoreSourceID     int     `xorm:"not null INT(11)"`
	CoreDataTypeID   int     `xorm:"not null INT(11)"`
	EventSeverity    int     `xorm:"not null INT(11)"`
	StateRuleID      int     `xorm:"not null INT(11)"`
	Readable         bool    `xorm:"Bool"`
	Writable         bool    `xorm:"Bool"`
	Masked           bool    `xorm:"Bool"`
	DefaultValue     string  `xorm:"VARCHAR(32)"`
	Step             float32 `xorm:"Float"`
	OriginStandardID string  `xorm:"VARCHAR(20)"`
	StandardID       int     `xorm:"not null INT(11)"`
	Cron             string  `xorm:"VARCHAR(128)"`
	Expression       string  `xorm:"VARCHAR(255)"`
	UniqueID         string  `xorm:"VARCHAR(128)"`
}
