package entity

type EmailTemplate struct {
	ID                string `json:"id,omitempty" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
	EmailTypeID       string `json:"emailTypeId,omitempty" gorm:"type:UUID;not null;uniqueIndex:idx_email_template;"`
	Name              string `json:"name,omitempty" gorm:"not null;size:100;uniqueIndex:idx_email_template;"`
	Message           string `json:"message,omitempty" gorm:"not null;"`
	IsDefaultTemplate bool   `json:"isDefaultTemplate"`

	// Association
	//EmailType EmailType `json:"-"`

	// Extra information
	Vars []*EmailVar `json:"emailVars,omitempty" gorm:"-"`
}

type EmailVar struct {
	ID          string `json:"id,omitempty" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
	EmailTypeID string `json:"emailTypeId,omitempty" gorm:"type:UUID;not null;uniqueIndex:idx_email_var;"`
	Name        string `json:"name,omitempty" gorm:"not null;size:100;uniqueIndex:idx_email_var;"`
	Description string `json:"description,omitempty" gorm:"not null;size:100;"`

	// Association
	//EmailType EmailType `json:"-"`
}
