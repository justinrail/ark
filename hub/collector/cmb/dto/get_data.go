package dto

import "encoding/xml"

//GetDataRequest GetDataRequest
type GetDataRequest struct {
	XMLName    xml.Name   `xml:"Info"`
	FSUID      string     `xml:"FSUID"`
	DeviceList DeviceList `xml:"DeviceList"`
}

//DeviceList DeviceList
type DeviceList struct {
	XMLName xml.Name `xml:"DeviceList"`
	Devices []GetDataDevice
}

//GetDataDevice GetDataDevice
type GetDataDevice struct {
	XMLName  xml.Name `xml:"Device"`
	DeviceID string   `xml:"ID,attr"`
	IDs      []GetDataSignalID
}

//GetDataSignalID GetDataSignalID
type GetDataSignalID struct {
	XMLName   xml.Name `xml:"ID"`
	InnerText string   `xml:",innerxml"`
}

//GetDataResponse GetDataResponse
type GetDataResponse struct {
	XMLName      xml.Name         `xml:"Info"`
	FSUID        string           `xml:"FSUID"`
	Result       int              `xml:"Result"`
	Values       GetDataAckValues `xml:"Values"`
	FailureCause string           `xml:"FailureCause"`
}

//GetDataAckValues GetDataAckValues
type GetDataAckValues struct {
	XMLName       xml.Name      `xml:"Values"`
	DeviceAckList DeviceAckList `xml:"DeviceList"`
}

//DeviceAckList DeviceAckList
type DeviceAckList struct {
	XMLName xml.Name `xml:"DeviceList"`
	Device  []GetDataAckDevice
}

//GetDataAckDevice GetDataAckDevice
type GetDataAckDevice struct {
	XMLName    xml.Name `xml:"Device"`
	DeviceID   string   `xml:"ID,attr"`
	TSemaphore []TSemaphore
}

//TSemaphore TSemaphore
type TSemaphore struct {
	XMLName      xml.Name `xml:"TSemaphore"`
	ID           string   `xml:"ID,attr"`
	SignalNumber string   `xml:"SignalNumber,attr"`
	Type         int      `xml:"Type,attr"`
	MeasuredVal  float32  `xml:"MeasuredVal,attr"`
	SetupVal     string   `xml:"SetupVal,attr"`
	Status       int      `xml:"Status,attr"`
	Time         string   `xml:"Time,attr"`
}
