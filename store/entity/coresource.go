package entity

//CoreSource 采集单元的领域对象
type CoreSource struct {
	CoreSourceID int `xorm:"not null pk autoincr INT(11)"`
	GatewayID    int `xorm:"not null INT(11)"`
	//UniqueID 低端采集设备的对应ID
	UniqueID   string `xorm:"not null VARCHAR(128)"`
	SourceName string `xorm:"VARCHAR(128)"`
	State      string `xorm:"VARCHAR(128)"`
}
