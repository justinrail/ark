package main

import (
	"ark/hub"
	"ark/store/influxstore"
	"ark/store/mysql"
	"ark/util/cfg"
	"ark/web/server"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//init db
	mysql.OpenConnection()
	influxstore.OpenConnection()

	//start hub
	hub.Run()

	//prof server
	if cfg.Read().App.ProfileServerRunFuse {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}

	//start web server
	if cfg.Read().App.WebRunFuse {
		server.Start()
	}

}
