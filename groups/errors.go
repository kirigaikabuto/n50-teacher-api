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
	ErrUserIdNotProvided      = com.NewMiddleError(errors.New("user id is not provided"), 400, 503)
	ErrNoUserIdInToken        = com.NewMiddleError(errors.New("no user id in token"), 400, 504)
	ErrNoUserTypeInToken      = com.NewMiddleError(errors.New("no user type in token"), 400, 505)
	ErrNoAccessPermissions    = com.NewMiddleError(errors.New("no access permissions"), 400, 506)
)
