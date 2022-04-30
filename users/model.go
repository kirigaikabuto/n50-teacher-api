package users

import "github.com/kirigaikabuto/n50-teacher-api/common"

type User struct {
	Id          string   `json:"id"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Email       string   `json:"email"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	CreatedDate string   `json:"created_date"`
	Type        common.UserType `json:"type"`
}

type UserUpdate struct {
	Id        string    `json:"id"`
	Username  *string   `json:"username"`
	Password  *string   `json:"password"`
	Email     *string   `json:"email"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Type      *common.UserType `json:"type"`
}

