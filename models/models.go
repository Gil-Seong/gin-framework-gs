package models

import (
	"time"
)

type CustomModel struct {
	// ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type TestModel struct {
	CustomModel
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type User1 struct {
	// gorm.Model
	CustomModel
	Id    int    `form:"id"`
	Name  string `form:"name"`
	Email string `form:"email"`
}
type User2 struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type User3 struct {
	Id    int    `uri:"id"`
	Name  string `uri:"name"`
	Email string `uri:"email"`
}
type User4 struct {
	Id int `header:"id"`
}

type User struct {
	Id       string
	Email    string
	Password string
}

type Info struct {
	CustomModel
	Id    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
