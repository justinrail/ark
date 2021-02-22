package job

import (
	"ark/hub/domain"
	"ark/store/mongostore"
	"ark/util/metrics"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

func sendCOR2Phoenix() {
	css := domain.CoreSources.Iter()

	for cs := range css {
		coreSouce := cs.Value.(*domain.CoreSource)

		cps := coreSouce.CorePoints.Iter()
		CPs := make([]*bson.M, 0)

		for cp := range cps {
			corePoint := cp.Value.(*domain.CorePoint)
			curValue, _ := corePoint.GetCurrentValue()
			if len(curValue) > 0 {
				pos := strings.Index(curValue, " ")
				if pos > 0 {
					curValue = curValue[0:pos]
				}
			}

			CP := &bson.M{
				"corePointId":  corePoint.CorePointID,
				"currentValue": curValue,
				"standardId":   corePoint.StandardID,
				"sampleTime":   time.Now(),
			}

			CPs = append(CPs, CP)
		}
		cor := bson.M{"_id": coreSouce.CoreSourceID, "sendTime": bson.Now(), "dataPoints": CPs}
		mongostore.UpsertID("phoenix", "COR", coreSouce.CoreSourceID, cor)
		metrics.AppMetrics.HubTopoPhoenixCORCounter.Inc(1)
	}
}
