package service

import (
	"ark/web/vm"
	"strconv"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

//GetAllRuntime Get all runtime information
func GetAllRuntime() []vm.LiteItem {
	items := make([]vm.LiteItem, 0)

	items = readNetConnections(items)

	items = readNetIOCounters(items)

	items = readProcess(items)

	return items
}

func readProcess(items []vm.LiteItem) []vm.LiteItem {
	procs, err := process.Processes()
	check(err)

	for _, p := range procs {

		pName, err3 := p.Name()
		check(err3)

		proName := "[" + strconv.Itoa(int(p.Pid)) + "] " + pName

		exePath, err6 := p.Exe()
		check(err6)

		item := vm.LiteItem{Category: "process", ItemName: proName, ItemValue: exePath}
		items = append(items, item)

	}

	return items
}

func readNetIOCounters(items []vm.LiteItem) []vm.LiteItem {
	iocs, err := net.IOCounters(true)
	check(err)

	for _, ioc := range iocs {
		//conName := con.Laddr.String() + "--" + con.Raddr.String()
		sendData := strconv.FormatUint(ioc.BytesSent/1024/1024, 10) + " MB"
		item := vm.LiteItem{Category: "netio", ItemName: "[" + ioc.Name + "] Send Data", ItemValue: sendData}
		items = append(items, item)
		sendPacket := strconv.FormatUint(ioc.PacketsSent, 10)
		item = vm.LiteItem{Category: "netio", ItemName: "[" + ioc.Name + "] Send Packets", ItemValue: sendPacket}
		items = append(items, item)

		recvData := strconv.FormatUint(ioc.BytesSent/1024/1024, 10) + " MB"
		item = vm.LiteItem{Category: "netio", ItemName: "[" + ioc.Name + "] Recv Data", ItemValue: recvData}
		items = append(items, item)
		recvPacket := strconv.FormatUint(ioc.PacketsRecv, 10)
		item = vm.LiteItem{Category: "netio", ItemName: "[" + ioc.Name + "] Recv Packets", ItemValue: recvPacket}
		items = append(items, item)

	}

	return items
}

func readNetConnections(items []vm.LiteItem) []vm.LiteItem {
	cons, err := net.Connections("inet")
	check(err)

	for _, con := range cons {
		conName := con.Laddr.String() + "--" + con.Raddr.String()
		item := vm.LiteItem{Category: "connection", ItemName: conName, ItemValue: con.Status}
		items = append(items, item)

	}

	return items
}
