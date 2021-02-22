package service

import (
	"ark/util/log"
	"ark/util/str"
	"ark/web/vm"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

//GetAllMachine Get all machine information
func GetAllMachine() []vm.LiteItem {
	items := make([]vm.LiteItem, 0)

	items = readDiskInfo(items)

	items = readHostInfo(items)

	items = readCPUInfo(items)

	items = readMemoryInfo(items)

	items = readNetInterface(items)

	return items
}

func readNetInterface(items []vm.LiteItem) []vm.LiteItem {
	nis, err := net.Interfaces()
	check(err)

	for _, ni := range nis {

		//perString := strconv.FormatFloat(per, 'f', 2, 64) + "%"
		item := vm.LiteItem{Category: "interface", ItemName: "[" + ni.Name + "] MAC", ItemValue: ni.HardwareAddr}
		items = append(items, item)

		item = vm.LiteItem{Category: "interface", ItemName: "[" + ni.Name + "] Flags", ItemValue: strings.Join(ni.Flags[:], ",")}
		items = append(items, item)

		for i, ad := range ni.Addrs {
			item = vm.LiteItem{Category: "interface", ItemName: "[" + ni.Name + "] Addr " + strconv.Itoa(i), ItemValue: ad.String()}
			items = append(items, item)
		}
	}

	return items
}

func readMemoryInfo(items []vm.LiteItem) []vm.LiteItem {
	v, err := mem.VirtualMemory()
	check(err)

	totalMemory := strconv.FormatUint(v.Total/1024/1024/1024, 10) + " GiB"
	item := vm.LiteItem{Category: "mem", ItemName: "Total Memory", ItemValue: totalMemory}
	items = append(items, item)

	// freeMemory := strconv.FormatUint(v.Free/1024/1024/1024, 10) + " GiB"
	// item = vm.LiteItem{Category: "host", ItemName: "Free Memory", ItemValue: freeMemory}
	// items = append(items, item)

	perString := strconv.FormatFloat(v.UsedPercent, 'f', 2, 64) + "%"
	item = vm.LiteItem{Category: "mem", ItemName: "Memory Usage", ItemValue: perString}
	items = append(items, item)

	return items
}

func readCPUInfo(items []vm.LiteItem) []vm.LiteItem {
	cpus, err := cpu.Info()
	check(err)

	c, err1 := cpu.Counts(true)
	check(err1)

	cpuCount := strconv.Itoa(int(c))
	item := vm.LiteItem{Category: "cpu", ItemName: "Logic CPU Count", ItemValue: cpuCount}
	items = append(items, item)

	pers, err2 := cpu.Percent(0, true)
	check(err2)

	for i, per := range pers {

		perString := strconv.FormatFloat(per, 'f', 2, 64) + "%"
		item := vm.LiteItem{Category: "cpu", ItemName: "Logic CPU " + strconv.Itoa(i) + " Usage ", ItemValue: perString}
		items = append(items, item)
	}

	for _, info := range cpus {
		items = updateCPUInfo(info, items)
	}

	return items
}

func updateCPUInfo(u cpu.InfoStat, items []vm.LiteItem) []vm.LiteItem {

	cpuIndex := strconv.Itoa(int(u.CPU))
	// item := vm.LiteItem{Category: "host", ItemName: "CPU [" + cpuIndex + "] ", ItemValue: cpuIndex}
	// items = append(items, item)

	item := vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] VendorID ", ItemValue: u.VendorID}
	items = append(items, item)

	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] Family ", ItemValue: u.Family}
	items = append(items, item)

	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] Model ", ItemValue: u.Model}
	items = append(items, item)

	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] PhysicalID ", ItemValue: u.PhysicalID}
	items = append(items, item)

	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] CoreID ", ItemValue: u.CoreID}
	items = append(items, item)

	cpuCores := strconv.Itoa(int(u.Cores))
	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] Cores ", ItemValue: cpuCores}
	items = append(items, item)

	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] ModelName ", ItemValue: u.ModelName}
	items = append(items, item)

	mhzString := strconv.FormatFloat(u.Mhz, 'f', 2, 64)
	item = vm.LiteItem{Category: "cpu", ItemName: "CPU [" + cpuIndex + "] mhz ", ItemValue: mhzString}
	items = append(items, item)

	// partitionFree := strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB"
	// item = vm.LiteItem{Category: "disk", ItemName: "Partition Free [" + u.Path + "]", ItemValue: partitionFree}
	// items = append(items, item)

	return items
}

func readHostInfo(items []vm.LiteItem) []vm.LiteItem {
	info, err := host.Info()
	check(err)
	item := vm.LiteItem{Category: "host", ItemName: "Host ID", ItemValue: info.HostID}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "Host Name", ItemValue: info.Hostname}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "OS", ItemValue: info.OS}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "Platform", ItemValue: info.Platform}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "PlatformFamily", ItemValue: info.PlatformFamily}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "PlatformVersion", ItemValue: info.PlatformVersion}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "KernelVersion", ItemValue: info.KernelVersion}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "VirtualizationSystem", ItemValue: info.VirtualizationSystem}
	items = append(items, item)

	item = vm.LiteItem{Category: "host", ItemName: "VirtualizationRole", ItemValue: info.VirtualizationRole}
	items = append(items, item)

	procCount := strconv.FormatUint(info.Procs, 10)
	item = vm.LiteItem{Category: "host", ItemName: "Process Count", ItemValue: procCount}
	items = append(items, item)

	upTimeString := str.TimeToString(int64(info.Uptime))
	item = vm.LiteItem{Category: "host", ItemName: "Uptime", ItemValue: upTimeString}
	items = append(items, item)

	bootBootTimeString := str.TimeToString(int64(info.BootTime))
	item = vm.LiteItem{Category: "host", ItemName: "BootTime", ItemValue: bootBootTimeString}
	items = append(items, item)

	return items
}

func readDiskInfo(items []vm.LiteItem) []vm.LiteItem {
	parts, err := disk.Partitions(false)
	check(err)

	var usage []*disk.UsageStat

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		if err == nil {
			usage = append(usage, u)
			items = updateDiskUsage(u, items)
		}
	}

	return items
}

func updateDiskUsage(u *disk.UsageStat, items []vm.LiteItem) []vm.LiteItem {

	partitionUsage := strconv.FormatFloat(u.UsedPercent, 'f', 2, 64) + "% full."
	item := vm.LiteItem{Category: "disk", ItemName: "Partition Usage [" + u.Path + "]", ItemValue: partitionUsage}
	items = append(items, item)

	partitionTotal := strconv.FormatUint(u.Total/1024/1024/1024, 10) + " GiB"
	item = vm.LiteItem{Category: "disk", ItemName: "Partition Total [" + u.Path + "]", ItemValue: partitionTotal}
	items = append(items, item)

	partitionFree := strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB"
	item = vm.LiteItem{Category: "disk", ItemName: "Partition Free [" + u.Path + "]", ItemValue: partitionFree}
	items = append(items, item)

	return items
}

func check(err error) {
	if err != nil {
		log.Error(err)
	}
}
