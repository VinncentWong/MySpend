package infrastructure

import (
	"fmt"
	"module/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() error {
	dsn := fmt.Sprintf(`
	host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"),
		os.Getenv("password"), os.Getenv("dbname"))
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = _db
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func Migrate() error {
	err := db.AutoMigrate(
		domain.User{},
		domain.BudgetFivety{},
		domain.BudgetThirty{},
		domain.BudgetTwenty{},
		domain.FinancialAccount{},
		domain.PaymentHistory{},
	)
	return err
}
