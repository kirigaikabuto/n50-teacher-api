package lessons

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrLessonNotFound            = com.NewMiddleError(errors.New("lesson not found"), 404, 700)
	ErrLessonIdNotProvided       = com.NewMiddleError(errors.New("lesson id is not provided"), 400, 701)
	ErrCreateLessonUnknown       = com.NewMiddleError(errors.New("could not create lesson: unknown error"), 500, 702)
	ErrNothingToUpdate           = com.NewMiddleError(errors.New("nothing to update"), 400, 703)
	ErrNoAccessPermissions       = com.NewMiddleError(errors.New("no access permissions"), 400, 704)
	ErrNoUserIdInToken           = com.NewMiddleError(errors.New("no user id in token"), 400, 705)
	ErrNoUserTypeInToken         = com.NewMiddleError(errors.New("no user type in token"), 400, 706)
	ErrGroupSubjectIdNotProvided = com.NewMiddleError(errors.New("group subject id is not provided"), 400, 707)
)
