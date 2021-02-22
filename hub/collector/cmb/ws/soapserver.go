package ws

import (
	"ark/util/exe"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

//DataIn 数据处理入口队列
var DataIn chan *Conversation

func init() {
	DataIn = make(chan *Conversation)
}

func newSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", soapHandler)
	return mux
}

//NewSOAPServer create soap server
func NewSOAPServer(addr string) *http.Server {
	mux := newSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}
	return server
}

func soapHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cvs, ok := NewConversation(string(rawBody), w, r)
		//fmt.Println(string(rawBody))
		if ok {
			DataIn <- cvs
			<-cvs.Done
		}

	case "GET":
		if strings.ToLower(r.RequestURI) == "/services/lscservice?wsdl" {
			filename := exe.Info().AppPath + filepath.FromSlash("/conf/LSCService.wsdl")
			body, _ := ioutil.ReadFile(filename)
			w.Header().Set("Content-Type", "text/xml")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		}
	default:
		// this will be a soap fault !?
		w.Write([]byte("this is a soap server\n"))
	}

}
