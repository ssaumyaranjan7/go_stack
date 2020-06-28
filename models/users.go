package models

import (
	"fmt"

	"github.com/Kamva/mgm"
	"github.com/ssaumyaranjan7/go_stack/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	mgm.DefaultModel `bson:"inline"`
	Basics           Basics `json:"basics" bson:"basics"`
	Role             string `json:"role" bson:"role"`
	Address          string `json:"address" bson:"address"`
	Location         struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
	}
}

type Basics struct {
	FirstName       string `json:"first_name" bson:"firstName"`
	LastName        string `json:"last_name" bson:"lastName"`
	Email           string `json:"email" bson:"email"`
	Password        string `json:"password" bson:"password"`
	MobileNumber    string `json:"mobile_number" bson:"mobileNumber"`
	IsActive        bool   `json:"is_active" bson:"isActive"`
	ProfileImageURL string `json:"profile_image_url" bson:"profileImageURL"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// Register is used to insert any user into db
func Register(user *User) bool {
	err := mgm.Coll(user).Create(user)
	if err != nil {
		return false
	}
	return true
}

// Creating is used as a pre-insert hook
func (model *User) Creating() error {
	// Call to DefaultModel Creating hook
	if err := model.DefaultModel.Creating(); err != nil {
		return err
	}
	fmt.Println("control Here")
	// We can check if model fields is not valid, return error to
	// cancel document insertion .
	pwd := []byte(model.Basics.Password)
	encPwd := helpers.PasswordEncrypt(pwd)

	fmt.Println(encPwd)

	model.Basics.Password = string(encPwd)
	return nil
}

// FindUserByEmail is used to retrive user details from DB by Email
func FindUserByEmail(Email string) *User {
	u := &User{}
	coll := mgm.Coll(u)
	if err := coll.First(bson.M{"basics.email": Email}, u); err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}

// FindUserByID us used to retrive suer details from DB by ID
func FindUserByID(id string) *User {
	u := &User{}
	coll := mgm.Coll(u)
	if err := coll.FindByID(id, u); err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}

func UpdateUser(u *User) *User {
	if err := mgm.Coll(u).Update(u); err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}
