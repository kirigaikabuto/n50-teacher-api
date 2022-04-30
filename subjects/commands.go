package subjects

type CreateSubjectCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListSubjectsCommand struct {
}

type GetSubjectByIdCommand struct {
	Id string `json:"id"`
}

type CreateTeacherSubject struct {
}
