package users

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Type      *UserType `json:"type"`
}

type UserUpdate struct {
	Id        string    `json:"id"`
	Username  *string   `json:"username"`
	Password  *string   `json:"password"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Type      *UserType `json:"type"`
}

type UserType string


