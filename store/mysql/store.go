package mysql

import (
	"ark/util/cfg"
	"ark/util/log"
	"os"
	"strings"

	//init mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// //数据库配置
// const (
// 	userName = "root"
// 	password = "siteweb1!"
// 	ip       = "127.0.0.1"
// 	port     = "3306"
// 	dbName   = "ark"
// )

//Store DB object
var store *Store

func init() {
	store = &Store{}
}

//Store mysql db manage object
type Store struct {
	Engine *xorm.Engine
}

//Engine return engine
func Engine() *xorm.Engine {
	return store.Engine
}

//OpenConnection create or open database
func OpenConnection() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{
		cfg.Read().App.MySQLServerUserName,
		":",
		cfg.Read().App.MySQLServerPassword,
		"@tcp(",
		cfg.Read().App.MySQLServerIP,
		":",
		cfg.Read().App.MySQLServerPort,
		")/",
		cfg.Read().App.MySQLServerDBName, "?charset=utf8"}, "")
	var err error
	store.Engine, err = xorm.NewEngine("mysql", path)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	store.Engine.ImportFile("conf/init.sql")

	//连接测试
	// if err := store.Engine.Ping(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//日志打印SQL
	//store.Engine.ShowSQL(true)

	store.Engine.SetMapper(core.SameMapper{})

	//设置连接池的空闲数大小
	store.Engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	store.Engine.SetMaxOpenConns(5)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1024*1024)
	store.Engine.SetDefaultCacher(cacher)

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
