package routers

import (
	"fmt"
	"gin-framework-gs/controllers"
	"gin-framework-gs/database"
	"gin-framework-gs/lib/jwt"
	"gin-framework-gs/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestServer() *gin.Engine {
	router := gin.Default()
	database.ConnectDatabase()

	router.SetFuncMap(template.FuncMap{})
	router.LoadHTMLGlob("templates/*.html")
	// router.LoadHTMLGlob("templates/*")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "메인페이지 입니다.",
			"message": "제목",
		})
	})

	router.GET("/user", func(c *gin.Context) {
		var userObj models.User1
		if err := c.ShouldBindQuery(&userObj); err == nil {
			fmt.Printf("user obj - %+v \n", userObj)
		} else {
			fmt.Printf("error - %+v \n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	router.POST("/user", func(c *gin.Context) {
		var userObj models.User2
		if err := c.ShouldBindJSON(&userObj); err == nil {
			fmt.Printf("user obj - %+v \n", userObj)
		} else {
			fmt.Printf("error - %+v \n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	router.PUT("/user/:id/:name/:email", func(c *gin.Context) {
		var userObj models.User3
		if err := c.ShouldBindUri(&userObj); err == nil {
			fmt.Printf("user obj uri binding - %+v \n", userObj)
		} else {
			fmt.Printf("error - %+v \n", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	router.PUT("/user", func(c *gin.Context) {
		var userObj models.User4
		if err := c.ShouldBindHeader(&userObj); err == nil {
			fmt.Printf("user obj - %+v \n", userObj)
		} else {
			fmt.Printf("error - %+v \n", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})

	router.GET("/user2/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action

		c.String(http.StatusOK, message)
	})

	router.POST("/add", func(c *gin.Context) {
		// var req = &Bind{}
		var data models.TestModel
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
			fmt.Println("data.id : ", data.Id)
			fmt.Println("data.name : ", data.Name)

		}
	})

	router.GET("/:name", func(c *gin.Context) { // :은 gin에게 url 이후에 오는 것이 name 변수로 받아진다는 것
		var val = c.Param("name") // 파라미터 name의 값을 변수 val의 값으로 초기화
		c.JSON(http.StatusOK, gin.H{
			"value": val,
		})

	})

	return router
}

func NewServer() *gin.Engine {

	router := gin.Default()
	database.ConnectDatabase()

	router.SetFuncMap(template.FuncMap{})
	router.LoadHTMLGlob("templates/*.html")

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

func TokenAuthMiddleware(c *gin.Context) {
	authToken := c.Request.Header.Get("auth-token")
	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No token",
		})
		return
	}

	if authToken != "secret-token" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}

	log.Println("authenticateMiddleware passing")
	c.Next()
	log.Println("authenticateMiddleware passed already")
}

// func (c *gin.Context) Next() // 미들웨어 내에서만 사용, 호출 핸들러 내부의 체인에서 보류 중인 핸들러를 실행한다.
// func (c *gin.Context) Abort() // 보류 중인 핸들러 호출을 방지한다. -> 여기서 response를 주고 다음 실행 예정인 핸들러를 실행시키지 않고 종료한다고 볼 수 있다.
// func (c *gin.Context) AbortWithStatusJSON(code int, json any) // Abort()호출 후 JSON을 호출한다. c.Abort() 후 c.JSON(code, json)
