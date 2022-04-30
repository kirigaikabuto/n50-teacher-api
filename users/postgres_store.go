package users

import (
	"database/sql"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

var marketplaceAppRepoQueries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		id TEXT,
		username TEXT,
		password TEXT,
		email TEXT,
		first_name TEXT,
		last_name TEXT,
		type TEXT,
		created_date date,
		PRIMARY KEY(id)
	);`,
}

type usersStore struct {
	db *sql.DB
}

func NewPostgresUsersStore(cfg common.PostgresConfig) (UsersStore, error) {
	db, err := common.GetDbConn(common.GetConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range marketplaceAppRepoQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &usersStore{db: db}
	return store, nil
}

func (u *usersStore) Update(user *UserUpdate) (*User, error) {
	q := "update users set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if user.Password != nil {
		cnt++
		hashPassword, err := setdata_common.HashPassword(*user.Password)
		if err != nil {
			return nil, err
		}
		*user.Password = hashPassword
		parts = append(parts, "password = $"+strconv.Itoa(cnt))
		values = append(values, user.Password)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, user.Id)
	result, err := u.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrUserNotFound
	}
	return u.Get(user.Id)
}

func (u *usersStore) Create(user *User) (*User, error) {
	hashPassword, err := setdata_common.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashPassword
	result, err := u.db.Exec("INSERT INTO users (id, username, password, email, first_name, last_name, type, created_date) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, current_date)",
		user.Id, user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.Type.ToString(),
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateUserUnknown
	}
	return user, nil
}

func (u *usersStore) Get(id string) (*User, error) {
	user := &User{}
	err := u.db.QueryRow("select id, username, password, email, first_name, last_name, type, created_date "+
		"from users where id = $1 limit 1", id).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Type, &user.CreatedDate)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersStore) List() ([]User, error) {
	users := []User{}
	var values []interface{}
	q := "select id, username, password, email, first_name, last_name, type, created_date from users"
	//cnt := 1
	rows, err := u.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Type, &user.CreatedDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *usersStore) GetByUsernameAndPassword(username, password string) (*User, error) {
	user := &User{}
	err := u.db.QueryRow("select id, username, password, email, first_name, last_name, type, created_date "+
		"from users where username = $1 limit 1", &username).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Type, &user.CreatedDate)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	compare := setdata_common.CheckPasswordHash(password, user.Password)
	if !compare {
		return nil, ErrUserPasswordNotCorrect
	}
	return user, nil
}

func (u *usersStore) Delete(id string) error {
	result, err := u.db.Exec("delete from users where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrUserNotFound
	}
	return nil
}
