package vm

//CorePointView corepoint view model
type CorePointView struct {
	CorePointID    int    `xorm:"not null pk autoincr INT(11)"`
	PointName      string `xorm:"VARCHAR(128)"`
	Accuracy       string `xorm:"VARCHAR(32)"`
	Unit           string `xorm:"VARCHAR(32)"`
	Max            string `xorm:"VARCHAR(32)"`
	Min            string `xorm:"VARCHAR(32)"`
	CoreSourceID   int    `xorm:"not null INT(11)"`
	CoreDataTypeID int    `xorm:"not null INT(11)"`
	EventSeverity  int    `xorm:"not null INT(11)"`
	StateRuleID    int    `xorm:"not null INT(11)"`
	Readable       bool   `xorm:"Bool"`
	Writable       bool   `xorm:"Bool"`
	Masked         bool   `xorm:"Bool"`
	DefaultValue   string `xorm:"VARCHAR(32)"`
	Step           string `xorm:"Float"`
	StandardID     string `xorm:"VARCHAR(20)"`
	Cron           string `xorm:"VARCHAR(128)"`
	Expression     string `xorm:"VARCHAR(255)"`

	GatewayID int
	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	//事件：开始，确认，结束
	CurrentEventState int
	//数据是否有效
	IsAvailabe bool
	//数据是否有效,0:ok,1:low limit exceed 2: up limit exceed
	LimitState int
	//更新时间
	UpdateTimeString string
	//常用的State开始对应的，开始时间，如果CurrentEventState可用，否则存的是上次的结束时间
	StartTime int64
	//常用的State结束对应的，结束时间，如果CurrentEndState可用，否则存的是上次的结束时间
	EndTime int64
}
