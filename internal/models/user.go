package models

import "time"

type User struct{
	ID           uint       `gorm:"primary_key" json:"id"`
	Name  string     `gorm:"type:varchar(16);NOT NULL" json:"name"`
	Password string  `gorm:"type:varchar(64);NOT NULL" json:"password"`
	Phone     uint64      `gorm:"type:bigint unsigned;NOT NULL" json:"phone"`
	Email string `gorm:"type:varchar(64);NOT NULL" json:"email"`

	CreatedAt    time.Time  `gorm:"type:timestamp;DEFAULT:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;DEFAULT:current_timestamp ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index"`
}
