package models

import (
	"time"
)

type Users struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	FullName   string    `gorm:"type:varchar(300)" json:"full_name"`
	Email      string    `gorm:"type:varchar(300)" json:"email"`
	Age        int       `gorm:"type:int" json:"age"`
	Religion   string    `gorm:"type:varchar(300)" json:"religion"`
	Gender     string    `gorm:"type:varchar(300)" json:"gender"`
	Username   string    `gorm:"type:varchar(300)" json:"username"`
	Password   string    `gorm:"type:varchar(300)" json:"password"`
	Photo      string    `gorm:"type:varchar(300)" json:"photo"`
	LastLogin  time.Time `gorm:"type:timestamp" json:"last_login"`
	IsVerified bool      `gorm:"type:bool" json:"is_verified"`
	Deleted    bool      `gorm:"type:bool" json:"deleted"`
	DeletedAt  time.Time `gorm:"type:timestamp" json:"deleted_at"`
}
