package service

import (
	"ark/web/vm"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/process"
)

//GetAllApplication return all application info items
func GetAllApplication() []vm.LiteItem {
	items := make([]vm.LiteItem, 0)
	pro, err2 := process.NewProcess(int32(os.Getpid()))
	check(err2)

	item := vm.LiteItem{Category: "basic", ItemName: "GO Version", ItemValue: runtime.Version()}
	items = append(items, item)

	item = vm.LiteItem{Category: "basic", ItemName: "Enviroment GO ROOT", ItemValue: os.Getenv("GOPATH")}
	items = append(items, item)

	item = vm.LiteItem{Category: "basic", ItemName: "Execute GO ROOT", ItemValue: runtime.GOROOT()}
	items = append(items, item)

	exename, err1 := os.Executable()
	check(err1)
	item = vm.LiteItem{Category: "abasicpps", ItemName: "Execute file path", ItemValue: exename}
	items = append(items, item)

	cmdl, err7 := pro.Cmdline()
	check(err7)
	item = vm.LiteItem{Category: "basic", ItemName: "Command Line", ItemValue: cmdl}
	items = append(items, item)

	wd, err := os.Getwd()
	check(err)
	item = vm.LiteItem{Category: "basic", ItemName: "Current Directory", ItemValue: wd}
	items = append(items, item)

	uname, err4 := pro.Username()
	check(err4)
	item = vm.LiteItem{Category: "basic", ItemName: "User Name", ItemValue: uname}
	items = append(items, item)

	ctime, err3 := pro.CreateTime()
	check(err3)

	cts := time.Unix(0, ctime*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
	item = vm.LiteItem{Category: "basic", ItemName: "Startup Time", ItemValue: cts}
	items = append(items, item)

	item = vm.LiteItem{Category: "probe", ItemName: "PID", ItemValue: strconv.Itoa(int(os.Getpid()))}
	items = append(items, item)

	item = vm.LiteItem{Category: "probe", ItemName: "UID", ItemValue: strconv.Itoa(int(os.Getuid()))}
	items = append(items, item)

	item = vm.LiteItem{Category: "probe", ItemName: "GID", ItemValue: strconv.Itoa(int(os.Getegid()))}
	items = append(items, item)

	nic, err8 := pro.Nice()
	check(err8)
	item = vm.LiteItem{Category: "probe", ItemName: "Nice (priority)", ItemValue: strconv.Itoa(int(nic))}
	items = append(items, item)

	cper, err5 := pro.CPUPercent()
	check(err5)
	cperString := strconv.FormatFloat(cper, 'f', 2, 64) + "%"
	item = vm.LiteItem{Category: "probe", ItemName: "CPU Usage", ItemValue: cperString}
	items = append(items, item)

	mper, err6 := pro.MemoryPercent()
	check(err6)
	mperString := strconv.FormatFloat(float64(mper), 'f', 2, 64) + "%"
	item = vm.LiteItem{Category: "probe", ItemName: "Memory Usage", ItemValue: mperString}
	items = append(items, item)

	nh, err11 := pro.NumThreads()
	check(err11)
	item = vm.LiteItem{Category: "probe", ItemName: "Thread Count", ItemValue: strconv.Itoa(int(nh))}
	items = append(items, item)

	item = vm.LiteItem{Category: "probe", ItemName: "Number of Routines", ItemValue: strconv.Itoa(int(runtime.NumGoroutine()))}
	items = append(items, item)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	almem := strconv.Itoa(int(m.Alloc/1024/1024)) + " MB"
	item = vm.LiteItem{Category: "probe", ItemName: "Allocate Memory", ItemValue: almem}
	items = append(items, item)

	tlmem := strconv.Itoa(int(m.TotalAlloc/1024/1024)) + " MB"
	item = vm.LiteItem{Category: "probe", ItemName: "Total Allocate Memory", ItemValue: tlmem}
	items = append(items, item)

	flmem := strconv.Itoa(int(m.Frees/1024/1024)) + " MB"
	item = vm.LiteItem{Category: "probe", ItemName: "Freed Memory", ItemValue: flmem}
	items = append(items, item)

	halmem := strconv.Itoa(int(m.HeapAlloc/1024/1024)) + " MB"
	item = vm.LiteItem{Category: "probe", ItemName: "Heap Allocated Memory", ItemValue: halmem}
	items = append(items, item)

	hrlmem := strconv.Itoa(int(m.HeapReleased/1024/1024)) + " MB"
	item = vm.LiteItem{Category: "probe", ItemName: "Heap Released Memory", ItemValue: hrlmem}
	items = append(items, item)

	item = vm.LiteItem{Category: "probe", ItemName: "Heap Object Count", ItemValue: strconv.Itoa(int(m.HeapObjects))}
	items = append(items, item)

	item = vm.LiteItem{Category: "probe", ItemName: "Number of GC Times", ItemValue: strconv.Itoa(int(m.NumGC))}
	items = append(items, item)

	lstGC := time.Unix(0, int64(m.LastGC)*int64(time.Nanosecond)).Format("2006-01-02 15:04:05")
	item = vm.LiteItem{Category: "probe", ItemName: "Last GC Time", ItemValue: lstGC}
	items = append(items, item)

	iocs, err9 := pro.IOCounters()
	check(err9)
	item = vm.LiteItem{Category: "probe", ItemName: "IO Counter", ItemValue: iocs.String()}
	items = append(items, item)

	// nf, err10 := pro.NumFDs()
	// check(err10)
	// item = vm.LiteItem{Category: "apps", ItemName: "File Descriptor Count", ItemValue: strconv.Itoa(int(nf))}
	// items = append(items, item)

	// ofs, err12 := pro.OpenFiles()
	// check(err12)

	// for i, of := range ofs {
	// 	item := vm.LiteItem{Category: "apps", ItemName: "Opened File " + strconv.Itoa(i), ItemValue: of.String()}
	// 	items = append(items, item)
	// }

	// item = vm.LiteItem{Category: "apps", ItemName: "Environment Variables", ItemValue: strings.Join(os.Environ()[:], ",")}
	// items = append(items, item)

	// items = readHostInfo(items)

	return items
}
