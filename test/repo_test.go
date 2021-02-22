package test

import (
	"ark/store/mysql"
	"ark/store/mysql/repo"
	"testing"
)

func Test_GatewayRepo(t *testing.T) {

	mysql.OpenConnection()
	gws := repo.GetGatewaysByCollector("stub")

	if len(gws) > 0 {
		t.Log("passed")
	} else {
		t.Error("can not load repo data")
	}

}
