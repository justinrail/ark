package str

import "time"

//TimeToString conv time to timestamp style string
func TimeToString(ctime int64) string {
	strtime := time.Unix(ctime, 0).Format("2006-01-02 15:04:05")
	return strtime
}

//TimstampStringToTime convert timestamp to int64
func TimstampStringToTime(str string) int64 {

	tm2, _ := time.Parse("2006-01-02 15:04:05", str)

	return tm2.Unix()
}

//LocalTimstampStringToTime convert timestamp to int64
func LocalTimstampStringToTime(str string) int64 {

	local2, err2 := time.LoadLocation("Local")
	if err2 != nil {
		return -1
	}

	//tm2, _ := time.Parse("2006-01-02 15:04:05", str)
	tm2, _ := time.ParseInLocation("2006-01-02 15:04:05", str, local2)

	return tm2.Unix()
}

//NowString return now to timestamp style string
func NowString() string {
	return TimeToString(time.Now().Unix())
}

//TimeDurationString give string for time span
func TimeDurationString(timeStart int64, timeEnd int64) string {
	since := time.Unix(timeStart, 0).Sub(time.Unix(timeEnd, 0))
	return since.String()
}

//TimeDurationStringToNow give set time to now duration string
func TimeDurationStringToNow(ctime int64) string {

	return TimeDurationString(ctime, time.Now().Unix())
}
