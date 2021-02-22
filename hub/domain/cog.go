package domain

//COG Change of gateway gateway dto object
type COG struct {
	ID        int
	Name      string
	Address   string
	Flag      int
	Timestamp int64

	CORs []COR
}

//Clone clone cog
func (cog *COG) Clone() *COG {
	pcog := *cog

	return &pcog
}
