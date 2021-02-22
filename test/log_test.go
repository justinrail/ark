package test

import (
	"ark/util/log"
	"testing"
)

func Test_log(t *testing.T) {
	log.Info("DDDD")

	t.Log("log existed")
}

func Benchmark_log(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间

	for i := 0; i < 10000; i++ {
		log.Info("test %d")
	}

}
