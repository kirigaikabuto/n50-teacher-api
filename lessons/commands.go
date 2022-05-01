package lessons

type CreateLessonCommand struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	VideoFileUrl    string `json:"video_file_url"`
	DocumentFileUrl string `json:"document_file_url"`
	GroupSubjectId  string `json:"group_subject_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *CreateLessonCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(LessonService).CreateLesson(cmd)
}

type ListLessonByGroupSubjectIdCommand struct {
	GroupSubjectId  string `json:"group_subject_id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *ListLessonByGroupSubjectIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(LessonService).ListLessonByGroupSubjectId(cmd)
}

type GetLessonByIdCommand struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *GetLessonByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(LessonService).GetLessonById(cmd)
}

type UpdateLessonCommand struct {
	Id              string  `json:"-"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	VideoFileUrl    *string `json:"video_file_url"`
	DocumentFileUrl *string `json:"document_file_url"`
	GroupSubjectId  *string `json:"group_subject_id"`
	CurrentUserType string  `json:"-"`
	CurrentUserId   string  `json:"-"`
}

func (cmd *UpdateLessonCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(LessonService).UpdateLesson(cmd)
}

type DeleteLessonCommand struct {
	Id              string `json:"id"`
	CurrentUserType string `json:"-"`
	CurrentUserId   string `json:"-"`
}

func (cmd *DeleteLessonCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(LessonService).DeleteLesson(cmd)
}
