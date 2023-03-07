package routers

import (
	"gin-framework-gs/controllers"
	"gin-framework-gs/database"
	"gin-framework-gs/lib/jwt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func NewServer() *gin.Engine {

	// router := gin.New() // 커스텀이 필요하다면 New를 사용
	router = gin.Default()
	database.ConnectDatabase()

	router.SetFuncMap(template.FuncMap{})
	router.LoadHTMLGlob("templates/*.html")
	// router.LoadHTMLGlob("templates/*")

	swaggerConfig()

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "router list",
			"urlPath": router.Routes(),
		})
	})

	//인증
	// /v1
	v1 := router.Group("v1")
	auth := v1.Group("auth")
	auth.POST("/signin", controllers.Signin)
	auth.POST("/logout", controllers.Logout)
	auth.POST("/token-test", controllers.TokenTest)
	// auth.POST("/re-token", jwt.VerifyRefreshToken, jwt.CreateReissuanceToken, controllers.TokenTest)

	// /v1/product
	product := v1.Group("product")
	product.Use(jwt.VerifyAccessToken)
	{
		product.GET("/info", controllers.ReadInfo)
		product.POST("/info", controllers.CreateInfo)
		product.PUT("/info/:id", controllers.UpdateInfo)
		product.DELETE("/info/:id", controllers.DeleteInfo)
	}

	return router

}

// func (c *gin.Context) Next() // 미들웨어 내에서만 사용, 호출 핸들러 내부의 체인에서 보류 중인 핸들러를 실행한다.
// func (c *gin.Context) Abort() // 보류 중인 핸들러 호출을 방지한다. -> 여기서 response를 주고 다음 실행 예정인 핸들러를 실행시키지 않고 종료한다고 볼 수 있다.
// func (c *gin.Context) AbortWithStatusJSON(code int, json any) // Abort()호출 후 JSON을 호출한다. c.Abort() 후 c.JSON(code, json)
