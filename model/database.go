package model

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // justifying
)

const (
	//DbPath 数据库路径
	dbPath = "gb28181.db"
)

var (
	// GRWLock 数据库全局读写锁
	GRWLock *sync.RWMutex
)

func init() {
	GRWLock = new(sync.RWMutex)
	GRWLock.Lock()
	db, err := GetDbHandler()
	if err != nil {
		GRWLock.Unlock()
		os.Exit(-1)
	}

	defer func() {
		db.Close()
		GRWLock.Unlock()
	}()

	db.AutoMigrate(&Device{})
	db.AutoMigrate(&User{})

	user := &User{
		Model: gorm.Model{
			ID: 1,
		},
		Username:      "admin",
		Password:      "21232f297a57a5a743894a0e4a801fc3",
		Role:          "超级管理员",
		Enable:        true,
		HasAllChannel: true,
	}

	db.Model(&User{}).Save(user)
}

// GetDbHandler 获取数据库句柄
func GetDbHandler() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
