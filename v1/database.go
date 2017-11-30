package v1

import (
	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DataBase handler
type DataBase struct {
	gormDb      *gorm.DB
	isConnected bool
}

type todoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func newDatabase() *DataBase {
	return &DataBase{}
}

func (db *DataBase) open() error {
	database, err := gorm.Open("sqlite3", "data/gorm.db")
	if err != nil {
		return err
	}

	db.gormDb = database
	db.isConnected = true
	return nil
}

func (db *DataBase) close() {
	if db == nil {
		return
	}

	db.isConnected = false
	db.gormDb.Close()
}

func (db *DataBase) autoMigrate() {
	if db == nil || !db.isConnected {
		panic("database not open or initialized")
	}

	db.gormDb.AutoMigrate(&todoModel{})
}

func (db *DataBase) save(value interface{}) {
	if db == nil || !db.isConnected {
		panic("database not open or initialized")
	}

	db.gormDb.Save(value)
}

func (db *DataBase) find(value interface{}) {
	if db == nil || !db.isConnected {
		panic("database not open or initialized")
	}

	db.gormDb.Find(value)
}

func (db *DataBase) findBy(value interface{}, where ...interface{}) {
	if db == nil || !db.isConnected {
		panic("database not open or initialized")
	}

	db.gormDb.Find(value, where)
}
