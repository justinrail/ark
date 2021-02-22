package service

import (
	"ark/util/cfg"
	"ark/web/vm"
	"strconv"
)

//GetAllConfig return all config items
func GetAllConfig() []vm.LiteItem {
	items := make([]vm.LiteItem, 0)

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "WebRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().App.WebRunFuse),
		Remark:    "是否启动自管理网站",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "ProfileServerRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().App.ProfileServerRunFuse),
		Remark:    "是否启动性能诊断服务",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "WebGinDebugMode",
		ItemValue: strconv.FormatBool(cfg.Read().App.WebGinDebugMode),
		Remark:    "是否开启自管理网站调试模式",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "WebServerPort",
		ItemValue: strconv.Itoa(cfg.Read().App.WebServerPort),
		Remark:    "自管理网站的对外HTTP端口",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MySQLServerIP",
		ItemValue: cfg.Read().App.MySQLServerIP,
		Remark:    "MYSQL的服务器IP",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MySQLServerPort",
		ItemValue: cfg.Read().App.MySQLServerPort,
		Remark:    "MYSQL的服务器端口",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MySQLServerUserName",
		ItemValue: cfg.Read().App.MySQLServerUserName,
		Remark:    "MYSQL的服务器用户名",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MySQLServerPassword",
		ItemValue: cfg.Read().App.MySQLServerPassword,
		Remark:    "MYSQL的服务器密码",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MySQLServerDBName",
		ItemValue: cfg.Read().App.MySQLServerDBName,
		Remark:    "MYSQL的服务器默认数据库",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "InfluxDBServerIP",
		ItemValue: cfg.Read().App.InfluxDBServerIP,
		Remark:    "InfluxDB的服务器IP",
	})

	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "InfluxDBServerPort",
		ItemValue: cfg.Read().App.InfluxDBServerPort,
		Remark:    "InfluxDB的服务器Port",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "InfluxDBServerUserName",
		ItemValue: cfg.Read().App.InfluxDBServerUserName,
		Remark:    "InfluxDB的服务器用户名",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "InfluxDBServerPassword",
		ItemValue: cfg.Read().App.InfluxDBServerPassword,
		Remark:    "InfluxDB的服务器密码",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "InfluxDBServerDBName",
		ItemValue: cfg.Read().App.InfluxDBServerDBName,
		Remark:    "InfluxDB的服务器默认数据库",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MongoDBServerIP",
		ItemValue: cfg.Read().App.MongoDBServerIP,
		Remark:    "MongoDB的服务器IP",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MongoDBServerPort",
		ItemValue: cfg.Read().App.MongoDBServerPort,
		Remark:    "MongoDB的服务器Port",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MongoDBServerUserName",
		ItemValue: cfg.Read().App.MongoDBServerUserName,
		Remark:    "MongoDB的服务器用户名",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MongoDBServerPassword",
		ItemValue: cfg.Read().App.MongoDBServerPassword,
		Remark:    "MongoDB的服务器密码",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "MongoDBServerDBName",
		ItemValue: cfg.Read().App.MongoDBServerDBName,
		Remark:    "MongoDB的服务器默认数据库",
	})
	items = append(items, vm.LiteItem{
		Category:  "app",
		ItemName:  "JSONDBPath",
		ItemValue: cfg.Read().App.JSONDBPath,
		Remark:    "JSON本地存储主目录",
	})

	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "CollectorCMBRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.CollectorCMBRunFuse),
		Remark:    "是否启动中国移动B接口采集器",
	})
	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "CollectorStubRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.CollectorStubRunFuse),
		Remark:    "是否启动内置B接口采集器模拟器",
	})
	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "JobHisPointRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.JobHisPointRunFuse),
		Remark:    "是否保存历史CorePoint记录",
	})
	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "JobHisComplexIndexRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.JobHisComplexIndexRunFuse),
		Remark:    "是否保存历史指标",
	})
	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "JobComplexIndexRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.JobComplexIndexRunFuse),
		Remark:    "是否启动指标计算",
	})
	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "NotifyRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.NotifyRunFuse),
		Remark:    "是否启动事件通知机制",
	})
	items = append(items, vm.LiteItem{
		Category:  "hub",
		ItemName:  "BoltHisEventRunFuse",
		ItemValue: strconv.FormatBool(cfg.Read().Hub.BoltHisEventRunFuse),
		Remark:    "是否启动历史告警存储",
	})

	items = append(items, vm.LiteItem{
		Category:  "phoenix",
		ItemName:  "PhoenixServerIP",
		ItemValue: cfg.Read().Phoenix.PhoenixServerIP,
		Remark:    "Phoenix服务器地址",
	})

	items = append(items, vm.LiteItem{
		Category:  "phoenix",
		ItemName:  "PhoenixServerPort",
		ItemValue: cfg.Read().Phoenix.PhoenixServerPort,
		Remark:    "Phoenix服务器端口",
	})

	items = append(items, vm.LiteItem{
		Category:  "phoenix",
		ItemName:  "JobSendCOR",
		ItemValue: strconv.FormatBool(cfg.Read().Phoenix.JobSendCOR),
		Remark:    "是否启动实时数据存储到Phoenix",
	})

	items = append(items, vm.LiteItem{
		Category:  "phoenix",
		ItemName:  "BoltSendCOV",
		ItemValue: strconv.FormatBool(cfg.Read().Phoenix.BoltSendCOV),
		Remark:    "是否启动事件发送至Phoenix",
	})

	return items
}
