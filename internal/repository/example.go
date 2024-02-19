package repository

import (
	"fmt"
	"github.com/urcop/go-fiber-template/internal/config"
	"github.com/urcop/go-fiber-template/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ExampleRepository interface {
	GetRow() ([]*models.Example, error)
}

type ExampleRepositoryImpl struct {
	db *gorm.DB
}

func (e *ExampleRepositoryImpl) GetRow() ([]*models.Example, error) {
	var example []*models.Example
	result := e.db.Find(&example)

	if result.Error != nil {
		return nil, result.Error
	}
	return example, nil

}

func NewExampleRepository() ExampleRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.Name, cfg.Db.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	pgSvc := &ExampleRepositoryImpl{db: db}
	err = db.AutoMigrate(&models.Example{})
	if err != nil {
		panic(err)
	}
	return pgSvc
}
