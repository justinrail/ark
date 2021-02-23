package mongostore

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// const (
// 	//UserName 用户名
// 	UserName = "admin"
// 	//Password 密码
// 	Password = "admin"
// 	//IP 数据库IP
// 	IP = "10.169.42.220"
// 	//Port 数据库端口
// 	Port = "27017"
// 	//DatabaseName  数据库名
// 	DatabaseName = "ark"
// )

//Connection mongoDB的session
var Connection *mgo.Session

// TODO: 因为mongodb会提前初始化，影响启动，所以先关掉初始化函数
// func init() {
// 	session, err := mgo.Dial(cfg.Read().App.MongoDBServerIP + ":" + cfg.Read().App.MongoDBServerPort)
// 	if err != nil {
// 		log.Error(err)
// 		session.Close()
// 		Connection = nil
// 		return
// 	}

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
// 	Connection = session
// 	//defer session.Close()
// }

//Insert 插入文档数据
func Insert(db string, table string, docs ...interface{}) error {
	if Connection == nil {
		return fmt.Errorf("Connection %s is not opened", db)
	}
	c := Connection.DB(db).C(table)
	err := c.Insert(docs)
	return err
}

//UpsertID 更新文档如果不存在的话插入文档数据
func UpsertID(db string, table string, id interface{}, doc interface{}) error {
	c := Connection.DB(db).C(table)
	_, err := c.UpsertId(id, doc)

	return err
}
