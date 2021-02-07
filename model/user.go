package model

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//User 用户信息
type User struct {
	gorm.Model
	Username      string    `form:"username" gorm:"type:varchar(255);unique_index" binding:"required" `
	Password      string    `form:"password" gorm:"type:varchar(255)" binding:"required" `
	Role          string    `form:"-" gorm:"type:varchar(255)"`
	Enable        bool      `form:"-" gorm:"DEFAULT:true"`
	HasAllChannel bool      `form:"-" `
	Creator       string    `form:"-" gorm:"type:varchar(255)"`
	LastLoginAt   time.Time `form:"-" `
	URLTokenOnly  bool      `form:"url_token_only" gorm:"-"`
}

//TableName
func (User) TableName() string {
	return "t_users"
}

// GetPasswordHash 获取密码hash
func (u *User) GetPasswordHash() string {
	h := md5.New()
	h.Write([]byte(u.Password))
	return hex.EncodeToString(h.Sum(nil))
}

// Verify 验证密码
func (u *User) Verify() bool {
	GRWLock.Lock()
	db, err := GetDbHandler()
	if err != nil {
		GRWLock.Unlock()
		return false
	}

	defer func() {
		db.Close()
		GRWLock.Unlock()
	}()
	var user User
	db.Model(&User{}).Where("username =?", u.Username).First(&user)
	return strings.ToLower(user.Password) == strings.ToLower(u.Password)
	// fmt.Println(u.GetPasswordHash())
	// fmt.Println(u.Password)
	// return strings.ToLower(u.GetPasswordHash()) == strings.ToLower(u.Password)
}
