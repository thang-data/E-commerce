package entity

type AdminInvitation struct {
	ID        string `json:"id,omitempty" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
	Email     string `json:"email,omitempty" gorm:"size:320;not null;unique;"`
	LastName  string `json:"lastName,omitempty" gorm:"size:100;not null;"`
	FirstName string `json:"firstName,omitempty" gorm:"size:100;not null;"`

	// Extra information
	CountPermission int `json:"countPermission,omitempty" gorm:"-"`
}
