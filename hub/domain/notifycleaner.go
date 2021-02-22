package domain

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	"gopkg.in/jeevatkm/go-model.v1"
	"gopkg.in/oleiade/reflections.v1"
)

//NotifyCleaner 通知状态清除器
type NotifyCleaner struct {
	messageVarName   string
	messageClassName string
	messageFields    string
	whereExpression  string

	calcExpression *govaluate.EvaluableExpression
}

//Check 检查出发条件，并触发事件
func (notifyCleaner *NotifyCleaner) Check(notifyRule *NotifyRule) {
	if notifyCleaner.calcExpression != nil {

		needRemoves := make([]*list.Element, 0)

		for iter := notifyRule.TargetObjects.Back(); iter != nil; iter = iter.Prev() {
			parameters := make(map[string]interface{}, 2)
			parameters["target"] = iter.Value
			result, _ := notifyCleaner.calcExpression.Evaluate(parameters)
			//如果表达式成功执行，则进行message发送
			if result.(bool) {
				needRemoves = append(needRemoves, iter)
				notifyCleaner.postHandle(iter.Value, notifyRule)
			}
		}

		//清除target
		for _, obj := range needRemoves {
			notifyRule.TargetObjects.Remove(obj)
		}
	}
}

//进行message发送
func (notifyCleaner *NotifyCleaner) postHandle(obj interface{}, notifyRule *NotifyRule) {

	messageObj := notifyCleaner.getMessageObject(obj, notifyRule)
	if messageObj != nil {
		notifyRule.PushMessage(messageObj)
	}
}

func (notifyCleaner *NotifyCleaner) getMessageObject(target interface{}, notifyRule *NotifyRule) interface{} {
	var res interface{}

	//emit message
	if len(notifyCleaner.messageClassName) > 0 {
		switch notifyCleaner.messageClassName {
		case "CoreLiteEvent":
			if notifyRule.Trigger.targetClassName == "CorePoint" {
				cp, ok := CorePoints.Get((target.(*CorePoint)).CorePointID)
				if ok {
					//如果有预置参数，则进行设置
					cle := &CoreLiteEvent{}
					model.Copy(cle, cp)
					if len(notifyCleaner.messageFields) > 0 {
						notifyCleaner.updateFields(notifyCleaner.messageFields, cp, cle)
					}

					res = cle
				}
			}
		}
	}

	return res
}

func (notifyCleaner *NotifyCleaner) updateFields(expression string, fromObj interface{}, toObj interface{}) {

	setStrs := strings.Split(expression, ",")

	for _, str := range setStrs {
		fds := strings.Split(str, "=")
		toFieldStr := strings.TrimSpace(fds[0])
		fromFieldStr := strings.TrimSpace(fds[1])
		toFieldName := strings.TrimSpace(toFieldStr[strings.Index(toFieldStr, ".")+1 : len(toFieldStr)])

		if strings.Contains(fromFieldStr, ".") {
			fromFieldName := strings.TrimSpace(fromFieldStr[strings.Index(fromFieldStr, ".")+1 : len(fromFieldStr)])
			fromField, err := reflections.GetField(fromObj, fromFieldName)
			if err == nil {
				reflections.SetField(toObj, toFieldName, fromField)
			}
		} else {
			//如果不是对象属性设置则直接认为是数字
			i, err := strconv.Atoi(fromFieldStr)
			if err == nil {
				reflections.SetField(toObj, toFieldName, i)
			}
		}
	}
}

//NewNotifyCleaner 创建NotifyCleaner
func NewNotifyCleaner(str string) (*NotifyCleaner, error) {
	cleaner := &NotifyCleaner{}

	err := cleaner.parseCleanerSentense(str)
	if err == nil {
		if len(strings.TrimSpace(cleaner.whereExpression)) == 0 {
			return nil, fmt.Errorf("new cleaner: no valid where clause exists")
		}

		calcExpression, err2 := govaluate.NewEvaluableExpressionWithFunctions(cleaner.whereExpression, cleaner.expressionFuncs())

		if err2 == nil {
			cleaner.calcExpression = calcExpression
		} else {
			fmt.Print(err2)
			return nil, fmt.Errorf("new trigger: fail to create calc expression")
		}

		return cleaner, nil
	}

	return nil, fmt.Errorf("new trigger: fail to create trigger")
}

func (notifyCleaner *NotifyCleaner) parseCleanerSentense(str string) error {
	nc := notifyCleaner

	arcPos := strings.Index(str, " <= ")

	if arcPos > 0 {
		frontPart := str[0:arcPos]
		ss := strings.Split(strings.TrimSpace(frontPart), " ")
		nc.messageClassName = strings.TrimSpace(ss[0])
		nc.messageVarName = strings.TrimSpace(ss[1])
		nc.messageFields = strings.TrimSpace(ss[2])
		nc.whereExpression = str[arcPos+4 : len(str)]
	} else {
		nc.whereExpression = str
	}

	return nil
}

func (notifyCleaner *NotifyCleaner) expressionFuncs() map[string]govaluate.ExpressionFunction {
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
		//AVG, 获取一系列值的平均值
		"avg": func(args ...interface{}) (interface{}, error) {
			numcount := len(args)

			sum := 0.0
			if numcount > 0 {
				for _, v := range args {
					item := v.(float64)
					sum = item + sum
				}

				return sum / float64(numcount), nil
			}

			return -1, fmt.Errorf("avg item count is 0")
		},
		//now, get now unix int64 value
		"now": func(args ...interface{}) (interface{}, error) {
			return (float64)(time.Now().Unix()), nil
		},
		//duration, get duration between two time
		"duration": func(args ...interface{}) (interface{}, error) {

			tm := time.Unix(int64(args[0].(float64)), 0)
			tm2 := time.Unix(int64(args[1].(float64)), 0)
			return tm2.Sub(tm).Seconds(), nil
		},
	}
	return functions
}
