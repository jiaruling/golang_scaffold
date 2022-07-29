package api

import (
	"GolangScaffold/docs"
	"GolangScaffold/global"
	"GolangScaffold/service/api/base"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func RegisterRouter() {
	if global.ServerConfig.S.Env == "dev" {
		docs.SwaggerInfo.Title = global.ServerConfig.SW.Title
		docs.SwaggerInfo.Description = global.ServerConfig.SW.Desc
		docs.SwaggerInfo.Version = global.ServerConfig.SW.Version
		docs.SwaggerInfo.Host = global.ServerConfig.SW.Host
		docs.SwaggerInfo.BasePath = global.ServerConfig.SW.BasePath
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		global.GinRouter.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	base.Router()
}
