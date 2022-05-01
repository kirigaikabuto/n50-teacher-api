package lessons

import (
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"github.com/kirigaikabuto/n50-teacher-api/subjects"
)

type LessonService interface {
	CreateLesson(cmd *CreateLessonCommand) (*Lesson, error)
	ListLessonByGroupSubjectId(cmd *ListLessonByGroupSubjectIdCommand) ([]Lesson, error)
	GetLessonById(cmd *GetLessonByIdCommand) (*Lesson, error)
	UpdateLesson(cmd *UpdateLessonCommand) (*Lesson, error)
	DeleteLesson(cmd *DeleteLessonCommand) error
}

type lessonService struct {
	lessonStore  LessonStore
	subjectStore subjects.SubjectStore
}

func NewLessonService(l LessonStore, s subjects.SubjectStore) LessonService {
	return &lessonService{lessonStore: l, subjectStore: s}
}

func (l *lessonService) CreateLesson(cmd *CreateLessonCommand) (*Lesson, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	_, err := l.subjectStore.GetGroupSubjectsById(cmd.GroupSubjectId)
	if err != nil {
		return nil, err
	}
	return l.lessonStore.CreateLesson(&Lesson{
		Name:            cmd.Name,
		Description:     cmd.Description,
		VideoFileUrl:    cmd.VideoFileUrl,
		DocumentFileUrl: cmd.DocumentFileUrl,
		GroupSubjectId:  cmd.GroupSubjectId,
	})
}

func (l *lessonService) ListLessonByGroupSubjectId(cmd *ListLessonByGroupSubjectIdCommand) ([]Lesson, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return l.lessonStore.ListLessonByGroupSubjectId(cmd.GroupSubjectId)
}

func (l *lessonService) GetLessonById(cmd *GetLessonByIdCommand) (*Lesson, error) {
	if !common.IsAvailableResource(cmd.CurrentUserType, []string{common.Teacher.ToString(), common.Admin.ToString()}) {
		return nil, ErrNoAccessPermissions
	}
	return l.lessonStore.GetLessonById(cmd.Id)
}

func (l *lessonService) UpdateLesson(cmd *UpdateLessonCommand) (*Lesson, error) {
	updateLesson := &LessonUpdate{}
	oldLesson, err := l.lessonStore.GetLessonById(cmd.Id)
	if err != nil {
		return nil, err
	}
	if cmd.Name != "" && cmd.Name != oldLesson.Name {
		updateLesson.Name = &cmd.Name
	}
	if cmd.Description != "" && cmd.Description != oldLesson.Description {
		updateLesson.Description = &cmd.Description
	}
	if cmd.VideoFileUrl != "" && cmd.VideoFileUrl != oldLesson.VideoFileUrl {
		updateLesson.VideoFileUrl = &cmd.VideoFileUrl
	}
	if cmd.DocumentFileUrl != "" && cmd.DocumentFileUrl != oldLesson.DocumentFileUrl {
		updateLesson.DocumentFileUrl = &cmd.DocumentFileUrl
	}
	if cmd.GroupSubjectId != "" && cmd.GroupSubjectId != oldLesson.GroupSubjectId {
		updateLesson.GroupSubjectId = &cmd.GroupSubjectId
	}
	return l.lessonStore.UpdateLesson(updateLesson)
}

func (l *lessonService) DeleteLesson(cmd *DeleteLessonCommand) error {
	return l.lessonStore.DeleteLesson(cmd.Id)
}
