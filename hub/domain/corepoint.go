package domain

import (
	"ark/hub/enum"
	"ark/util/exp"
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//CorePoint point of domain coresource
type CorePoint struct {
	CorePointID      int    `xorm:"not null pk autoincr INT(11)"`
	PointName        string `xorm:"VARCHAR(128)"`
	Accuracy         string `xorm:"VARCHAR(32)"`
	Unit             string `xorm:"VARCHAR(32)"`
	Max              string `xorm:"VARCHAR(32)"`
	Min              string `xorm:"VARCHAR(32)"`
	CoreSourceID     int    `xorm:"not null INT(11)"`
	CoreDataTypeID   int    `xorm:"not null INT(11)"`
	EventSeverity    int    `xorm:"not null INT(11)"`
	StateRuleID      int    `xorm:"not null INT(11)"`
	Readable         bool   `xorm:"Bool"`
	Writable         bool   `xorm:"Bool"`
	Masked           bool   `xorm:"Bool"`
	DefaultValue     string `xorm:"VARCHAR(32)"`
	Step             string `xorm:"Float"`
	OriginStandardID string `xorm:"VARCHAR(20)"`
	StandardID       int    `xorm:"not null INT(11)"`
	Cron             string `xorm:"VARCHAR(128)"`
	Expression       string `xorm:"VARCHAR(255)"`
	UniqueID         string `xorm:"VARCHAR(128)"`

	GatewayID int
	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	//事件：开始，确认，结束
	CurrentEventState int
	//数据是否有效
	IsAvailabe bool
	//数据是否有效,0:ok,1:low limit exceed 2: up limit exceed
	LimitState int
	//缓存全部COV的包
	PacketLogs *list.List
	//缓存COV包中间的告警事件部分
	StateLogs *list.List
	//更新时间
	UpdateTime int64
	//常用的State开始对应的，开始时间，如果CurrentEventState可用，否则存的是上次的结束时间
	StartTime int64
	//常用的State结束对应的，结束时间，如果CurrentEndState可用，否则存的是上次的结束时间
	EndTime int64
	lock    sync.Mutex
}

//NewCorePoint new corepoint
func NewCorePoint(coresource *CoreSource) *CorePoint {
	corePoint := &CorePoint{GatewayID: coresource.GatewayID}
	corePoint.PacketLogs = list.New()
	corePoint.StateLogs = list.New()

	return corePoint
}

//UpdateStateLogs  update state
func (corePoint *CorePoint) UpdateStateLogs(cov *COV) {
	defer corePoint.lock.Unlock()
	corePoint.lock.Lock()

	if cov.IsValid == true {
		switch cov.StateID {
		case enum.CorePointFlagStart:
			corePoint.appendState(cov)
		case enum.CorePointFlagConfirm:
			corePoint.appendState(cov)
		case enum.CorePointFlagEnd:
			corePoint.appendState(cov)
		}
	}
}

func (corePoint *CorePoint) appendState(cov *COV) {

	corePoint.StateLogs.PushBack(cov) // Enqueue

	if corePoint.StateLogs.Len() > CorePacketLogQueueSize {
		e := corePoint.StateLogs.Front()

		if e != nil {
			corePoint.StateLogs.Remove(e)
		}
	}
}

//UpdateLimitState  自动计算模拟量的超限情况
func (corePoint *CorePoint) UpdateLimitState(cov *COV) {
	defer corePoint.lock.Unlock()
	corePoint.lock.Lock()

	if corePoint.CoreDataTypeID == 2 {
		corePoint.LimitState = 0
		if len(corePoint.Max) > 0 {
			max, err := strconv.ParseFloat(corePoint.Max, 32)
			exp.CheckError(err)
			if float32(max) < corePoint.CurrentNumericValue {
				corePoint.LimitState = 2
			}
		}
		if len(corePoint.Min) > 0 {
			min, err := strconv.ParseFloat(corePoint.Min, 32)
			exp.CheckError(err)
			if float32(min) > corePoint.CurrentNumericValue {
				corePoint.LimitState = 1
			}
		}
	}
}

//updatePacketLogs update packet logs
func (corePoint *CorePoint) updatePacketLogs(cov *COV) {

	corePoint.PacketLogs.PushBack(cov.Clone()) // Enqueue

	if corePoint.PacketLogs.Len() > CorePacketLogQueueSize {
		e := corePoint.PacketLogs.Front()

		if e != nil {
			corePoint.PacketLogs.Remove(e)
		}

	}
}

//GetCurrentValue GetCurrentValue
func (corePoint *CorePoint) GetCurrentValue() (string, error) {
	if !corePoint.DataIsReady() {
		return "", fmt.Errorf("data is not ready")
	}

	if corePoint.CoreDataTypeID == 1 {
		return fmt.Sprintf("%d", int(corePoint.CurrentNumericValue)), nil
	} else if corePoint.CoreDataTypeID == 2 {
		if len(corePoint.Accuracy) == 0 {
			return fmt.Sprintf("%.2f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil
		}
		dotPos := strings.Index(corePoint.Accuracy, ".")
		if dotPos < 0 {
			return fmt.Sprintf("%.2f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil
		}

		if dotPos == 1 {
			return fmt.Sprintf("%.1f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil
		}

		if dotPos == 2 {
			return fmt.Sprintf("%.2f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil
		}

		if dotPos == 3 {
			return fmt.Sprintf("%.3f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil
		}

		if dotPos == 4 {
			return fmt.Sprintf("%.4f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil
		}

		return fmt.Sprintf("%.2f %s", corePoint.CurrentNumericValue, corePoint.Unit), nil

	} else if corePoint.CoreDataTypeID == 3 {
		return corePoint.CurrentStringValue, nil
	} else {
		return "", fmt.Errorf("datatype not supported")
	}
}

//UpdateData update corepoint data
func (corePoint *CorePoint) UpdateData(cov *COV) {
	defer corePoint.lock.Unlock()
	corePoint.lock.Lock()

	corePoint.IsAvailabe = cov.IsValid
	corePoint.CurrentNumericValue = cov.CurrentNumericValue
	corePoint.CurrentStringValue = cov.CurrentStringValue
	corePoint.UpdateTime = cov.Timestamp
	corePoint.updatePacketLogs(cov)
}

//DataIsReady 实时数据是否可用(判断gateway，coresource和isvalid)
func (corePoint *CorePoint) DataIsReady() bool {

	gw, ok := Gateways.Get(corePoint.GatewayID)
	if ok {
		gateway := gw.(*Gateway)
		if gateway.ConState == enum.GatewayConStateOnline {
			cr, exist := gateway.CoreSources.Get(corePoint.CoreSourceID)
			if exist {
				coresource := cr.(*CoreSource)
				if coresource.State == enum.CoreSourceConStateOnline {
					if corePoint.IsAvailabe {
						return true
					}
				}
			}
		}
	}

	return false
}
