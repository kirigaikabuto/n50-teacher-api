package subjects

type CreateSubjectCommand struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *CreateSubjectCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).CreateSubject(cmd)
}

type ListSubjectsCommand struct {
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *ListSubjectsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).ListSubjects(cmd)
}

type GetSubjectByIdCommand struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetSubjectByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetSubjectById(cmd)
}

type GetSubjectsByGroupId struct {
	GroupId         string `json:"group_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetSubjectsByGroupId) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetSubjectsByGroupId(cmd)
}

type CreateTeacherSubjectCommand struct {
	TeacherId       string `json:"teacher_id"`
	SubjectId       string `json:"subject_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *CreateTeacherSubjectCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).CreateTeacherSubject(cmd)
}

type ListTeacherSubjectsCommand struct {
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *ListTeacherSubjectsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).ListTeacherSubjects(cmd)
}

type GetTeacherSubjectByIdCommand struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetTeacherSubjectByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetTeacherSubjectById(cmd)
}

type GetTeacherSubjectsByTeacherIdCommand struct {
	TeacherId       string `json:"teacher_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetTeacherSubjectsByTeacherIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetTeacherSubjectsByTeacherId(cmd)
}

type GetTeacherSubjectsBySubjectIdCommand struct {
	SubjectId       string `json:"subject_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetTeacherSubjectsBySubjectIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetTeacherSubjectsBySubjectId(cmd)
}

type CreateGroupSubjectCommand struct {
	GroupId          string `json:"group_id"`
	TeacherSubjectId string `json:"teacher_subject_id"`
	CurrentUserType  string `json:"-"`
	CurrentUserId    string `json:"-"`
}

func (cmd *CreateGroupSubjectCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).CreateGroupSubject(cmd)
}

type ListGroupSubjectsCommand struct {
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *ListGroupSubjectsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).ListGroupSubjects(cmd)
}

type GetGroupSubjectById struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetGroupSubjectById) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetGroupSubjectsById(cmd)
}

type GetGroupSubjectByIdTeacherSub struct {
	TeacherSubjectId string `json:"teacher_subject_id"`
	CurrentUserType  string `json:"-"`
	CurrentUserId    string `json:"-"`
}

func (cmd *GetGroupSubjectByIdTeacherSub) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetGroupSubjectByIdTeacherSub(cmd)
}

type GetGroupSubjectByGroupId struct {
	GroupId         string `json:"group_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetGroupSubjectByGroupId) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetGroupSubjectByGroupId(cmd)
}

type GetGroupSubjectByTeacherGroupIdsCommand struct {
	GroupId          string `json:"group_id"`
	TeacherSubjectId string `json:"teacher_subject_id"`
	CurrentUserType  string `json:"-"`
	CurrentUserId    string `json:"-"`
}

func (cmd *GetGroupSubjectByTeacherGroupIdsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetGroupSubjectByTeacherGroupIds(cmd)
}

type GetTeacherSubjectsByTokenCommand struct {
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetTeacherSubjectsByTokenCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(SubjectService).GetTeacherSubjectByToken(cmd)
}
