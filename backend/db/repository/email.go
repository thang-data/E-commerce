package repository

import (
	"github.com/thang-data/backend/entity"
	"gorm.io/gorm"
)

func GetDefaultEmailTemplate(tx *gorm.DB, emailTypeId string) (*entity.EmailTemplate, error) {
	var item entity.EmailTemplate
	result := tx.First(&item, "email_type_id = ? AND is_default_template", emailTypeId)

	return &item, result.Error
}
