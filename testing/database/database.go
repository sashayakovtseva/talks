package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	// driver
	_ "github.com/go-sql-driver/mysql"
)

type (
	users struct{}

	// User represents user entity
	User struct {
		ID    int64
		Name  string
		Email string
	}
)

var (
	db *sql.DB

	// Users offer access to users table
	Users users
)

func init() {
	var err error
	user, psw := os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PSW")
	instance, dbName := os.Getenv("MYSQL_INSTANCE"), os.Getenv("MYSQL_DB_NAME")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, psw, instance, dbName))
	if err != nil {
		log.Fatal("could not connect to db", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("unreachable db: ", err)
	}
}

// Fetch fetches user from db
func (users) Fetch(ctx context.Context, id int64) (User, error) {
	row := db.QueryRowContext(ctx, `
SELECT id, name, email
  FROM users
 WHERE id = ?
`, id)
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

// IsManager fetches is_manager flag for user
func (users) IsManager(ctx context.Context, id int64) (bool, error) {
	row := db.QueryRowContext(ctx, `
SELECT is_manager
  FROM users
 WHERE id = ?
`, id)
	var isManager bool
	err := row.Scan(&isManager)
	return isManager, err
}
