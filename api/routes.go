package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func (server *Server) SetupRoutes(version string) http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger ERP Documentation"
	docs.SwaggerInfo.Description = "This is an ERP Backend with Golang."
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/* CORS */
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/sites", server.generateSite)
		v1.GET("/sites/:key", server.getSite)
	}

	return r
}
