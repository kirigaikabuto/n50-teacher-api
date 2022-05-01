package subjects

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrSubjectNotFound      = com.NewMiddleError(errors.New("subject not found"), 404, 600)
	ErrSubjectIdNotProvided = com.NewMiddleError(errors.New("subject id is not provided"), 400, 601)
	ErrCreateSubjectUnknown = com.NewMiddleError(errors.New("could not create subject: unknown error"), 500, 602)

	ErrTeacherSubjectNotFound      = com.NewMiddleError(errors.New("teacher subject not found"), 404, 603)
	ErrTeacherSubjectIdNotProvided = com.NewMiddleError(errors.New("subject id is not provided"), 400, 604)
	ErrCreateTeacherSubjectUnknown = com.NewMiddleError(errors.New("could not create teacher subject: unknown error"), 500, 605)

	ErrUserIdNotProvided   = com.NewMiddleError(errors.New("user id is not provided"), 400, 606)
	ErrNoUserIdInToken     = com.NewMiddleError(errors.New("no user id in token"), 400, 607)
	ErrNoUserTypeInToken   = com.NewMiddleError(errors.New("no user type in token"), 400, 608)
	ErrNoAccessPermissions = com.NewMiddleError(errors.New("no access permissions"), 400, 609)

	ErrGroupSubjectNotFound      = com.NewMiddleError(errors.New("group subject not found"), 404, 610)
	ErrGroupSubjectIdNotProvided = com.NewMiddleError(errors.New("group subject id is not provided"), 400, 611)
	ErrCreateGroupSubjectUnknown = com.NewMiddleError(errors.New("could not create group subject: unknown error"), 500, 612)

	ErrInsertedUserIsNotTeacher = com.NewMiddleError(errors.New("inserted user is not teahcer"), 400, 613)
	ErrTeacherIdNotProvided     = com.NewMiddleError(errors.New("teacher id not provided"), 400, 614)
	ErrGroupIdNotProvided       = com.NewMiddleError(errors.New("group id not provided"), 400, 615)
)
