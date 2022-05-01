package lessons

type Lesson struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	VideoFileUrl    string `json:"video_file_url"`
	DocumentFileUrl string `json:"document_file_url"`
	GroupSubjectId  string `json:"group_subject_id"`
	CreatedDate     string `json:"created_date"`
}

type LessonUpdate struct {
	Id              string  `json:"id"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	VideoFileUrl    *string `json:"video_file_url"`
	DocumentFileUrl *string `json:"document_file_url"`
	GroupSubjectId  *string `json:"group_subject_id"`
}
