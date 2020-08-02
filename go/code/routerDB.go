package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Problem_router struct {
	gorm.Model
	Question string
	Hint     string
	Anser    string
}

// DB接続
func dbInit_router() {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("dbInit_router失敗")
	}
	db.AutoMigrate(&Problem_router{})
	defer db.Close()
}

func check_router(id int, anser string) (Problem_router, string) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("router_check失敗")
	}
	var result string
	var router Problem_router
	if err := db.Where("id = ? AND anser = ?", id, anser).First(&router).Error; err != nil {
		result = "不正解"
	} else {
		result = "正解"
	}
	db.Close()
	return router, result
}

func routerGetAll() []Problem_router {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("データベース開けず(dbGetAll)")
	}
	var router []Problem_router
	db.Order("created_at desc").Find(&router)
	db.Close()
	return router
}

func routerGetOne(id int) Problem_router {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("データベース開けず(dbGetOne)")
	}
	var router Problem_router
	db.First(&router, id)
	db.Close()
	return router
}

func routerInsert(question string, anser string, hint string) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("routerInsert失敗")
	}
	db.Create(&Problem_router{Question: question, Anser: anser, Hint: hint})
	defer db.Close()
}

func routerUpdate(id int, question string, hint string, anser string) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("routerUpdate失敗")
	}
	var router Problem_router
	db.First(&router, id)
	router.Question = question
	router.Hint = hint
	router.Anser = anser
	db.Save(&router)
	db.Close()
}

func routerDelete(id int) {
	db, err := gorm.Open("mysql", "root:tmcit@/new_score?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("routerDelete失敗")
	}
	var router Problem_router
	db.Where("id = ?", id).Delete(&router)
	db.Close()
}
