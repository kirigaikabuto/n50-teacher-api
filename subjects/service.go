package subjects

import (
	"github.com/kirigaikabuto/n50-teacher-api/groups"
	"github.com/kirigaikabuto/n50-teacher-api/users"
)

type SubjectService interface {

}

type subjectService struct {
	subjectStore SubjectStore
	userStore    users.UsersStore
	groupStore   groups.UserGroupStore
}

func NewSubjectService(s SubjectStore, u users.UsersStore, g groups.UserGroupStore) SubjectService {
	return subjectService{
		subjectStore: s,
		userStore:    u,
		groupStore:   g,
	}
}
