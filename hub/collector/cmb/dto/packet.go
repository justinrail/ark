package dto

import "encoding/xml"

//Request 请求包
type Request struct {
	XMLName xml.Name `xml:"Request"`
	PKType  PKType   `xml:"PK_Type"`
	Info    Info     `xml:"Info"`
}

//Response 返回包
type Response struct {
	XMLName xml.Name `xml:"Response"`
	PKType  PKType   `xml:"PK_Type"`
	Info    Info     `xml:"Info"`
}

//PKType 包类型
type PKType struct {
	XMLName xml.Name `xml:"PK_Type"`
	Name    string   `xml:"Name"`
}

//Info Info
type Info struct {
	XMLName   xml.Name `xml:"Info"`
	InnerText string   `xml:",innerxml"`
}

// //Info 内容区
// type Info struct {
// 	XMLName      xml.Name `xml:"Info"`
// 	UserName     UserName `xml:"UserName,omitempty"`
// 	PassWord     string   `xml:"PassWord"`
// 	FSUID        string   `xml:"FSUID"`
// 	FSUIP        string   `xml:"FSUIP"`
// 	FSUMAC       string   `xml:"FSUMAC"`
// 	FSUVER       string   `xml:"FSUVER"`
// 	Result       int      `xml:"Result"`
// 	FailureCause string   `xml:"FailureCause"`
// }

// // //PKName PK_TYPE_NAME内容区
// // type PKName struct {
// // 	XMLName   xml.Name `xml:"Name"`
// // 	InnerText string   `xml:",innerxml"`
// // }

// //Response 返回包
// type Response struct {
// 	XMLName xml.Name `xml:"Request"`
// 	PKType  PKType   `xml:"PK_Type"`
// 	Info    string     `xml:"Info"`
// }
