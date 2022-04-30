package groups

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"log"
)

var userGroupQueries = []string{
	`create table if not exists groups(
		id text,
		name text,
		created_date date,
		primary key(id)
	);`,
	`create table if not exists user_groups(
		id text,
		user_id text,
		group_id text,
		created_date date,
		primary key(id)
	);`,
}

type userGroupStore struct {
	db *sql.DB
}

func NewUserGroupPostgreStore(cfg common.PostgresConfig) (UserGroupStore, error) {
	db, err := common.GetDbConn(common.GetConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range userGroupQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &userGroupStore{db: db}
	return store, nil
}

func (g *userGroupStore) CreateGroup(model *Group) (*Group, error) {
	model.Id = uuid.New().String()
	result, err := g.db.Exec("INSERT INTO groups (id, name, created_date) "+
		"VALUES ($1, $2, current_date)",
		model.Id, model.Name,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateGroupUnknown
	}
	return model, nil
}

func (g *userGroupStore) ListGroup() ([]Group, error) {
	var objects []Group
	var values []interface{}
	q := "select " +
		"id, name, created_date " +
		"from groups"
	rows, err := g.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := Group{}
		err = rows.Scan(
			&obj.Id, &obj.Name, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (g *userGroupStore) GetGroupById(id string) (*Group, error) {
	obj := &Group{}
	err := g.db.QueryRow("select id, name, created_date from groups where id = $1", id).
		Scan(&obj.Id, &obj.Name, &obj.CreatedDate,
		)
	if err == sql.ErrNoRows {
		return nil, ErrGroupNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (g *userGroupStore) CreateUserGroup(model *UserGroup) (*UserGroup, error) {
	model.Id = uuid.New().String()
	result, err := g.db.Exec("INSERT INTO user_groups (id, user_id, group_id, created_date) "+
		"VALUES ($1, $2, $3, current_date)",
		model.Id, model.UserId, model.GroupId,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateUserGroupUnknown
	}
	return model, nil
}

func (g *userGroupStore) GetUserGroupByGroupId(groupId string) ([]UserGroup, error) {
	var objects []UserGroup
	var values []interface{}
	values = append(values, groupId)
	q := "select " +
		"id, user_id, group_id, created_date " +
		"from user_groups where group_id = $1"
	rows, err := g.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := UserGroup{}
		err = rows.Scan(
			&obj.Id, &obj.UserId, &obj.GroupId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (g *userGroupStore) GetUserGroupByUserId(userId string) ([]UserGroup, error) {
	var objects []UserGroup
	var values []interface{}
	values = append(values, userId)
	q := "select " +
		"id, user_id, group_id, created_date " +
		"from user_groups where user_id = $1"
	rows, err := g.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := UserGroup{}
		err = rows.Scan(
			&obj.Id, &obj.UserId, &obj.GroupId, &obj.CreatedDate)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (g *userGroupStore) RemoveUserGroupById(id string) error {
	_, err := g.db.Exec("DELETE FROM user_groups WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
