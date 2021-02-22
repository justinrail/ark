package domain

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	"gopkg.in/jeevatkm/go-model.v1"
	"gopkg.in/oleiade/reflections.v1"
)

//NotifyTrigger 通知发生判断器
type NotifyTrigger struct {
	messageVarName   string
	messageClassName string
	messageFields    string
	targetVarName    string
	targetClassName  string
	targetFields     string
	tableVarName     string
	tableClassName   string
	whereExpression  string

	calcExpression *govaluate.EvaluableExpression
}

//Check 检查出发条件，并触发事件
func (notifyTrigger *NotifyTrigger) Check(notifyRule *NotifyRule) {
	if notifyTrigger.calcExpression != nil {
		parameters := make(map[string]interface{}, 2)

		rows, err := notifyTrigger.getTableObject(notifyTrigger.tableClassName)
		if err == nil {
			for _, row := range rows {
				parameters[notifyTrigger.tableVarName] = row
				result, _ := notifyTrigger.calcExpression.Evaluate(parameters)

				//如果表达式成功执行，则进行target留存或message发送
				if result.(bool) {
					notifyTrigger.postHandle(row, notifyRule)
				}
			}
		}
	}
}

func (notifyTrigger *NotifyTrigger) getTargetObject(obj interface{}) interface{} {
	var res interface{}

	switch notifyTrigger.targetClassName {
	case "CorePoint":
		if notifyTrigger.tableClassName == "CoreLiveEvent" {
			cp, ok := CorePoints.Get((obj.(*CoreLiveEvent)).CorePointID)
			if ok {
				//如果有预置参数，则进行设置
				if len(notifyTrigger.targetFields) > 0 {
					notifyTrigger.updateFields(notifyTrigger.targetFields, obj, cp)
				}
				res = cp
			}
		}
		if notifyTrigger.tableClassName == "CorePoint" {
			cp, ok := CorePoints.Get((obj.(*CorePoint)).CorePointID)
			if ok {
				res = cp
			}
		}
	case "CoreLiveEvent":
	}

	return res
}

func (notifyTrigger *NotifyTrigger) getMessageObject(target interface{}) interface{} {
	var res interface{}

	//emit message
	if len(notifyTrigger.messageClassName) > 0 {
		switch notifyTrigger.messageClassName {
		case "CoreLiteEvent":
			if notifyTrigger.targetClassName == "CorePoint" {
				cp, ok := CorePoints.Get((target.(*CorePoint)).CorePointID)
				if ok {
					//如果有预置参数，则进行设置
					cle := &CoreLiteEvent{}
					model.Copy(cle, cp)
					if len(notifyTrigger.messageFields) > 0 {
						notifyTrigger.updateFields(notifyTrigger.messageFields, cp, cle)
					}

					res = cle
				}
			}
		}
	}

	return res
}

//进行target留存或message发送
func (notifyTrigger *NotifyTrigger) postHandle(obj interface{}, notifyRule *NotifyRule) {
	firstFire := true

	//target add if not exist
	if len(notifyTrigger.tableClassName) > 0 {
		target := notifyTrigger.getTargetObject(obj)

		for iter := notifyRule.TargetObjects.Back(); iter != nil; iter = iter.Prev() {
			if target == iter.Value {
				firstFire = false
				break
			}
		}

		if firstFire {
			notifyRule.TargetObjects.PushBack(target)
			messageObj := notifyTrigger.getMessageObject(target)
			if messageObj != nil {
				notifyRule.PushMessage(messageObj)
			}
		}
	}
}

func (notifyTrigger *NotifyTrigger) updateFields(expression string, fromObj interface{}, toObj interface{}) {

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

func (notifyTrigger *NotifyTrigger) getTableObject(className string) ([]interface{}, error) {
	results := make([]interface{}, 0)

	switch className {
	case "CorePoint":
		cps := CorePoints.Iter()
		for kv := range cps {
			results = append(results, kv.Value)
		}
	case "CoreLiveEvent":
		cles := CoreLiveEvents.Iter()
		for kv := range cles {
			results = append(results, kv.Value)
		}
	}

	return results, nil
}

//NewNotifyTrigger 创建
func NewNotifyTrigger(str string) (*NotifyTrigger, error) {
	trigger := &NotifyTrigger{}

	err := trigger.parseTriggerSentense(str)
	if err == nil {
		//prepare expression

		if len(strings.TrimSpace(trigger.whereExpression)) == 0 {
			return nil, fmt.Errorf("new trigger: no valid where clause exists")
		}
		if len(strings.TrimSpace(trigger.tableClassName)) == 0 {
			return nil, fmt.Errorf("new trigger: no valid table class name exists")
		}
		if len(strings.TrimSpace(trigger.tableVarName)) == 0 {
			return nil, fmt.Errorf("new trigger: no valid table var name exists")
		}

		calcExpression, err2 := govaluate.NewEvaluableExpressionWithFunctions(trigger.whereExpression, trigger.expressionFuncs())

		if err2 == nil {
			trigger.calcExpression = calcExpression
		} else {
			fmt.Print(err2)
			return nil, fmt.Errorf("new trigger: fail to create calc expression")
		}

		return trigger, nil
	}

	return nil, fmt.Errorf("new trigger: fail to create trigger")
}

func (notifyTrigger *NotifyTrigger) parseTriggerSentense(str string) error {
	nt := notifyTrigger
	//TODO: 此处暂不支持中文字符计算（中文字符计算会作为2个字符）导致乱码
	fromPos := strings.Index(str, " from ")

	if fromPos > 0 {
		//fmt.Printf("%d, %d, %d\n", int(fromPos+6), len(str), len(str)-fromPos-6)

		wherePos := strings.Index(str, " where ")
		if wherePos > 0 {
			wherePart := str[wherePos+7 : len(str)]
			nt.whereExpression = strings.TrimSpace(wherePart)

			fromPart := str[fromPos+6 : wherePos]

			ss := strings.Split(strings.TrimSpace(fromPart), " ")
			nt.tableClassName = strings.TrimSpace(ss[0])
			nt.tableVarName = strings.TrimSpace(ss[1])
		} else {
			return fmt.Errorf("new trigger: not found where")
		}

		selectPos := strings.Index(str, " select ")
		if selectPos > 0 {
			selectPart := str[selectPos+8 : fromPos]
			ss := strings.Split(strings.TrimSpace(selectPart), " ")
			nt.targetClassName = strings.TrimSpace(ss[0])
			nt.targetVarName = strings.TrimSpace(ss[1])
			nt.targetFields = strings.TrimSpace(ss[2])

		} else {
			return fmt.Errorf("new trigger: not found select")
		}

		arcPos := strings.Index(str, " <= ")

		if arcPos > 0 {
			frontPart := str[0:arcPos]
			ss := strings.Split(strings.TrimSpace(frontPart), " ")
			nt.messageClassName = strings.TrimSpace(ss[0])
			nt.messageVarName = strings.TrimSpace(ss[1])
			nt.messageFields = strings.TrimSpace(ss[2])
		} else {
			return fmt.Errorf("new trigger: not found <=")
		}

	} else {
		return fmt.Errorf("new trigger: not found from")
	}

	return nil
}

func (notifyTrigger *NotifyTrigger) expressionFuncs() map[string]govaluate.ExpressionFunction {
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
