package groups

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrGroupNotFound          = com.NewMiddleError(errors.New("group not found"), 404, 400)
	ErrGroupIdNotProvided     = com.NewMiddleError(errors.New("group id is not provided"), 400, 401)
	ErrCreateGroupUnknown     = com.NewMiddleError(errors.New("could not create group: unknown error"), 500, 402)
	ErrUserGroupNotFound      = com.NewMiddleError(errors.New("user group not found"), 404, 500)
	ErrUserGroupIdNotProvided = com.NewMiddleError(errors.New("user group id is not provided"), 400, 501)
	ErrCreateUserGroupUnknown = com.NewMiddleError(errors.New("could not create user group: unknown error"), 500, 502)
)
