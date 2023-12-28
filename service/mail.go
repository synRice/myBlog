package service

import (
	"fmt"
	models "myBlog/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMessage(username string) string {
	dsn := "root:synrice@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	var result models.TestModel

	if err := db.Where("username = ?", "synrice").First(&result).Error; err != nil {
		fmt.Println("Error:", err) // 调试用
		if err == gorm.ErrRecordNotFound {
			return ("Record not found")
		} else {
			return ("Error")
		}
	} else {
		// 输出查询到的结果
		return result.Content
	}
}
