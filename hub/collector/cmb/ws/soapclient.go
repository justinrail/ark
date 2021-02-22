package ws

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//SOAPInvoke client invoke
func SOAPInvoke(ip string, cmdToken string, reqObject interface{}) string {
	str, ok := MakeFSUInvokeMessage(cmdToken, reqObject)
	if ok {
		url := fmt.Sprintf("http://%s:8080/services/FSUService?wsdl", ip)
		//fmt.Println(url)
		//fmt.Println(str)
		resp, err2 := http.Post(url, "text/xml", strings.NewReader(str))
		if err2 == nil {
			body, err3 := ioutil.ReadAll(resp.Body)
			if err3 == nil {
				content, correct := GetSOAPXmlInvokeReturnString(string(body))
				if correct {
					return content
				}
			}
		}
		defer resp.Body.Close()
	}

	return ""
}
