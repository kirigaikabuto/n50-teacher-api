package subjects

type Subject struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedDate string `json:"created_date"`
}

type TeacherSubject struct {
	Id          string `json:"id"`
	TeacherId   string `json:"teacher_id"`
	SubjectId   string `json:"subject_id"`
	CreatedDate string `json:"created_date"`
}

type GroupSubject struct {
	Id               string `json:"id"`
	GroupId          string `json:"group_id"`
	TeacherSubjectId string `json:"teacher_subject_id"`
	CreatedDate      string `json:"created_date"`
}

type SubjectFullInfo struct {
	Subject        Subject `json:"subject"`
	TeacherSubject TeacherSubject `json:"teacher_subject"`
	GroupSubject   GroupSubject `json:"group_subject"`
}
