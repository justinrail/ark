package domain

//COR CoreSource
type COR struct {
	GateWayID    int
	CoreSourceID int
	UniqueID     string
	Name         string
	Address      string
	Flag         int
	Timestamp    int64

	COVs []COV
}

//Clone clone cor
func (cor *COR) Clone() *COR {
	pcor := *cor

	return &pcor
}
