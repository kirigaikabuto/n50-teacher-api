package groups

type UserGroupService interface {
	CreateGroup(cmd *CreateGroupCommand) (*Group, error)
	ListGroup(cmd *ListGroupCommand) ([]Group, error)
	GetGroupById(cmd *GetGroupByIdCommand) (*Group, error)
	CreateUserGroup(cmd *CreateUserGroupCommand) (*UserGroup, error)
	GetUserGroupByGroupId(cmd *GetUserGroupByGroupId) ([]UserGroup, error)
	GetUserGroupByUserId(cmd *GetUserGroupByUserId) ([]UserGroup, error)
	DeleteUserGroupById(cmd *DeleteUserGroupById) error
}

type userGroupService struct {
	userGroupStore UserGroupStore
}

func NewUserGroupService(userGroupStore UserGroupStore) UserGroupService {
	return &userGroupService{userGroupStore: userGroupStore}
}

func (u *userGroupService) CreateGroup(cmd *CreateGroupCommand) (*Group, error) {
	return u.userGroupStore.CreateGroup(&Group{
		Name: cmd.Name,
	})
}

func (u *userGroupService) ListGroup(cmd *ListGroupCommand) ([]Group, error) {
	return u.userGroupStore.ListGroup()
}

func (u *userGroupService) GetGroupById(cmd *GetGroupByIdCommand) (*Group, error) {
	return u.userGroupStore.GetGroupById(cmd.Id)
}

func (u *userGroupService) CreateUserGroup(cmd *CreateUserGroupCommand) (*UserGroup, error) {
	return u.userGroupStore.CreateUserGroup(&UserGroup{
		UserId:  cmd.UserId,
		GroupId: cmd.GroupId,
	})
}

func (u *userGroupService) GetUserGroupByGroupId(cmd *GetUserGroupByGroupId) ([]UserGroup, error) {
	return u.userGroupStore.GetUserGroupByGroupId(cmd.GroupId)
}

func (u *userGroupService) GetUserGroupByUserId(cmd *GetUserGroupByUserId) ([]UserGroup, error) {
	return u.userGroupStore.GetUserGroupByUserId(cmd.UserId)
}

func (u *userGroupService) DeleteUserGroupById(cmd *DeleteUserGroupById) error {
	return u.userGroupStore.RemoveUserGroupById(cmd.Id)
}
