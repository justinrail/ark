package dto

//CMBGateway CMBGateway
type CMBGateway struct {
	ID        string
	GatewayID int
	UserName  string
	IP        string
	Devices   []Device
}
