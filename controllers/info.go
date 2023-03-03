package controllers

import (
	"gin-framework-gs/models"
	"net/http"

	"gin-framework-gs/database"

	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Id    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ReadInfo(c *gin.Context) {

	// 전체 읽기
	var infoList []models.Info
	database.DB.Find(&infoList)

	// 특정 컬럼 읽기
	//var info models.Info
	//models.DB.First(&info, 3) // primary key 기준으로 info 찾기
	//models.DB.First(&info, "id = ?", 5) // id가 5인 info 찾기

	if infoList == nil {
		c.JSON(http.StatusNoContent, nil)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "infoList",
	})
}

func CreateInfo(c *gin.Context) {

	input := CreateInput{}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 생성하기
	info := models.Info{Id: input.Id, Name: input.Name, Email: input.Email}
	database.DB.Create(&info)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   input,
	})
}

func UpdateInfo(c *gin.Context) {
	var info models.Info
	if err := database.DB.Where("id = ?", c.Param("id")).First(&info).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	input := UpdateInput{}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	// 수정하기
	database.DB.Model(&info).Updates(input)
	// 수정하기 - 특정 필드 수정하기
	//models.DB.Model(&info).Update("Name", "Lee")
	//models.DB.Model(&info).Update(models.Info{Name: "Lee", Email: "rhdtha01@gmail.com"})
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   info,
	})

}

func DeleteInfo(c *gin.Context) {
	var info models.Info
	if err := database.DB.Where("id = ?", c.Param("id")).First(&info).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// 삭제하기
	database.DB.Delete(&info)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   info,
	})
}
