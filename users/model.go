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

var (
	Admin   UserType = "админ"
	Teacher UserType = "учитель"
	Student UserType = "ученик"
)

var (
	userTypeToString = map[UserType]string{
		Admin:   "админ",
		Teacher: "учитель",
		Student: "ученик",
	}
	stringToUserType = map[string]UserType{
		"админ":   Admin,
		"учитель": Teacher,
		"ученик":  Student,
	}
)

func (u UserType) ToString() string {
	return userTypeToString[u]
}

func ToUserType(s string) UserType {
	return stringToUserType[s]
}

func IsUserTypeExist(s string) bool {
	userTypes := []string{"админ", "учитель", "ученик"}
	for _, v := range userTypes {
		if v == s {
			return true
		}
	}
	return false
}
