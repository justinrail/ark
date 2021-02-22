package ws

import (
	"ark/hub/collector/cmb/dto"
	"ark/util/log"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
)

//Conversation 请求及处理过程
type Conversation struct {
	w       http.ResponseWriter
	r       *http.Request
	Request *dto.Request
	Done    chan bool
}

//NewConversation 创建新的处理会话
func NewConversation(rawBody string, w http.ResponseWriter, r *http.Request) (*Conversation, bool) {
	c := &Conversation{}
	c.w = w
	c.r = r
	c.parseRequest(rawBody)
	c.Done = make(chan bool)
	if c.Request == nil {
		return nil, false
	}

	return c, true
}

//SendResponse 返回数据
func (conversation *Conversation) SendResponse(packetType string, obj interface{}) {

	conversation.w.Header().Set("Content-Type", "text/xml")
	conversation.w.WriteHeader(http.StatusOK)

	msg, fine := MakeLSCReturnMessage(packetType, obj)

	if fine {
		conversation.w.Write([]byte(msg))
	}
	//conversation.w.Write([]byte(packetType))
	conversation.Done <- true
}

//GetClientIP GetClientIP
func (conversation *Conversation) GetClientIP() string {
	ip, _, err := net.SplitHostPort(conversation.r.RemoteAddr)
	if err == nil {
		return ip
	}
	return "NA"
}

//GetRequest 从数据包中解出要用的对象
func (conversation *Conversation) GetRequest(obj interface{}) {
	xmlContent := fmt.Sprintf("<Info>%s</Info>", conversation.Request.Info.InnerText)
	xml.Unmarshal([]byte(xmlContent), obj)
}

func (conversation *Conversation) parseRequest(rawBody string) {
	xmlContent, ok := GetSOAPXmlBodyString(rawBody)

	if ok {
		req := &dto.Request{}
		err := xml.Unmarshal([]byte(xmlContent), req)

		if err != nil {
			conversation.w.WriteHeader(http.StatusInternalServerError)
			return
		}
		conversation.Request = req
	} else {
		log.Error("fail to get xml body\n")
	}

}
