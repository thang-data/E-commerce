package entity

import "time"

type Session struct {
	ID             string      `json:"sessionId,omitempty" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
	UserID         *string     `json:"userId,omitempty" gorm:"type:UUID;"` // userId/signerId/viewerId/adminId
	Email          string      `json:"email,omitempty" gorm:"size:320;"`
	Type           SessionType `json:"type,omitempty"`
	ExpirationDate time.Time   `json:"expirationDate,omitempty" gorm:"not null;"`

	// Extra information
	Is2FARequired bool   `json:"is2FARequired,omitempty" gorm:"-"`
	TOTP          *TOTP  `json:"totp,omitempty" gorm:"-"`
	Partner       string `json:"-"`
}

type TOTP struct {
	Secret string `json:"secret"`
	QRCode string `json:"qrCode"` // base64 image
}

const (
	SessionTypeAdmin = "admin"
)

type OnboardingStep string
type SessionType string
