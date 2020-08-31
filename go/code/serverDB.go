package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Problem_server struct {
	gorm.Model
	Question string
	Hint     string
	Anser    string
}

// DB接続
func dbInit_server() error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("dbInit_linux失敗: %w", err)
	}
	defer db.Close()
	db.AutoMigrate(&Problem_server{})
	return nil
}

func check_server(id int, anser string) (Problem_server, string, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return Problem_server{}, "", fmt.Errorf("server_check失敗: %w", err)
	}
	defer db.Close()
	var result string
	var server Problem_server
	if err := db.Where("id = ? AND anser = ?", id, anser).First(&server).Error; err != nil {
		result = "不正解"
	} else {
		result = "正解"
	}
	return server, result, nil
}

func serverGetAll() ([]Problem_server, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return nil, fmt.Errorf("データベース開けず(dbGetAll): %w", err)
	}
	defer db.Close()
	var server []Problem_server
	err = db.Order("created_at desc").Find(&server).Error
	if err != nil {
		return nil, err
	}
	return server, nil
}

func serverGetOne(id int) (Problem_server, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return Problem_server{}, fmt.Errorf("データベース開けず(dbGetOne): %w", err)
	}
	defer db.Close()
	var server Problem_server
	db.First(&server, id)
	return server, nil
}

func serverInsert(question string, anser string, hint string) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("serverInsert失敗: %w", err)
	}
	defer db.Close()
	db.Create(&Problem_server{Question: question, Anser: anser, Hint: hint})
	return nil
}

func serverUpdate(id int, question string, hint string, anser string) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("serverUpdate失敗: %w", err)
	}
	defer db.Close()
	var server Problem_server
	db.First(&server, id)
	server.Question = question
	server.Hint = hint
	server.Anser = anser
	db.Save(&server)
	return nil
}

func serverDelete(id int) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("serverDelete失敗: %w", err)
	}
	defer db.Close()
	var server Problem_server
	db.Where("id = ?", id).Delete(&server)
	return nil
}
