package main

import (
	"fmt"

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
func dbInit_linux() error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("dbInit_linux失敗: %w", err)
	}
	defer db.Close()
	db.AutoMigrate(&Problem_linux{})
	return nil
}

func check_linux(id int, anser string) (Problem_linux, string, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return Problem_linux{}, "", fmt.Errorf("linux_check失敗: %w", err)
	}
	defer db.Close()
	var result string
	var linux Problem_linux
	if err := db.Where("id = ? AND anser = ?", id, anser).First(&linux).Error; err != nil {
		result = "不正解"
	} else {
		result = "正解"
	}
	return linux, result, nil
}

func linuxGetAll() ([]Problem_linux, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return nil, fmt.Errorf("dbGetAll失敗: %w", err)
	}
	defer db.Close()
	var linux []Problem_linux
	err = db.Order("created_at desc").Find(&linux).Error
	if err != nil {
		return nil, err
	}
	return linux, nil
}

func linuxGetOne(id int) (Problem_linux, error) {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return Problem_linux{}, fmt.Errorf("linuxGetone失敗: %w", err)
	}
	defer db.Close()
	var linux Problem_linux
	err = db.First(&linux, id).Error
	if err != nil {
		return Problem_linux{}, err
	}
	return linux, nil
}

func linuxInsert(question string, anser string, hint string) error {
	db, err := gorm.Open("mysql")
	if err != nil {
		return fmt.Errorf("linuxInsert失敗: %w", err)
	}
	defer db.Close()
	db.Create(&Problem_linux{Question: question, Anser: anser, Hint: hint})
	return nil
}

func linuxUpdate(id int, question string, hint string, anser string) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("linuxUpdate失敗: %w", err)
	}
	defer db.Close()
	var linux Problem_linux
	db.First(&linux, id)
	linux.Question = question
	linux.Anser = anser
	linux.Hint = hint
	db.Save(&linux)
	return nil
}

func linuxDelete(id int) error {
	db, err := gorm.Open("mysql", connectString)
	if err != nil {
		return fmt.Errorf("linuxDelete失敗: %w", err)
	}
	defer db.Close()
	var linux Problem_linux
	db.Where("id = ?", id).Delete(&linux)
	return nil
}
