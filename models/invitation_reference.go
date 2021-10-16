package models

type InvitationReference struct {
	ID              int `gorm:"primaryKey;autoIncrement"`
	MemberID        int `gorm:"not null"`
	InvitedMemberID int `gorm:"unique;not null"`

	Member        Member
	InvitedMember Member
}
