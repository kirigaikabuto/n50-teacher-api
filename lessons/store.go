package lessons

type LessonStore interface {
	CreateLesson(model *Lesson) (*Lesson, error)
	ListLessonByGroupSubjectId(id string) ([]Lesson, error)
	GetLessonById(id string) (*Lesson, error)
	UpdateLesson(model *LessonUpdate) (*Lesson, error)
	DeleteLesson(id string) error
}
