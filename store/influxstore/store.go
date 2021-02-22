package influxstore

import (
	"ark/util/cfg"
	"ark/util/log"
	"fmt"
	"net/url"

	client "github.com/influxdata/influxdb1-client"
)

var conn *client.Client

// const (
// 	//UserName 用户名
// 	UserName = "admin"
// 	//Password 密码
// 	Password = "admin"
// 	//IP 数据库IP
// 	IP = "127.0.0.1"
// 	//Port 数据库端口
// 	Port = 8086
// 	//DatabaseName  数据库名
// 	DatabaseName = "ark"
// )

//OpenConnection 使用长连接
func OpenConnection() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%s", cfg.Read().App.InfluxDBServerIP, cfg.Read().App.InfluxDBServerPort))
	if err != nil {
		log.Error(err)
	}

	conf := client.Config{
		URL:      *host,
		Username: cfg.Read().App.InfluxDBServerUserName,
		Password: cfg.Read().App.InfluxDBServerPassword,
	}
	con, err := client.NewClient(conf)
	conn = con
	if err != nil {
		log.Error(err)
		return
	}

	Query(fmt.Sprintf("CREATE DATABASE %s", cfg.Read().App.InfluxDBServerDBName), cfg.Read().App.InfluxDBServerDBName)
	Query(fmt.Sprintf("create retention policy temp_flow on %s duration 1d replication 1 default", cfg.Read().App.InfluxDBServerDBName), cfg.Read().App.InfluxDBServerDBName)
	//Query("drop measurement hiscomplexindexs", DatabaseName)
	Query("drop measurement hiscoreliveevents", cfg.Read().App.InfluxDBServerDBName)
	// rps := Query(fmt.Sprintf("show retention policies on %s", DatabaseName), DatabaseName)
	// for _, row := range rps[0].Series[0].Values {
	// 	fmt.Println(row)
	// }
}

//Query 查询数据库
func Query(queryString string, dbName string) []client.Result {

	q := client.Query{
		Command:  queryString,
		Database: dbName,
	}

	response, err := conn.Query(q)

	if err == nil && response.Error() == nil {
		return response.Results
	}

	return nil
}

//Write 写数据库
func Write(pts []client.Point, dbName string, rpName string) {

	bps := client.BatchPoints{
		Points:          pts,
		Database:        dbName,
		RetentionPolicy: rpName,
	}
	_, err := conn.Write(bps)
	if err != nil {
		log.Error(err)
	}
}
