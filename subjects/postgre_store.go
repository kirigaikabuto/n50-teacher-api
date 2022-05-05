package subjects

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"log"
)

var subjectQueries = []string{
	`create table if not exists subjects(
		id text,
		name text,
		description text,
		created_date date,
		primary key(id)
	);`,
	`create table if not exists teacher_subjects(
		id text,
		teacher_id text,
		subject_id text,
		created_date date,
		primary key(id)
	);`,
	`create table if not exists group_subjects(
		id text,
		group_id text,
		teacher_subject_id text,
		created_date date,
		primary key(id)
	);`,
}

type subjectStore struct {
	db *sql.DB
}

func NewSubjectPostgreStore(cfg common.PostgresConfig) (SubjectStore, error) {
	db, err := common.GetDbConn(common.GetConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range subjectQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &subjectStore{db: db}
	return store, nil
}

//subjects
func (s *subjectStore) CreateSubject(model *Subject) (*Subject, error) {
	model.Id = uuid.New().String()
	result, err := s.db.Exec("INSERT INTO subjects (id, name, description, created_date) "+
		"VALUES ($1, $2, $3, current_date)",
		model.Id, model.Name, model.Description,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateSubjectUnknown
	}
	return model, nil
}

func (s *subjectStore) ListSubjects() ([]Subject, error) {
	var objects []Subject
	var values []interface{}
	q := "select " +
		"id, name, description, created_date " +
		"from subjects"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := Subject{}
		err = rows.Scan(
			&obj.Id, &obj.Name, &obj.Description, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *subjectStore) GetSubjectById(id string) (*Subject, error) {
	obj := &Subject{}
	err := s.db.QueryRow("select id, name, description, created_date from subjects where id = $1", id).
		Scan(&obj.Id, &obj.Name, &obj.Description, &obj.CreatedDate,
		)
	if err == sql.ErrNoRows {
		return nil, ErrSubjectNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

//teacher subjects
func (s *subjectStore) CreateTeacherSubject(model *TeacherSubject) (*TeacherSubject, error) {
	model.Id = uuid.New().String()
	result, err := s.db.Exec("INSERT INTO teacher_subjects (id, teacher_id, subject_id, created_date) "+
		"VALUES ($1, $2, $3, current_date)",
		model.Id, model.TeacherId, model.SubjectId,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateTeacherSubjectUnknown
	}
	return model, nil
}

func (s *subjectStore) ListTeacherSubjects() ([]TeacherSubject, error) {
	var objects []TeacherSubject
	var values []interface{}
	q := "select " +
		"id, teacher_id, subject_id, created_date " +
		"from teacher_subjects"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := TeacherSubject{}
		err = rows.Scan(
			&obj.Id, &obj.TeacherId, &obj.SubjectId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *subjectStore) GetTeacherSubjectById(id string) (*TeacherSubject, error) {
	obj := &TeacherSubject{}
	err := s.db.QueryRow("select id, teacher_id, subject_id, created_date from teacher_subjects where id = $1", id).
		Scan(&obj.Id, &obj.TeacherId, &obj.SubjectId, &obj.CreatedDate,
		)
	if err == sql.ErrNoRows {
		return nil, ErrTeacherSubjectNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *subjectStore) GetTeacherSubjectsByTeacherId(id string) ([]TeacherSubject, error) {
	var objects []TeacherSubject
	var values []interface{}
	values = append(values, id)
	q := "select " +
		"id, teacher_id, subject_id, created_date " +
		"from teacher_subjects  where teacher_id = $1"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := TeacherSubject{}
		err = rows.Scan(
			&obj.Id, &obj.TeacherId, &obj.SubjectId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *subjectStore) GetTeacherSubjectsBySubjectId(id string) ([]TeacherSubject, error) {
	var objects []TeacherSubject
	var values []interface{}
	values = append(values, id)
	q := "select " +
		"id, teacher_id, subject_id, created_date " +
		"from teacher_subjects  where subject_id = $1"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := TeacherSubject{}
		err = rows.Scan(
			&obj.Id, &obj.TeacherId, &obj.SubjectId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

//group subjects
func (s *subjectStore) CreateGroupSubject(model *GroupSubject) (*GroupSubject, error) {
	model.Id = uuid.New().String()
	model.Id = uuid.New().String()
	result, err := s.db.Exec("INSERT INTO group_subjects (id, group_id, teacher_subject_id, created_date) "+
		"VALUES ($1, $2, $3, current_date)",
		model.Id, model.GroupId, model.TeacherSubjectId,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateGroupSubjectUnknown
	}
	return model, nil
}

func (s *subjectStore) ListGroupSubjects() ([]GroupSubject, error) {
	var objects []GroupSubject
	var values []interface{}
	q := "select " +
		"id, group_id, teacher_subject_id, created_date " +
		"from group_subjects"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := GroupSubject{}
		err = rows.Scan(
			&obj.Id, &obj.GroupId, &obj.TeacherSubjectId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *subjectStore) GetGroupSubjectsById(id string) (*GroupSubject, error) {
	obj := &GroupSubject{}
	err := s.db.QueryRow("select id, group_id, teacher_subject_id, created_date from group_subjects where id = $1", id).
		Scan(&obj.Id, &obj.GroupId, &obj.TeacherSubjectId, &obj.CreatedDate,
		)
	if err == sql.ErrNoRows {
		return nil, ErrGroupSubjectNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *subjectStore) GetGroupSubjectByIdTeacherSub(id string) ([]GroupSubject, error) {
	var objects []GroupSubject
	var values []interface{}
	values = append(values, id)
	q := "select " +
		"id, group_id, teacher_subject_id, created_date " +
		"from group_subjects where teacher_subject_id = $1"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := GroupSubject{}
		err = rows.Scan(
			&obj.Id, &obj.GroupId, &obj.TeacherSubjectId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *subjectStore) GetGroupSubjectByGroupId(id string) ([]GroupSubject, error) {
	var objects []GroupSubject
	var values []interface{}
	values = append(values, id)
	q := "select " +
		"id, group_id, teacher_subject_id, created_date " +
		"from group_subjects where group_id = $1"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := GroupSubject{}
		err = rows.Scan(
			&obj.Id, &obj.GroupId, &obj.TeacherSubjectId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *subjectStore) GetGroupSubjectByTeacherGroupIds(idTeacher string, idGroup string) (*GroupSubject, error) {
	obj := &GroupSubject{}
	err := s.db.QueryRow("select id, group_id, teacher_subject_id, created_date from group_subjects where teacher_subject_id = $1 and group_id = $2", idTeacher, idGroup).
		Scan(&obj.Id, &obj.GroupId, &obj.TeacherSubjectId, &obj.CreatedDate,
		)
	if err == sql.ErrNoRows {
		return nil, ErrGroupSubjectNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}
