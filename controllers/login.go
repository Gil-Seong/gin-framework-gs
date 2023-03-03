package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// DB 대신 값
var user = User{
	ID:       1,
	Username: "bono",
	Password: "wjd0606@@",
}

// func (user *User) LoginUser() bool {
//     db, err := ConnectToDb()
//     if err != nil {
//         log.Println(err)
//     }
//     defer db.Close()

//     _ = db.QueryRow(
//         "SELECT user_no, user_name "+
//             "FROM user_info "+
//             "WHERE user_id = $1 AND user_pw = $2 AND is_enabled = 1",
//         user.Email, sha512hash(user.Pw)).Scan(&user.UserNo, &user.Name)

//     if user.UserNo != 0 {
//         return true
//     } else {
//         return false
//     }
// }

func Login(c *gin.Context) {

	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	//JWT token 생성
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	c.Header("Last-Modified", time.Now().String())
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "-1")

	//쿠키 만료 시간 30분

	// 쿠키의 이름 = "access-token"
	// 값 = accessToken,
	// MaxAge = 1일 (초단위),
	// Path = ""(Root),
	// 도메인 = ""
	// Secure = false (https에서만 쿠키사용가능),
	// httpOnly = false (JavaScript가 접근하지 못하게하는 설정)
	c.SetCookie("access-token", token, 60*60*24, "", "", false, false)

	c.JSON(http.StatusOK, token)
	// c.JSON(http.StatusOK, token)

}

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
