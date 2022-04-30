package groups

type CreateGroupCommand struct {
	Name            string `json:"name"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *CreateGroupCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserGroupService).CreateGroup(cmd)
}

type ListGroupCommand struct {
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *ListGroupCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserGroupService).ListGroup(cmd)
}

type GetGroupByIdCommand struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetGroupByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserGroupService).GetGroupById(cmd)
}

type CreateUserGroupCommand struct {
	UserId          string `json:"user_id"`
	GroupId         string `json:"group_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *CreateUserGroupCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserGroupService).CreateUserGroup(cmd)
}

type GetUserGroupByGroupId struct {
	GroupId         string `json:"group_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetUserGroupByGroupId) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserGroupService).GetUserGroupByGroupId(cmd)
}

type GetUserGroupByUserId struct {
	UserId          string `json:"user_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetUserGroupByUserId) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserGroupService).GetUserGroupByUserId(cmd)
}

type DeleteUserGroupById struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *DeleteUserGroupById) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(UserGroupService).DeleteUserGroupById(cmd)
}
