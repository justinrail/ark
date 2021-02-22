package server

import (
	"ark/util/cfg"
	"ark/util/exe"
	"ark/util/log"
	"ark/web/controller"
	"ark/web/handler"
	"html/template"
	"path/filepath"
	"strconv"
	"time"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

//ArkWebServer ark web server
type ArkWebServer struct {
	Engine *gin.Engine
}

var webServer *ArkWebServer

func init() {
	webServer = prepare()
}

func prepare() *ArkWebServer {
	router := gin.Default()

	srv := &ArkWebServer{router}

	//template web for simple embedded
	loadTemplateWebSiteRoutes(srv)

	loadRESTAPIRoutes(srv)

	//serv static files
	staticPath := exe.Info().AppPath + filepath.FromSlash("/web/static/coreui")
	log.Info("web static root path: " + staticPath)
	//static middleware 解决不能使用/作为静态文件的serv虚拟目录问题
	srv.Engine.Use(static.Serve("/", static.LocalFile(staticPath, false)))
	if cfg.Read().App.WebGinDebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	}
	return srv
}

func loadRESTAPIRoutes(server *ArkWebServer) {

	api := server.Engine.Group("/api")
	{
		api.GET("/gatewaypackets/:gatewayId", controller.GetGatewayPacketByGatewayID)
		api.GET("/coresourcepackets/", controller.GetCoreSourcePacketByCoreSourceID)
		api.GET("/corepointpackets/", controller.GetCorePointPacketByCorePointID)
		api.GET("/corepointstatepackets/", controller.GetCorePointStatePacketByCorePointID)
		api.GET("/litepoints/", controller.GetLitePoints)
		api.GET("/complexindexs/:complexIndexId", controller.GetComplexIndexs)
		api.GET("/complexindexs/", controller.GetComplexIndexs)
		api.GET("/hiscomplexindexs/", controller.GetHisComplexIndexs)
		api.GET("/gateways/", controller.GetGateways)
		api.GET("/gateways/:gatewayId", controller.GetGatewayByID)
		api.DELETE("/gateways/:gatewayId", controller.DeleteGatewayByID)
		api.PUT("/gateways/:gatewayId", controller.UpdateGateway)
	}

}

func loadTemplateWebSiteRoutes(server *ArkWebServer) {
	templatePath := exe.Info().AppPath + filepath.FromSlash("/web/template/")
	//log.Info("web template path: " + templatePath)
	//TODO: Load template files recursive is not offical surported. So all temp files in same dir

	//new template engine
	server.Engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      templatePath,
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/header", "partials/footer", "partials/side"},
		Funcs: template.FuncMap{"formatAsDate": func(ctime int64) string {
			strtime := time.Unix(ctime, 0).Format("2006-01-02 15:04:05")
			return strtime
		}},
		DisableCache: true,
	})

	//with no gin-template plugin
	//templatePath := exe.Info().AppPath + filepath.FromSlash("/web/template/*.html")
	//server.Engine.LoadHTMLGlob(templatePath)

	server.Engine.NoRoute(handler.Error404Get)
	server.Engine.GET("/", handler.IndexGet)
	server.Engine.GET("/index", handler.IndexGet)
	server.Engine.GET("/log", handler.LogGet)
	server.Engine.GET("/config", handler.ConfigGet)
	server.Engine.GET("/runtime", handler.RuntimeGet)
	server.Engine.GET("/machine", handler.MachineGet)
	server.Engine.GET("/application", handler.ApplicationGet)
	server.Engine.GET("/domain", handler.DomainGet)
	server.Engine.GET("/gateway/:gatewayId", handler.GatewayGet)
	server.Engine.GET("/coresource", handler.CoreSourceGet)
	server.Engine.GET("/liveevent", handler.LiveEventGet)
	server.Engine.GET("/hisliveevent", handler.HisLiveEventGet)
	server.Engine.GET("/hispoint", handler.HisPointGet)
	server.Engine.GET("/topology", handler.TopologyGet)
	server.Engine.GET("/complexindex", handler.ComplexIndexGet)
	server.Engine.GET("/hiscomplexindex", handler.HisComplexIndexGet)
	server.Engine.GET("/stub", handler.StubGet)
	server.Engine.GET("/notifyrule", handler.NotifyRuleGet)
	server.Engine.GET("/notifymessage", handler.NotifyMessageGet)
	server.Engine.GET("/cmb", handler.CMBGet)
	server.Engine.GET("/removegateway", handler.DeleteGatewayByGatewayID)
}

//Start start ark embedded website server
func Start() {
	webServer.Engine.Run(":" + strconv.Itoa(cfg.Read().App.WebServerPort))
}
