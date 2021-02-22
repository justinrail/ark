package dto

import "encoding/xml"

//LoginRequest LoginRequest
type LoginRequest struct {
	XMLName  xml.Name `xml:"Info"`
	UserName string   `xml:"UserName"`
	PassWord string   `xml:"PassWord"`
	FSUID    string   `xml:"FSUID"`
	FSUIP    string   `xml:"FSUIP"`
	FSUMAC   string   `xml:"FSUMAC"`
	FSUVER   string   `xml:"FSUVER"`
}

//LoginResponse LoginRespone
type LoginResponse struct {
	XMLName      xml.Name `xml:"Info"`
	FSUID        string   `xml:"FSUID"`
	Result       int      `xml:"Result"`
	FailureCause string   `xml:"FailureCause"`
}

// {
// 	  // 创建编码器
// 	  buffer := new(bytes.Buffer)
// 	  enc := xml.NewEncoder(buffer)

// 	  // 设置缩进，这里为4个空格
// 	  enc.Indent("", "    ")

// 	  // 开始生成XML
// 	  startExtension := start("extension", attrmap{"name": "rtp_multicast_page"})
// 	  enc.EncodeToken(startExtension)
// 	  startCondition := start("condition", attrmap{"field": "destination_number",
// 		  "expression": "^pagegroup$|^7243$"})
// 	  enc.EncodeToken(startCondition)
// 	  startAction := start("action", attrmap{"application": "answer"})
// 	  enc.EncodeToken(startAction)
// 	  enc.EncodeToken(xml.CharData("raw text"))
// 	  enc.EncodeToken(startAction.End())
// 	  startAction = start("action", attrmap{"application": "esf_page_group"})
// 	  enc.EncodeToken(startAction)
// 	  enc.EncodeToken(startAction.End())
// 	  enc.EncodeToken(startCondition.End())
// 	  enc.EncodeToken(startExtension.End())

// 	  // 写入XML
// 	  enc.Flush()

// 	  // 打印结果
// 	  fmt.Println(buffer)
// }
