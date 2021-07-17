package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Database struct {
	postgres *gorm.DB
}

func (db *Database) Init() {
	var err error
	// database should be created in pgAdmin
	db.postgres, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=ie-project-db sslmode=disable password=62442")
	if err != nil {
		log.Println("[Database]: failed to connect database")
		log.Println(err)
		os.Exit(-1)
	}
	log.Println("[Database]: db is up")
	// the bellow commented code creates tables in database
	//db.createTables()
}

func (db *Database) createTables() {
	db.postgres.CreateTable(User{})
	db.postgres.CreateTable(Admin{})
	db.postgres.CreateTable(Product{})
	db.postgres.CreateTable(Category{})
	db.postgres.CreateTable(Receipt{})
}

func (db *Database) tmp() {
	db.postgres.Exec("")
}

// this three function is for admin category JOB3

// InsertCategory COMPLETE
func (db *Database) InsertCategory(categoryName string) {
	categories := make([]Category, 10)
	db.postgres.Find(&categories, "Name =?", categoryName)
	if (len(categories)) > 0 {
		print("there is another category with same name")
		return
	}
	categories = []Category{
		{Name: categoryName},
	}
	for _, categs := range categories {
		db.postgres.Create(&categs)
	}
}

// ModifyCategory COMPLETE
func (db *Database) ModifyCategory(newName string, oldName string) {
	db.postgres.Model(Category{}).Where("name = ?", oldName).Updates(Category{Name: newName})
	db.postgres.Model(Product{}).Where("category = ?", oldName).Updates(Product{Category: newName})
}

// GetCategories COMPLETE
func (db *Database) GetCategories() []string {
	allCategories := make([]Category, 20)
	db.postgres.Find(&allCategories)
	result := make([]string, 0)
	for _, cat := range allCategories {
		if cat.Name != "" {
			result = append(result, cat.Name)
		}
	}
	return result
}

// DeleteCategory COMPLETE
func (db *Database) DeleteCategory(categoryName string) {
	db.postgres.Model(Product{}).Where("category = ?", categoryName).Updates(Product{Category: "NO CATEGORY"})
	db.postgres.Where("name = ?", categoryName).Delete(&Category{})
}

// SeeAllReceipt COMPLETE
func (db *Database) SeeAllReceipt() []Receipt {
	receipts := make([]Receipt, 10)
	result := db.postgres.Find(&receipts)
	if result.Error != nil {
		panic(result.Error)
	}
	return receipts
}

// SeeReceiptByCode COMPLETE
func (db *Database) SeeReceiptByCode(code string) []Receipt {
	receipts := make([]Receipt, 10)
	result := db.postgres.Where("tracingCode=?", code).Find(&receipts)
	if result.Error != nil {
		panic(result.Error)
	}
	return receipts
}

// AddReceipt COMPLETE
func (db *Database) AddReceipt(productName string, soldNumber int, customerEmail string, customerFirstname string, customerLastname string, customerAddress string, amount int, date string, tracingCode string, status string) {
	receipts := []Receipt{
		{ProductName: productName, SoldNumber: soldNumber, CustomerAddress: customerAddress, CustomerEmail: customerEmail, CustomerFirstname: customerFirstname, CustomerLastname: customerLastname, Amount: amount, Date: date, TracingCode: tracingCode, Status: status},
	}
	for _, rescps := range receipts {
		db.postgres.Create(&rescps)
	}
}

// ChangeReceiptStatus COMPLETE
func (db *Database) ChangeReceiptStatus(code string, status string) {
	db.postgres.Model(Receipt{}).Where("tracingCode = ?", code).Updates(Receipt{Status: status})
}
