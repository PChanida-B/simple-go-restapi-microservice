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

/*
func NewSqliteDBStore(dsn string) *GormStore {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&service.Request{})

	return &GormStore{db: db}
}*/

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) Create(value interface{}) error {
	return s.db.Create(value).Error
}

func (s *GormStore) Read(value interface{}) error {
	return s.db.Find(value).Error
}

/*Read with primary key*/
func (s *GormStore) ReadIndex(value interface{}, id int) error {
	return s.db.Find(value, id).Error
}

/*Update with primary key*/
func (s *GormStore) Update(value interface{}, id int) error {
	return s.db.Model(value).Where("id = ?", id).Updates(value).Error
}

/*Delete with primary key*/
func (s *GormStore) Delete(value interface{}, id int) error {
	return s.db.Delete(value, id).Error
}
