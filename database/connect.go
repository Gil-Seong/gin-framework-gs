package database

import (
	"gin-framework-gs/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("mysql", "root:wjd0606@@@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Asia%2FSeoul")
	if err != nil {
		panic(err)
	}

	// table 명은 구조체명의 복수형(s)으로 자동 명명된다.
	// column 명은 구조체 필드의 소문자 snake case의 이름을 가진다.
	DB.AutoMigrate(&models.Info{}, &models.User1{})
	// DB.AutoMigrate(&TestModel2{})

	DB.DB().SetMaxIdleConns(11)  // idle connection pool(유휴 연결 풀)의 최대 수 설정
	DB.DB().SetMaxOpenConns(100) // 데이터베이스에 대한 열린 연결의 최대 수 설정
}
