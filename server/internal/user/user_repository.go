package user

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	//todo: define db transaction methods --> done
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	//todo: return Repository --> done
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var existUsername string = ""
	chechExistQuery := "SELECT username FROM users WHERE username = $1"
	r.db.QueryRowContext(ctx, chechExistQuery, user.Username).Scan(&existUsername)
	if len(existUsername) > 0 {
		return user, fmt.Errorf("username already exist")
	}

	var lastInsertId int
	query := "INSERT INTO users(username, password) VALUES ($1, $2) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&lastInsertId)

	if err != nil {
		return nil, fmt.Errorf("failed to create user in repository: %s", err)
	}
	user.ID = lastInsertId // <-- typecast if using 64bit for ID
	return user, nil
}

func (r *repository) GetUsername(ctx context.Context, username string) (*User, error) {
	//todo: query username from db from user ID, username, password --> done
	u := User{}
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := r.db.QueryRowContext(ctx, query, username).Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		return nil, fmt.Errorf("failed to get username in repository: %s", err)
	}

	return &u, nil
}

func (r *repository) ListUsers(ctx context.Context) ([]*User, error) {
	query := "SELECT username FROM users"

	users, err := r.db.QueryContext(ctx, query)
	var result []*User
	for users.Next() {
		var username string
		err := users.Scan(&username)
		if err != nil {
			fmt.Println("error")
		}
		user := User{
			Username: username,
		}
		result = append(result, &user)
	}

	if err != nil {
		fmt.Printf("Failed at ListUsers() in repo: %s", err)
		return nil, fmt.Errorf("failed to list all users from repository: %s", err)
	}

	return result, nil
}
