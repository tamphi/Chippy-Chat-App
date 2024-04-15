// package db

// import (
// 	"database/sql"

// 	_ "github.com/lib/pq" // <-- postgres
// )

// type Database struct {
// 	//todo: implement db struct --> done
// 	db *sql.DB
// }

// // todo : singleton pattern
// func NewDataBase() (*Database, error) {
// 	//todo: init new db --> done
// 	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/chippy-chat?sslmode=disabled")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Database{db: db}, nil
// }

// func (d *Database) Close() {
// 	d.db.Close()
// }

// func (d *Database) GetDB() *sql.DB {
// 	return d.db
// }

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	//db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/chippy-chat?sslmode=disable") //comment out if running on docker
	db, err := sql.Open("postgres", "postgresql://root:password@host.docker.internal:5432/chippy-chat?sslmode=disable") //comment out if running on local
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to database")
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
