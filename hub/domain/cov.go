package domain

//COV change of corepoint
type COV struct {
	GateWayID    int
	CoreSourceID int
	CorePointID  int
	StateID      int
	IsValid      bool
	//当前数值值
	CurrentNumericValue float32
	//当前字符串值
	CurrentStringValue string

	Timestamp int64
}

//Clone clone cov
func (cov *COV) Clone() *COV {
	pcov := &COV{}
	pcov.GateWayID = cov.GateWayID
	pcov.CoreSourceID = cov.CoreSourceID
	pcov.CurrentNumericValue = cov.CurrentNumericValue
	pcov.CorePointID = cov.CorePointID
	pcov.CurrentStringValue = cov.CurrentStringValue
	pcov.IsValid = cov.IsValid
	pcov.StateID = cov.StateID
	pcov.Timestamp = cov.Timestamp

	return pcov
}
