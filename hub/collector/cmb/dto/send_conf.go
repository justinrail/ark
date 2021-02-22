package dto

import "encoding/xml"

//SendConfigRequest SendConfigRequest
type SendConfigRequest struct {
	XMLName xml.Name `xml:"Info"`
	FSUID   string   `xml:"FSUID"`
	Values  Values   `xml:"Values"`
}

//SendConfigResponse SendConfigResponse
type SendConfigResponse struct {
	XMLName      xml.Name `xml:"Info"`
	FSUID        string   `xml:"FSUID"`
	Result       int      `xml:"Result"`
	FailureCause string   `xml:"FailureCause"`
}

//Values Values
type Values struct {
	XMLName xml.Name `xml:"Values"`
	Device  []Device
}

//Device Device
type Device struct {
	XMLName       xml.Name `xml:"Device"`
	DeviceID      string   `xml:"DeviceID,attr"`
	DeviceName    string   `xml:"DeviceName,attr"`
	SiteID        string   `xml:"SiteID,attr"`
	RoomID        string   `xml:"RoomID,attr"`
	SiteName      string   `xml:"SiteName,attr"`
	RoomName      string   `xml:"RoomName,attr"`
	DeviceType    string   `xml:"DeviceType,attr"`
	DeviceSubType string   `xml:"DeviceSubType,attr"`
	Model         string   `xml:"Model,attr"`
	Brand         string   `xml:"Brand,attr"`
	RatedCapacity string   `xml:"RatedCapacity,attr"`
	Version       string   `xml:"Version,attr"`
	BeginRunTime  string   `xml:"BeginRunTime,attr"`
	DevDescribe   string   `xml:"DevDescribe,attr"`
	ConfRemark    string   `xml:"ConfRemark,attr"`
	Signals       Signals  `xml:"Signals"`
}

//Signals Signals
type Signals struct {
	XMLName xml.Name `xml:"Signals"`
	Signal  []Signal
}

//Signal Signal
type Signal struct {
	XMLName      xml.Name `xml:"Signal"`
	Type         int      `xml:"Type,attr"`
	ID           string   `xml:"ID,attr"`
	SignalName   string   `xml:"SignalName,attr"`
	SignalNumber string   `xml:"SignalNumber,attr"`
	AlarmLevel   int      `xml:"AlarmLevel,attr"`
	Thresbhold   string   `xml:"Thresbhold,attr"`
	NMAlarmID    string   `xml:"NMAlarmID,attr"`
}
