package sqlite

import (
	"github.com/go-xorm/xorm"
	//Db connection defaul init
	_ "github.com/mattn/go-sqlite3"
)

//Store DB object
var store *Store

func init() {
	store = &Store{}
}

//Store sqlite db manage object
type Store struct {
	engine *xorm.Engine
}

//OpenConnection create or open database
func OpenConnection() {
	var err error
	store.engine, err = xorm.NewEngine("sqlite3", "./ark.db")
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
