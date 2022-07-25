package entity

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID            string `json:"id,omitempty" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
	Email         string `json:"email,omitempty" form:"email" gorm:"size:365;not null;unique;"`
	Password      string `json:"-" form:"password" gorm:"size:100;"`
	LastName      string `json:"lastName,omitempty" form:"lastName" gorm:"size:100;not null;"`   // required when input profile
	FirstName     string `json:"firstName,omitempty" form:"firstName" gorm:"size:100;not null;"` // required when input profile
	EmailVerified bool   `json:"emailVerified,omitempty"`

	CreatedAt *time.Time     `json:"createdAt,omitempty" gorm:"<-:create"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Extra information
	OldPassword                 string `json:"-" form:"oldPassword" gorm:"-"`
	ResendEmailVerificationCode bool   `json:"-" gorm:"-"`
	CountPermission             int    `json:"countPermission,omitempty" gorm:"-"`
}

type AdminProfile struct {
	Admin
	EmailPending string `json:"emailPending"`
}
