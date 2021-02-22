package domain

import (
	"ark/util/log"
	"fmt"

	"github.com/Knetic/govaluate"
)

//ComplexIndex 复杂指标计算，类似VDS
type ComplexIndex struct {
	ComplexIndexID   int
	ComplexIndexName string
	Category         string
	Title            string
	Label            string
	ObjectTypeID     int
	BusinessID       string
	GlobalResourceID int
	CalcCron         string
	//0:默认的原始数据，1:SaveCron时的上次值之差(now-last)， 12:差值计算后进行后期处理，10：原始数据采集后进行后期处理
	CalcType int
	//后期处理Expression
	AfterCalc  string
	SaveCron   string
	Expression string
	Unit       string
	Remark     string
	//上次的值,用于暂存差值计算需要的上次原始值，非最终值
	LastValue float64
	//上次采集时间
	LastTimestamp int64
	//当前数据采集时间
	Timestamp int64
	//当前值
	CurrentValue float64

	//当前值是否有效
	IsValid bool

	evaluableExpression      *govaluate.EvaluableExpression
	afterEvaluableExpression *govaluate.EvaluableExpression
}

//Init 初始化
func (ci *ComplexIndex) Init() {
	calcExpression, err := govaluate.NewEvaluableExpressionWithFunctions(ci.Expression, ci.expressionFuncs())

	if err == nil {
		ci.evaluableExpression = calcExpression
	} else {
		log.Error(fmt.Sprintf("ComplexIndexID: %d Name:%s Err: %s", ci.ComplexIndexID, ci.ComplexIndexName, err))
	}

	if len(ci.AfterCalc) > 0 {
		afterExpression, err2 := govaluate.NewEvaluableExpressionWithFunctions(ci.AfterCalc, ci.expressionFuncs())

		if err2 == nil {
			ci.afterEvaluableExpression = afterExpression
		}
	}

}

//Calculate 计算实时值
func (ci *ComplexIndex) Calculate(calcTime int64) {
	if ci.evaluableExpression != nil {
		result, _ := ci.evaluableExpression.Evaluate(nil)
		//如果是原始值，则直接存储返回
		if ci.CalcType == 0 {
			ci.CurrentValue = result.(float64)
			ci.Timestamp = calcTime
			//只有第一次计算成功时，取到的值才是真实的
			ci.IsValid = true

		} else if ci.CalcType == 1 || ci.CalcType == 12 {
			//如果是差值，或者差值后计算，则比对上次的值，进行差值计算
			if ci.LastTimestamp > 0 {
				ci.CurrentValue = result.(float64) - ci.LastValue

				//如果计算成功，需要后期处理，则进行后期处理
				if ci.CalcType == 12 {
					postVal, _ := ci.afterEvaluableExpression.Evaluate(nil)
					ci.CurrentValue = postVal.(float64)
				}
				ci.Timestamp = calcTime
				ci.IsValid = true

			} else {
				ci.IsValid = false
			}

			ci.LastValue = result.(float64)
			ci.LastTimestamp = calcTime
		}
	}
}

func (ci *ComplexIndex) expressionFuncs() map[string]govaluate.ExpressionFunction {
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
		//获取corepoint的实时数字值的简写函数
		"cp": func(args ...interface{}) (interface{}, error) {
			id := (int)(args[0].(float64))
			cphs, ok := CorePoints.Get(id)
			if ok {
				return (float64)((cphs.(*CorePoint)).CurrentNumericValue), nil
			}

			return -1, fmt.Errorf("complex index func cp: %d core point signal not found", id)

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
		//selfval, 获取当前值
		"selfval": func(args ...interface{}) (interface{}, error) {
			return ci.CurrentValue, nil
		},
		//powerfee, 电费计算函数（暂时没有具体的算法，后期加入）
		"powerfee": func(args ...interface{}) (interface{}, error) {
			val := args[0].(float64)

			return (float64)(val * 10), nil
		},
	}
	return functions
}
