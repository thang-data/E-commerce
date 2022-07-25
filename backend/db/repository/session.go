package repository

import (
	"github.com/thang-data/backend/entity"
	"github.com/thang-data/backend/pkg"
	"gorm.io/gorm"
	"time"
)

func CreateSession(tx *gorm.DB, e *entity.Session) (*entity.Session, error) {
	if e.UserID != nil && *e.UserID == "" {
		e.UserID = nil
	}
	if e.ExpirationDate.IsZero() {
		e.ExpirationDate = time.Now().Add(pkg.SessionExpireAfter)
	}
	result := tx.Create(e)
	return e, result.Error
}
