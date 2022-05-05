package subjects

type SubjectStore interface {
	//subjects
	CreateSubject(model *Subject) (*Subject, error)
	ListSubjects() ([]Subject, error)
	GetSubjectById(id string) (*Subject, error)
	//teacher subjects
	CreateTeacherSubject(model *TeacherSubject) (*TeacherSubject, error)
	ListTeacherSubjects() ([]TeacherSubject, error)
	GetTeacherSubjectById(id string) (*TeacherSubject, error)
	GetTeacherSubjectsByTeacherId(id string) ([]TeacherSubject, error)
	GetTeacherSubjectsBySubjectId(id string) ([]TeacherSubject, error)
	//group subjects
	CreateGroupSubject(model *GroupSubject) (*GroupSubject, error)
	ListGroupSubjects() ([]GroupSubject, error)
	GetGroupSubjectsById(id string) (*GroupSubject, error)
	GetGroupSubjectByIdTeacherSub(id string) ([]GroupSubject, error)
	GetGroupSubjectByGroupId(id string) ([]GroupSubject, error)
	GetGroupSubjectByTeacherGroupIds(idTeacher string, idGroup string) (*GroupSubject, error)
}
