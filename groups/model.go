package groups

type Group struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CreatedDate string `json:"created_date"`
}

type UserGroup struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	GroupId     string `json:"group_id"`
	CreatedDate string `json:"created_date"`
}

type UpdateGroup struct {
	Id   string  `json:"id"`
	Name *string `json:"name"`
}

type UpdateUserGroup struct {
	Id      string  `json:"id"`
	UserId  *string `json:"user_id"`
	GroupId *string `json:"group_id"`
}
