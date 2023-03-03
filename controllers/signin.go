package controllers

import (
	"fmt"
	"gin-framework-gs/lib/jwt"

	"github.com/gin-gonic/gin"
)

// 하드코딩 테스트 방법
type TokenUser struct {
	ID        uint64 `json:"id"`
	UserID    string `json:"userid"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	IsManager bool   `json:"ismanager"`
}

// 하드코딩 테스트 방법
var tokenUser = TokenUser{
	ID:        1,
	UserID:    "bono915",
	Name:      "bono",
	Password:  "wjd0606@@",
	IsManager: true,
}

// SigninParam - 파라미터 형식 구조체
type SigninParam struct {
	ID string `json:"id" form:"id" query:"id"`
	PW string `json:"pw" form:"pw" query:"pw"`
}

// Signin - 로그인 메서드
func Signin(c *gin.Context) {
	u := new(SigninParam)

	if err := c.Bind(u); err != nil {
		return
	}

	fmt.Println(u.ID, u.PW)

	//DB 방식
	// User := &database.User{}
	// err := database.DB.Where("user_id = ? AND pw = ?", u.ID, u.Pw).Find(User).Error
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"status":       500,
	// 		"message":      "일치하는 회원이 없습니다.",
	// 		"refreshToken": "null",
	// 		"accessToken":  "null",
	// 	})
	// 	return
	// }

	//하드코딩 테스트 방법
	if tokenUser.UserID != u.ID || tokenUser.Password != u.PW {
		c.JSON(500, gin.H{
			"status":       500,
			"message":      "일치하는 회원이 없습니다.",
			"refreshToken": "null",
			"accessToken":  "null",
		})
		return
	}

	fmt.Println(tokenUser.UserID, tokenUser.Name)

	refreshToken, err := jwt.CreateRefreshToken(tokenUser.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       500,
			"message":      "refreshtoken 생성 중 에러",
			"refreshToken": "null",
			"accessToken":  "null",
		})
		return
	}

	accessToken, err := jwt.CreateAccessToken(tokenUser.Name, tokenUser.IsManager)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       500,
			"message":      "accesstoken 생성 중 에러",
			"refreshToken": refreshToken,
			"accessToken":  "null",
		})
		return
	}
	// c.SetCookie("access-token", accessToken, 60*60, "/", "localhost:3000", false, true)
	// c.SetCookie("refresh-token", refreshToken, 60*60*24*30, "/", "localhost:3000", false, true)

	//쿠키 만료 시간 30분

	// 쿠키의 이름 = "access-token"
	// 값 = accessToken,
	// MaxAge = 1일 (초단위),
	// Path = ""(Root),
	// 도메인 = ""
	// Secure = false (https에서만 쿠키사용가능),
	// httpOnly = false (JavaScript가 접근하지 못하게하는 설정)
	c.SetCookie("accessToken", accessToken, 60*60, "", "", false, false)
	c.SetCookie("refreshToken", refreshToken, 60*60*24*30, "", "", false, false)

	c.JSON(200, gin.H{
		"status":       200,
		"message":      "토큰 발급 완료",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
	return
}
