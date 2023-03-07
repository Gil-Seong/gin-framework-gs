package routers

import (
	"fmt"
	"gin-framework-gs/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestServer() *gin.Engine {
	router := gin.Default()

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
