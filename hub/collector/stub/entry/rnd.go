package entry

import (
	"ark/hub/collector/stub/dto"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/store/mysql/repo"
	"ark/util/exp"
	"math/rand"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

var covCache []*domain.COV

//RandomCOG generate random cog
func RandomCOG() *dto.Packet {
	gateways := repo.GetGatewaysByCollector("stub")

	x := rand.Intn(len(gateways))

	//y := rand.Intn(7)
	y := rand.Intn(enum.GatewayFlagSetInfoUpdateIntervalAck-enum.GatewayFlagUnkown+1) + enum.GatewayFlagUnkown
	//flagString := enum.GetEnumString(y)
	cog := domain.COG{ID: gateways[x].ID, Name: gateways[x].Name, Address: "127.0.0.1", Flag: y, Timestamp: time.Now().Unix()}
	return &dto.Packet{MessageType: enum.PacketCOG, Body: cog}
}

//randomCOR generate random cor
func randomCOR() domain.COR {
	cors := repo.GetCoreSourceByGateway(-1)

	x := rand.Intn(len(cors))
	u4, err := uuid.NewV4()
	exp.CheckError(err)

	//y := rand.Intn(5) + 7
	//flagString := enum.GetEnumString(y)
	seed := rand.Intn(enum.CoreSourceFlagConDown - enum.CoreSourceFlagUnkown + 1)
	y := enum.CoreSourceFlagUnkown + seed

	//fmt.Println(enum.GetEnumString(y))

	cor := domain.COR{CoreSourceID: cors[x].CoreSourceID, GateWayID: -1, UniqueID: u4.String(), Name: cors[x].SourceName,
		Flag: y, Timestamp: time.Now().Unix()}

	return cor

}

//RandomCORPacket random cor packet
func RandomCORPacket() *dto.Packet {
	return &dto.Packet{MessageType: enum.PacketCOR, Body: randomCOR()}
}

//randomCOV generate random cov
func randomCOV() domain.COV {
	cors := repo.GetCoreSourceByGateway(-1)

	x := rand.Intn(len(cors))

	gatewayID := -1
	corSourceID := cors[x].CoreSourceID

	y := rand.Intn(enum.CorePointFlagEnd-enum.CorePointFlagUnkown+1) + enum.CorePointFlagUnkown
	flagString := enum.GetEnumString(y)

	cops := repo.GetCorePointByByCoreSource(corSourceID)

	z := rand.Intn(len(cops))

	corePointID := cops[z].CorePointID

	v := rand.Float32() * 10
	cov := domain.COV{GateWayID: gatewayID, CoreSourceID: corSourceID, CorePointID: corePointID, StateID: y,
		CurrentStringValue: flagString, Timestamp: time.Now().Unix(), CurrentNumericValue: v, IsValid: true}
	return cov

}

//RandomCOVPacketPerfermance random cov packet
func RandomCOVPacketPerfermance() *dto.Packet {

	if covCache == nil {
		covs := make([]*domain.COV, 100)

		for i := range covs {
			cov := randomCOV()
			covs[i] = &cov
		}

		covCache = covs
	}

	batch := 1000
	res := make([]*domain.COV, 0)
	for a := 0; a < batch; a++ {
		y := rand.Intn(enum.CorePointFlagEnd-enum.CorePointFlagUnkown+1) + enum.CorePointFlagUnkown
		timeNow := time.Now().Unix()
		v := rand.Float32() * 10

		for _, cov := range covCache {
			cov.CurrentNumericValue = v
			cov.StateID = y
			cov.Timestamp = timeNow
		}

		res = append(res, covCache...)
	}

	return &dto.Packet{MessageType: enum.PacketCOV, Body: res}
}

//RandomCOVPacket 随机COV数据包
func RandomCOVPacket() *dto.Packet {
	covs := make([]*domain.COV, 10)

	for i := range covs {
		cov := randomCOV()
		covs[i] = &cov
	}

	return &dto.Packet{MessageType: enum.PacketCOV, Body: covs}
}
