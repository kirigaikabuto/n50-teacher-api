package subjects

import (
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"github.com/kirigaikabuto/n50-teacher-api/groups"
	"github.com/kirigaikabuto/n50-teacher-api/users"
)

type SubjectService interface {
	CreateSubject(cmd *CreateSubjectCommand) (*Subject, error)
	ListSubjects(cmd *ListSubjectsCommand) ([]Subject, error)
	GetSubjectById(cmd *GetSubjectByIdCommand) (*Subject, error)
	GetSubjectsByGroupId(cmd *GetSubjectsByGroupId) ([]SubjectFullInfo, error)

	CreateTeacherSubject(cmd *CreateTeacherSubjectCommand) (*TeacherSubject, error)
	ListTeacherSubjects(cmd *ListTeacherSubjectsCommand) ([]TeacherSubject, error)
	GetTeacherSubjectById(cmd *GetTeacherSubjectByIdCommand) (*TeacherSubject, error)
	GetTeacherSubjectsByTeacherId(cmd *GetTeacherSubjectsByTeacherIdCommand) ([]TeacherSubject, error)
	GetTeacherSubjectsBySubjectId(cmd *GetTeacherSubjectsBySubjectIdCommand) ([]TeacherSubject, error)
	GetTeacherSubjectByToken(cmd *GetTeacherSubjectsByTokenCommand) ([]TeacherSubject, error)

	CreateGroupSubject(cmd *CreateGroupSubjectCommand) (*GroupSubject, error)
	ListGroupSubjects(cmd *ListGroupSubjectsCommand) ([]GroupSubject, error)
	GetGroupSubjectsById(cmd *GetGroupSubjectById) (*GroupSubject, error)
	GetGroupSubjectByIdTeacherSub(cmd *GetGroupSubjectByIdTeacherSub) ([]GroupSubject, error)
	GetGroupSubjectByGroupId(cmd *GetGroupSubjectByGroupId) ([]GroupSubject, error)
	GetGroupSubjectByTeacherGroupIds(cmd *GetGroupSubjectByTeacherGroupIdsCommand) (*GroupSubject, error)
}

type subjectService struct {
	subjectStore SubjectStore
	userStore    users.UsersStore
	groupStore   groups.UserGroupStore
}

func NewSubjectService(s SubjectStore, u users.UsersStore, g groups.UserGroupStore) SubjectService {
	return &subjectService{
		subjectStore: s,
		userStore:    u,
		groupStore:   g,
	}
}

func (s *subjectService) CreateSubject(cmd *CreateSubjectCommand) (*Subject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	model := &Subject{
		Name:        cmd.Name,
		Description: cmd.Description,
	}
	return s.subjectStore.CreateSubject(model)
}

func (s *subjectService) ListSubjects(cmd *ListSubjectsCommand) ([]Subject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.ListSubjects()
}

func (s *subjectService) GetSubjectById(cmd *GetSubjectByIdCommand) (*Subject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetSubjectById(cmd.Id)
}

func (s *subjectService) GetSubjectsByGroupId(cmd *GetSubjectsByGroupId) ([]SubjectFullInfo, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Student.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	_, err := s.groupStore.GetGroupById(cmd.GroupId)
	if err != nil {
		return nil, err
	}
	groupSubjects, err := s.subjectStore.GetGroupSubjectByGroupId(cmd.GroupId)
	if err != nil {
		return nil, err
	}
	resp := []SubjectFullInfo{}
	for _, v := range groupSubjects {
		teacherSub, err := s.subjectStore.GetTeacherSubjectById(v.TeacherSubjectId)
		if err != nil {
			return nil, err
		}
		sub, err := s.subjectStore.GetSubjectById(teacherSub.SubjectId)
		if err != nil {
			return nil, err
		}
		groupSubject, err := s.subjectStore.GetGroupSubjectByTeacherGroupIds(teacherSub.Id, cmd.GroupId)
		if err != nil {
			return nil, err
		}
		temp := SubjectFullInfo{
			Subject:        *sub,
			TeacherSubject: *teacherSub,
			GroupSubject:   *groupSubject,
		}
		resp = append(resp, temp)
	}
	return resp, nil
}

func (s *subjectService) CreateTeacherSubject(cmd *CreateTeacherSubjectCommand) (*TeacherSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	if cmd.TeacherId == "" {
		if cmd.CurrentUserType == common.Teacher.ToString() {
			cmd.TeacherId = cmd.CurrentUserId
		} else if cmd.CurrentUserType == common.Admin.ToString() {
			return nil, ErrNoTeacherId
		}
	}
	teacher, err := s.userStore.Get(cmd.TeacherId)
	if err != nil {
		return nil, err
	}
	if teacher.Type != common.Teacher {
		return nil, ErrInsertedUserIsNotTeacher
	}
	_, err = s.subjectStore.GetSubjectById(cmd.SubjectId)
	if err != nil {
		return nil, err
	}
	return s.subjectStore.CreateTeacherSubject(&TeacherSubject{
		TeacherId: cmd.TeacherId,
		SubjectId: cmd.SubjectId,
	})
}

func (s *subjectService) ListTeacherSubjects(cmd *ListTeacherSubjectsCommand) ([]TeacherSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.ListTeacherSubjects()
}

func (s *subjectService) GetTeacherSubjectById(cmd *GetTeacherSubjectByIdCommand) (*TeacherSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetTeacherSubjectById(cmd.Id)
}

func (s *subjectService) GetTeacherSubjectsByTeacherId(cmd *GetTeacherSubjectsByTeacherIdCommand) ([]TeacherSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetTeacherSubjectsByTeacherId(cmd.TeacherId)
}

func (s *subjectService) GetTeacherSubjectsBySubjectId(cmd *GetTeacherSubjectsBySubjectIdCommand) ([]TeacherSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetTeacherSubjectsBySubjectId(cmd.SubjectId)
}

func (s *subjectService) CreateGroupSubject(cmd *CreateGroupSubjectCommand) (*GroupSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	_, err := s.subjectStore.GetTeacherSubjectById(cmd.TeacherSubjectId)
	if err != nil {
		return nil, err
	}
	_, err = s.groupStore.GetGroupById(cmd.GroupId)
	if err != nil {
		return nil, err
	}
	return s.subjectStore.CreateGroupSubject(&GroupSubject{
		GroupId:          cmd.GroupId,
		TeacherSubjectId: cmd.TeacherSubjectId,
	})
}

func (s *subjectService) ListGroupSubjects(cmd *ListGroupSubjectsCommand) ([]GroupSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.ListGroupSubjects()
}

func (s *subjectService) GetGroupSubjectsById(cmd *GetGroupSubjectById) (*GroupSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetGroupSubjectsById(cmd.Id)
}

func (s *subjectService) GetGroupSubjectByIdTeacherSub(cmd *GetGroupSubjectByIdTeacherSub) ([]GroupSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetGroupSubjectByIdTeacherSub(cmd.TeacherSubjectId)
}

func (s *subjectService) GetGroupSubjectByGroupId(cmd *GetGroupSubjectByGroupId) ([]GroupSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Student.ToString(), common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetGroupSubjectByGroupId(cmd.GroupId)
}

func (s *subjectService) GetGroupSubjectByTeacherGroupIds(cmd *GetGroupSubjectByTeacherGroupIdsCommand) (*GroupSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetGroupSubjectByTeacherGroupIds(cmd.TeacherSubjectId, cmd.GroupId)
}

func (s *subjectService) GetTeacherSubjectByToken(cmd *GetTeacherSubjectsByTokenCommand) ([]TeacherSubject, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return s.subjectStore.GetTeacherSubjectsByTeacherId(cmd.CurrentUserId)
}
