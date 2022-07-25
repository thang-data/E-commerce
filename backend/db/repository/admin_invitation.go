package repository

import (
	"github.com/thang-data/backend/entity"
	"gorm.io/gorm"
)

func DeleteAdminInvitationByEmail(tx *gorm.DB, email string) error {
	result := tx.Where("email = ?", email).Delete(&entity.AdminInvitation{})
	return result.Error
}
func CreateAdminInvation(tx *gorm.DB, e *entity.AdminInvitation) (*entity.AdminInvitation, error) {
	result := tx.Create(e)
	return e, result.Error
}
