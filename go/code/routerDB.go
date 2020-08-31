package main

import (
	"fmt"

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
func dbInit_router() error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("dbInit_router失敗: %w", err)
	}
	defer db.Close()
	db.AutoMigrate(&Problem_router{})
	return nil
}

func check_router(id int, anser string) (Problem_router, string, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return Problem_router{}, "", fmt.Errorf("router_check失敗: %w", err)
	}
	defer db.Close()
	var result string
	var router Problem_router
	if err := db.Where("id = ? AND anser = ?", id, anser).First(&router).Error; err != nil {
		result = "不正解"
	} else {
		result = "正解"
	}
	return router, result, nil
}

func routerGetAll() ([]Problem_router, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return nil, fmt.Errorf("データベース開けず(dbGetAll): %w", err)
	}
	defer db.Close()
	var router []Problem_router
	err = db.Order("created_at desc").Find(&router).Error
	if err != nil {
		return nil, err
	}
	return router, nil
}

func routerGetOne(id int) (Problem_router, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return Problem_router{}, fmt.Errorf("データベース開けず(dbGetOne): %w", err)
	}
	defer db.Close()
	var router Problem_router
	db.First(&router, id)
	return router, nil
}

func routerInsert(question string, anser string, hint string) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("routerInsert失敗: %w", err)
	}
	defer db.Close()
	db.Create(&Problem_router{Question: question, Anser: anser, Hint: hint})
	return nil
}

func routerUpdate(id int, question string, hint string, anser string) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("routerUpdate失敗: %w", err)
	}
	defer db.Close()
	var router Problem_router
	db.First(&router, id)
	router.Question = question
	router.Hint = hint
	router.Anser = anser
	db.Save(&router)
	return nil
}

func routerDelete(id int) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("routerDelete失敗: %w", err)
	}
	defer db.Close()
	var router Problem_router
	db.Where("id = ?", id).Delete(&router)
	return nil
}
