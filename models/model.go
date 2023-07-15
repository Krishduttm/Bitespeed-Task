package models

import (
	"time"
)

type Link string

const (
	Primary   Link = "primary"
	Secondary Link = "secondary"
)

type Contact struct {
	Id             int       `gorm:"primaryKey, auto_increment"`
	PhoneNumber    string    `gorm:"type:varchar(10)"`
	Email          string    `gorm:"type:varchar(50)"`
	LinkedIn       int       `gorm:"type:int"`
	LinkPrecedence Link      // "primary" if it's the first Contact in the link
	CreatedAt      time.Time `gorm:"type:timestamp"`
	UpdatedAt      time.Time `gorm:"type:timestamp"`
	// DeletedAt      time.Time `gorm:"type:timestamp"`
}
