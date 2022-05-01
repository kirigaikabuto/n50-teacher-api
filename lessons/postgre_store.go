package lessons

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"log"
	"strconv"
	"strings"
)

var lessonsQueries = []string{
	`create table if not exists lessons(
		id text,
		name text,
		description text,
		video_file_url text,
		document_file_url text,
		group_subject_id text,
		created_date date,
		primary key(id)
	);`,
}

type lessonPostgreStore struct {
	db *sql.DB
}

func NewLessonsPostgreStore(cfg common.PostgresConfig) (LessonStore, error) {
	db, err := common.GetDbConn(common.GetConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range lessonsQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &lessonPostgreStore{db: db}
	return store, nil
}

func (l *lessonPostgreStore) CreateLesson(model *Lesson) (*Lesson, error) {
	model.Id = uuid.New().String()
	result, err := l.db.Exec("INSERT INTO lessons (id, name, description, video_file_url, document_file_url, group_subject_id, created_date) "+
		"VALUES ($1, $2, $3, $4, $5, $6, current_date)",
		model.Id, model.Name, model.Description, model.VideoFileUrl, model.DocumentFileUrl, model.GroupSubjectId,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateLessonUnknown
	}
	return model, nil
}

func (l *lessonPostgreStore) ListLessonByGroupSubjectId(id string) ([]Lesson, error) {
	var objects []Lesson
	var values []interface{}
	values = append(values, id)
	q := "select " +
		"id, name, description, video_file_url, document_file_url, group_subject_id, created_date " +
		"from lessons where group_subject_id = $1"
	rows, err := l.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := Lesson{}
		err = rows.Scan(
			&obj.Id, &obj.Name, &obj.Description,
			&obj.VideoFileUrl, &obj.DocumentFileUrl, &obj.GroupSubjectId,
			&obj.CreatedDate,
		)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (l *lessonPostgreStore) GetLessonById(id string) (*Lesson, error) {
	obj := &Lesson{}
	err := l.db.QueryRow("select id, name, description, video_file_url, document_file_url, group_subject_id, created_date from lessons where id = $1", id).
		Scan(
			&obj.Id, &obj.Name, &obj.Description,
			&obj.VideoFileUrl, &obj.DocumentFileUrl, &obj.GroupSubjectId,
			&obj.CreatedDate,
		)
	if err == sql.ErrNoRows {
		return nil, ErrLessonNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (l *lessonPostgreStore) UpdateLesson(model *LessonUpdate) (*Lesson, error) {
	q := "update lessons set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if model.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, model.Name)
	}
	if model.Description != nil {
		cnt++
		parts = append(parts, "description = $"+strconv.Itoa(cnt))
		values = append(values, model.Description)
	}
	if model.VideoFileUrl != nil {
		cnt++
		parts = append(parts, "video_file_url = $"+strconv.Itoa(cnt))
		values = append(values, model.VideoFileUrl)
	}
	if model.DocumentFileUrl != nil {
		cnt++
		parts = append(parts, "document_file_url = $"+strconv.Itoa(cnt))
		values = append(values, model.DocumentFileUrl)
	}
	if model.GroupSubjectId != nil {
		cnt++
		parts = append(parts, "group_subject_id = $"+strconv.Itoa(cnt))
		values = append(values, model.GroupSubjectId)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, model.Id)
	result, err := l.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrLessonNotFound
	}
	return l.GetLessonById(model.Id)
}

func (l *lessonPostgreStore) DeleteLesson(id string) error {
	result, err := l.db.Exec("delete from lessons where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrLessonNotFound
	}
	return nil
}
