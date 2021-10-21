// go-swagger example
//
// Base go-swagger
//
//     	Schemes: http
//     	BasePath: /api
//     	Version: 0.1.0
//
//     	Consumes:
//     	- application/json
//     	Produces:
//     	- application/json
//
//
// swagger:meta
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"

	_ "github.com/matrixik/swag-perf/swagger/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title swaggo example
// @version 0.1.0
// @description.markdown
// @BasePath /api
func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("", mainHandler)
	}

	r.Static("/swagger-ui", "swagger-ui/")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// swagger:model mspIdentityResponse
type _ struct {
	// Response with identity data
	// in: body
	msp.IdentityResponse
}

// swagger:route GET / example
//
// go-swagger docs summary
//
// go-swagger docs description
//
// responses:
//   200: mspIdentityResponse
//
// @Summary swaggo docs summary
// @Description swaggo docs description
// @Success 200 {object} msp.IdentityResponse{attributes=msp.Attribute}
// @Produce json
// @Router / [get]
// @T
func mainHandler(c *gin.Context) {
	resp := msp.IdentityResponse{}
	c.JSON(http.StatusOK, resp)
}
