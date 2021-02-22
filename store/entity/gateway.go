package entity

//Gateway entity of gateway table
type Gateway struct {
	ID        int    `xorm:"not null pk autoincr INT(11)"`
	UUID      string `xorm:"not null VARCHAR(32)"`
	Name      string `xorm:"VARCHAR(128)"`
	Collector string `xorm:"VARCHAR(128)"`
	IP        string `xorm:"VARCHAR(64)"`
	State     string `xorm:"VARCHAR(64)"`
	Joined    bool   `xorm:"Bool"`
}
