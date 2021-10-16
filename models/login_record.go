package models

import "time"

type LoginRecord struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	IP        string `gorm:"size:64"`
	AdminID   *int
	MemberID  *int
	CreatedAt time.Time `gorm:"type:timestamp;not null"`

	Admin  *Admin
	Member *Member
}
