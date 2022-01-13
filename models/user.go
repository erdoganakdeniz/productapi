package models

import "productapi/utils"

type User struct {
	ID        string               `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string               `json:"email,omitempty" bson:"email"`
	UserName  string               `json:"username,omitempty" bson:"username"`
	Password  string               `json:"-" bson:"password,omitempty"`
	Products   Product `json:"products,omitempty" bson:"products"`
}
type SignupInputs struct {
	Email    string `json:"email,omitempty" bson:"email" valid:"email"`
	UserName string `json:"user_name,omitempty" bson:"user_name" valid:"length(3|30)"`
	Password string `json:"password,omitempty" bson:"password" valid:"length(6|30)"`
}
type LoginInputs struct {
	Email    string `json:"email,omitempty" bson:"email" valid:"email"`
	Password string `json:"password,omitempty" bson:"password" valid:"length(6|30)"`
}

func (i SignupInputs) Validate() error {
	return utils.Validator(i)
}
