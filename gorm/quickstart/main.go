package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&Product{})
	p := Product{Code: "D42", Price: 100}
	// Create
	result := db.Create(&p)
	if result.Error != nil {
		fmt.Println("err:", result.Error)
	}
	fmt.Printf("ID:%d, RowsAffected:%d\n", p.ID, result.RowsAffected)

	// Read
	var product Product
	result = db.First(&product, "id=?", 2) // 根据整型主键查找
	if result.Error != nil {
		fmt.Println("id:1->", result.Error)
		fmt.Println("id:1->", errors.Is(result.Error, gorm.ErrRecordNotFound))
	}

	fmt.Printf("id:1->%v\n", product)

	result = db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	if result.Error != nil {
		fmt.Println("code = D42->", result.Error)
	}

	fmt.Printf("code = D42->%v\n", product)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	db.Delete(&product, 1)
}
