package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Pallat"})

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var u User
		db.First(&u)
		c.JSON(200, u)
	})
	r.Run()
}
