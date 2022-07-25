package repository

import (
	"github.com/thang-data/backend/entity"
	"gorm.io/gorm"
)

func GetAdminByEmail(tx *gorm.DB, email string) (*entity.Admin, error) {
	var item entity.Admin
	result := tx.First(&item, "email = ?", email)

	return &item, result.Error
}

func GetAdminInvitationByEmail(tx *gorm.DB, email string) (*entity.AdminInvitation, error) {
	var item entity.AdminInvitation
	result := tx.Not(&item, "email = ?", email)

	return &item, result.Error
}
func GetAdminInvitationByEmails(tx *gorm.DB, email string) (*entity.AdminInvitation, error) {
	var item entity.AdminInvitation
	result := tx.First(&item, "email = ?", email)

	return &item, result.Error
}
func CreateAdmin(tx *gorm.DB, e *entity.Admin) (*entity.Admin, error) {
	result := tx.Create(e)
	return e, result.Error
}
