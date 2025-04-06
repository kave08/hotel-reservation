package types

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 12

type User struct {
	ID               string `bson:"_id" json:"id"`
	FirstName        string `bson:"first_name" json:"first_name "`
	LastName         string `bson:"last_name" json:"last_name "`
	Email            string `bson:"email" json:"email "`
	EncryptedPasword string `bson:"encrypted_pasword" json:"-"`
}

type UserRequest struct {
	FirstName string `json:"first_name "`
	LastName  string `json:"last_name "`
	Email     string `bson:"email" json:"email "`
	Pasword   string `json:"pasword "`
}

func NewUserParams(params UserRequest) (*User, error) {
	encpsw, err := bcrypt.GenerateFromPassword([]byte(params.Pasword), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:        params.FirstName,
		LastName:         params.LastName,
		Email:            params.Email,
		EncryptedPasword: string(encpsw),
	}, nil
}
