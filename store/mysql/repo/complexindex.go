package repo

import (
	"ark/store/entity"
	"ark/store/mysql"
)

//GetAllComplexIndex 获取所有complexIndex
func GetAllComplexIndex() []entity.ComplexIndex {
	complexindexs := make([]entity.ComplexIndex, 0)

	err := mysql.Engine().Find(&complexindexs)
	checkErr(err)

	return complexindexs
}
