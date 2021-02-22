package ws

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const (
	//XMLHeader XML包头
	XMLHeader = "<xmlData>"
	//XMLFooter XML包尾
	XMLFooter = "</xmlData>"
	//FSUInvokeTemplate FSUInvokeTemplate
	FSUInvokeTemplate = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:fsus="http://FSUService.chinamobile.com">
	<soapenv:Header/>
	<soapenv:Body>
	   <fsus:invoke soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
		  <xmlData xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
		  
		  <?xml version="1.0" encoding="UTF-8"?>
 <Request>
	 <PK_Type>
		 <Name>%s</Name>
	 </PK_Type>
	 %s  
 </Request>		  </xmlData>
	   </fsus:invoke>
	</soapenv:Body>
 </soapenv:Envelope>`

	//LSCInvokeReturnTemplate InvokeReturnTemplate
	LSCInvokeReturnTemplate = `<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:ns1="http://FSUService.chinamobile.com">
	<SOAP-ENV:Header/>
	<SOAP-ENV:Body SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	   <ns1:invokeResponse>
		  <invokeReturn><![CDATA[<?xml version="1.0" encoding="UTF-8"?>
 <Response>
	 <PK_Type>
		 <Name>%s</Name>
	 </PK_Type>
	 %s
 </Response>]]></invokeReturn>
	   </ns1:invokeResponse>
	</SOAP-ENV:Body>
 </SOAP-ENV:Envelope>`
	FSUInvokeReturnHeader = `<invokeReturn>`
	FSUInvokeReturnFooter = `</invokeReturn>`
)

//GetSOAPXmlInvokeReturnString 获取SOAP包的XML返回内容区的XML片段
func GetSOAPXmlInvokeReturnString(content string) (string, bool) {

	body := strings.Replace(content, "&#xA;", "\n", -1)
	body = strings.Replace(body, "&lt;", "<", -1)
	body = strings.Replace(body, "&gt;", ">", -1)

	headerPos := strings.Index(body, FSUInvokeReturnHeader)
	if headerPos <= 0 {
		return "", false
	}

	footerPos := strings.Index(body, FSUInvokeReturnFooter)
	if footerPos <= 0 {
		return "", false
	}

	if headerPos >= footerPos {
		return "", false
	}

	res := body[headerPos+len(FSUInvokeReturnHeader) : footerPos]

	return res, true

}

//GetSOAPXmlBodyString 获取SOAP包的XML内容区的XML片段
func GetSOAPXmlBodyString(content string) (string, bool) {

	body := strings.Replace(content, "&#xA;", "\n", -1)
	body = strings.Replace(body, "&lt;", "<", -1)
	body = strings.Replace(body, "&gt;", ">", -1)

	headerPos := strings.Index(body, XMLHeader)
	if headerPos <= 0 {
		return "", false
	}

	footerPos := strings.Index(body, XMLFooter)
	if footerPos <= 0 {
		return "", false
	}

	if headerPos >= footerPos {
		return "", false
	}

	res := body[headerPos+len(XMLHeader) : footerPos]

	return res, true

}

//MakeLSCReturnMessage 生成返回消息
func MakeLSCReturnMessage(packetType string, obj interface{}) (string, bool) {

	x, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", false
	}

	res := fmt.Sprintf(LSCInvokeReturnTemplate, packetType, x)

	return res, true
}

//MakeFSUInvokeMessage 生成FSU调用消息
func MakeFSUInvokeMessage(packetType string, obj interface{}) (string, bool) {

	x, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", false
	}

	res := fmt.Sprintf(FSUInvokeTemplate, packetType, x)

	return res, true
}
