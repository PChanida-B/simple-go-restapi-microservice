package store

import (
	"github.com/PChanida-B/simple-go-restapi-microservice/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewMariaDBStore(dsn string) *GormStore {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&service.Request{})

	return &GormStore{db: db}
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) Create(req service.Request) error {
	return s.db.Create(req).Error
}

func (s *GormStore) Read(req service.Request) error {
	return nil
}

func (s *GormStore) Update(req service.Request) error {
	return nil
}

func (s *GormStore) Delete(req service.Request) error {
	return nil
}
