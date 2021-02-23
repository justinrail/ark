package cfg

import (
	"fmt"

	"github.com/koding/multiconfig"
)

//arkConfig global config object
type arkConfig struct {
	App     App
	Hub     Hub
	Phoenix Phoenix
}

type (
	// App ark的自身通用配置
	App struct {
		WebRunFuse             bool   `default:"true"`
		ProfileServerRunFuse   bool   `default:"false"`
		WebGinDebugMode        bool   `default:"false"`
		WebServerPort          int    `default:"80"`
		MySQLServerIP          string `default:"127.0.0.1"`
		MySQLServerPort        string `default:"3306"`
		MySQLServerUserName    string `default:"root"`
		MySQLServerPassword    string `default:"root"`
		MySQLServerDBName      string `default:"ark"`
		InfluxDBServerIP       string `default:"127.0.0.1"`
		InfluxDBServerPort     string `default:"8086"`
		InfluxDBServerUserName string `default:"admin"`
		InfluxDBServerPassword string `default:"admin"`
		InfluxDBServerDBName   string `default:"ark"`
		MongoDBServerIP        string `default:"10.169.42.220"`
		MongoDBServerPort      string `default:"27017"`
		MongoDBServerUserName  string `default:"admin"`
		MongoDBServerPassword  string `default:"admin"`
		MongoDBServerDBName    string `default:"ark"`
		JSONDBPath             string `default:"."`
	}

	// Hub 采集Hub的核心配置
	Hub struct {
		CollectorCMBRunFuse       bool `required:"true"`
		CollectorStubRunFuse      bool `required:"true"`
		JobHisPointRunFuse        bool `required:"true"`
		JobHisComplexIndexRunFuse bool `required:"true"`
		JobComplexIndexRunFuse    bool `required:"true"`
		NotifyRunFuse             bool `required:"true"`
		BoltHisEventRunFuse       bool `required:"true"`
	}

	//Phoenix 一个第三方系统
	Phoenix struct {
		PhoenixServerIP   string `default:"10.169.42.89"`
		PhoenixServerPort string `default:"8100"`
		JobSendCOR        bool   `required:"true"`
		BoltSendCOV       bool   `required:"true"`
	}
)

var cfg *arkConfig

func init() {
	cfg = load()
}

//Load load config from file
func load() *arkConfig {
	m := multiconfig.NewWithPath("conf/config.toml") // supports TOML, JSON and YAML

	// Get an empty struct for your configuration
	arkConf := new(arkConfig)

	// Populated the serverConf struct
	err := m.Load(arkConf) // Check for error
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	m.MustLoad(arkConf) // Panic's if there is any error

	return arkConf
}

//Read 获取配置
func Read() *arkConfig {
	return cfg
}

//TODO: Config file use conf as the config file directory
