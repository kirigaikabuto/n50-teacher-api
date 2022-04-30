package users

import (
	"github.com/google/uuid"
	"github.com/kirigaikabuto/n50-teacher-api/auth"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type UserService interface {
	CreateUser(cmd *CreateUserCommand) (*User, error)
	UpdateUser(cmd *UpdateUserCommand) (*User, error)
	DeleteUser(cmd *DeleteUserCommand) error
	GetUser(cmd *GetUserCommand) (*User, error)
	ListUser(cmd *ListUserCommand) ([]User, error)
	GetUserByUsernameAndPassword(cmd *GetUserByUsernameAndPassword) (*User, error)
	Login(cmd *LoginCommand) (*auth.TokenDetails, error)
}

type userService struct {
	userStore  UsersStore
	tokenStore auth.TokenStore
}

func NewUserService(uStore UsersStore, aStore auth.TokenStore) UserService {
	return &userService{userStore: uStore, tokenStore: aStore}
}

func (u *userService) CreateUser(cmd *CreateUserCommand) (*User, error) {
	_, err := u.userStore.GetByUsernameAndPassword(cmd.Username, cmd.Password)
	if err != ErrUserNotFound {
		return nil, ErrUserWithUsernameAlreadyExist
	}
	user := &User{
		Id:        uuid.New().String(),
		Username:  cmd.Username,
		Password:  cmd.Password,
		Email:     cmd.Email,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Type:      ToUserType(cmd.Type),
	}
	return u.userStore.Create(user)
}

func (u *userService) UpdateUser(cmd *UpdateUserCommand) (*User, error) {
	if cmd.Id == "" {
		return nil, ErrUserIdNotProvided
	}
	oldUser, err := u.GetUser(&GetUserCommand{
		Id: cmd.Id,
	})
	if err != nil {
		return nil, err
	}
	userUpdate := &UserUpdate{Id: cmd.Id}
	if cmd.Password != "" {
		hashedPassword, err := setdata_common.HashPassword(cmd.Password)
		if err != nil {
			return nil, err
		}
		if oldUser.Password != hashedPassword {
			userUpdate.Password = &hashedPassword
		}
	}
	if cmd.Username != "" && cmd.Username != oldUser.Username {
		userUpdate.Username = &cmd.Username
	}
	if cmd.Email != "" && cmd.Email != oldUser.Email {
		userUpdate.Email = &cmd.Email
	}
	return u.userStore.Update(userUpdate)
}

func (u *userService) DeleteUser(cmd *DeleteUserCommand) error {
	return u.userStore.Delete(cmd.Id)
}

func (u *userService) GetUser(cmd *GetUserCommand) (*User, error) {
	return u.userStore.Get(cmd.Id)
}

func (u *userService) ListUser(cmd *ListUserCommand) ([]User, error) {
	return u.userStore.List()
}

func (u *userService) GetUserByUsernameAndPassword(cmd *GetUserByUsernameAndPassword) (*User, error) {
	return u.userStore.GetByUsernameAndPassword(cmd.Username, cmd.Password)
}

func (u *userService) Login(cmd *LoginCommand) (*auth.TokenDetails, error) {
	user, err := u.GetUserByUsernameAndPassword(&GetUserByUsernameAndPassword{
		Username: cmd.Username,
		Password: cmd.Password,
	})
	if err != nil {
		return nil, err
	}
	token, err := u.tokenStore.CreateToken(&auth.CreateTokenCommand{
		UserId:   user.Id,
		UserType: user.Type.ToString(),
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
