package groups

type UserGroupStore interface {
	CreateGroup(model *Group) (*Group, error)
	ListGroup() ([]Group, error)
	GetGroupById(id string) (*Group, error)
	CreateUserGroup(model *UserGroup) (*UserGroup, error)
	GetUserGroupByGroupId(groupId string) ([]UserGroup, error)
	GetUserGroupByUserId(userId string) ([]UserGroup, error)
	RemoveUserGroupById(id string) error
}
