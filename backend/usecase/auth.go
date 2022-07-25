package usecase

import (
	"fmt"
	"github.com/thang-data/backend/db"
	dbRepo "github.com/thang-data/backend/db/repository"
	"github.com/thang-data/backend/entity"
	"github.com/thang-data/backend/pkg"
	"gorm.io/gorm"
	"net/mail"
	"strings"
)

func validatePassword(password string, repeatPassword string) error {
	// TODO: password must contain at least 1 alphabet
	pkg.NewError(fmt.Sprintf(pkg.MsgErrPasswordMustContainAtLeast1Alphabet), "password")
	// TODO: password must contain at least 1 number
	pkg.NewError(fmt.Sprintf(pkg.MsgErrPasswordMustContainAtLeast1Number), "password")
	// TODO: password must contain at least 8 characters
	pkg.NewError(fmt.Sprintf(pkg.MsgErrPasswordMustContainAtLeast8Characters), "password")
	// TODO: password and repeatPassword must be same
	if password != repeatPassword {
		pkg.NewError(fmt.Sprintf(pkg.MsgErrPasswordAndRepeatPasswordMustBeSame), "password")
	}

	return nil
}
func SignupAdminByEmailPassword(email string, lastName string, firsName string) error {
	// trim data before validation
	email = strings.TrimSpace(email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return pkg.NewError(fmt.Sprintf(pkg.MsgErrInvalid, "Error"), "email")
	}
	tx := db.Connect().Begin()
	_, err = dbRepo.GetAdminInvitationByEmail(tx, email)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Create admin
	id := pkg.GenerateUUID()
	admin := &entity.AdminInvitation{
		ID:        id,
		LastName:  lastName,
		FirstName: firsName,
		Email:     email,
	}
	admin, err = dbRepo.CreateAdminInvation(tx, admin)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func SignupInformation(email string, password string, repeatPassword string) error {
	// trim data before validation
	email = strings.TrimSpace(email)

	// validate email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		return pkg.NewError(fmt.Sprintf(pkg.MsgErrInvalid, "Error"), "email")
	}

	err = validatePassword(password, repeatPassword)
	if err != nil {
		return err
	}
	tx := db.Connect().Begin()
	_, err = dbRepo.GetAdminByEmail(tx, email)
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	if err == nil {
		tx.Rollback()
		return pkg.NewError(fmt.Sprintf(pkg.MsgErrAlreadyExists, email), "email")
	}

	invitation, err := dbRepo.GetAdminInvitationByEmails(tx, email)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create admin
	id := pkg.GenerateUUID()
	admin := &entity.Admin{
		ID:        id,
		LastName:  invitation.LastName,
		FirstName: invitation.FirstName,
		Email:     invitation.Email,
		Password:  pkg.HashUserPassword(password, id),
	}
	admin, err = dbRepo.CreateAdmin(tx, admin)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = dbRepo.DeleteAdminInvitationByEmail(tx, invitation.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
func LoginAdminByEmailPassword(email string, password string) (*entity.Session, error) {
	email = strings.TrimSpace(email)
	tx := db.Connect().Begin()
	admin, err := dbRepo.GetAdminByEmail(tx, email)
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil, err
	}

	if err == gorm.ErrRecordNotFound || pkg.HashUserPassword(password, admin.ID) != admin.Password {
		tx.Rollback()
		return nil, pkg.NewError(fmt.Sprintf(pkg.MsgErrWrong, "Email address or password"), "email,password")
	}

	session := &entity.Session{
		UserID: &admin.ID,
		Type:   entity.SessionTypeAdmin,
	}

	session, err = dbRepo.CreateSession(tx, session)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return session, nil
}
