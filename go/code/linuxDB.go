package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Problem_linux struct {
	gorm.Model
	Question string
	Hint     string
	Anser    string
}

// DB接続
func dbInit_linux() {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("dbInit_linux失敗")
	}
	db.AutoMigrate(&Problem_linux{})
	defer db.Close()
}

func check_linux(id int, anser string) (Problem_linux, string) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("linux_check失敗")
	}
	var result string
	var linux Problem_linux
	if err := db.Where("id = ? AND anser = ?", id, anser).First(&linux).Error; err != nil {
		result = "不正解"
	} else {
		result = "正解"
	}
	db.Close()
	return linux, result
}

func linuxGetAll() []Problem_linux {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("データベース開けず(dbGetAll)")
	}
	var linux []Problem_linux
	db.Order("created_at desc").Find(&linux)
	db.Close()
	return linux
}

func linuxGetOne(id int) Problem_linux {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("データベース開けず(dbGetOne)")
	}
	var linux Problem_linux
	db.First(&linux, id)
	db.Close()
	return linux
}

func linuxInsert(question string, anser string, hint string) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("linuxInsert失敗")
	}
	db.Create(&Problem_linux{Question: question, Anser: anser, Hint: hint})
	defer db.Close()
}

func linuxUpdate(id int, question string, hint string, anser string) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("linuxUpdate失敗")
	}
	var linux Problem_linux
	db.First(&linux, id)
	linux.Question = question
	linux.Anser = anser
	linux.Hint = hint
	db.Save(&linux)
	db.Close()
}

func linuxDelete(id int) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("linuxDelete失敗")
	}
	var linux Problem_linux
	db.Where("id = ?", id).Delete(&linux)
	db.Close()
}
