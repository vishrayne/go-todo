package v1

import (
	"errors"

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

func (db *DataBase) autoMigrate() error {
	if db == nil || !db.isConnected {
		return errors.New("database not open or initialized")
	}

	db.gormDb.AutoMigrate(&todoModel{})
	return nil
}

func (db *DataBase) save(model interface{}) error {
	if db == nil || !db.isConnected {
		return errors.New("database not open or initialized")
	}

	db.gormDb.Save(model)
	return nil
}

func (db *DataBase) find(model interface{}) error {
	if db == nil || !db.isConnected {
		return errors.New("database not open or initialized")
	}

	db.gormDb.Find(model)
	return nil
}

func (db *DataBase) findBy(model interface{}, where ...interface{}) error {
	if db == nil || !db.isConnected {
		return errors.New("database not open or initialized")
	}

	db.gormDb.Find(model, where)
	return nil
}

func (db *DataBase) updateAttributes(model interface{}, attrs map[string]interface{}) error {
	if db == nil || !db.isConnected {
		return errors.New("database not open or initialized")
	}

	if len(attrs) <= 0 {
		return errors.New("nothing to update")
	}

	for key, value := range attrs {
		db.gormDb.Model(model).Update(key, value)
	}

	return nil
}

func (db *DataBase) delete(model interface{}) error {
	if db == nil || !db.isConnected {
		return errors.New("database not open or initialized")
	}

	db.gormDb.Delete(model)

	return nil
}
