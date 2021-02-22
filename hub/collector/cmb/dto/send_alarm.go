package dto

import "encoding/xml"

//SendAlarmRequest SendAlarmRequest
type SendAlarmRequest struct {
	XMLName xml.Name    `xml:"Info"`
	FSUID   string      `xml:"FSUID"`
	Values  AlarmValues `xml:"Values"`
}

//SendAlarmResponse SendAlarmResponse
type SendAlarmResponse struct {
	XMLName      xml.Name `xml:"Info"`
	FSUID        string   `xml:"FSUID"`
	Result       int      `xml:"Result"`
	FailureCause string   `xml:"FailureCause"`
}

//AlarmValues AlarmValues
type AlarmValues struct {
	XMLName    xml.Name `xml:"Values"`
	TAlarmList TAlarmList
}

//TAlarmList TAlarmList
type TAlarmList struct {
	XMLName xml.Name `xml:"TAlarmList"`
	TAlarm  []TAlarm
}

//TAlarm TAlarm
type TAlarm struct {
	XMLName      xml.Name `xml:"TAlarm"`
	SerialNo     string   `xml:"SerialNo"`
	ID           string   `xml:"ID"`
	DeviceID     string   `xml:"DeviceID"`
	NMAlarmID    string   `xml:"NMAlarmID"`
	AlarmTime    string   `xml:"AlarmTime"`
	AlarmLevel   int      `xml:"AlarmLevel"`
	AlarmFlag    string   `xml:"AlarmFlag"`
	AlarmDesc    string   `xml:"AlarmDesc"`
	EventValue   float32  `xml:"EventValue"`
	SignalNumber string   `xml:"SignalNumber"`
	AlarmRemark  string   `xml:"AlarmRemark"`
}
