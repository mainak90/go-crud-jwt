package userRepository

import (
	"database/sql"
	"fmt"
	"go-crud-jwt/models"
	"log"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	fmt.Println(err)
	logFatal(err)

	user.Password = ""
	return user
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	fmt.Println(err)
	if err != nil {
		return user, err
	}

	return user, nil
}
