package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connStr string) (*Db, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (d *Db) GetUsersByName(name string) []User {
	stmt, err := d.Prepare("SELECT * FORM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preperation Err", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Query Err:", err)
	}

	var r User
	users := []User{}

	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("Error scaning row:", err)
		}
		users = append(users, r)
	}

	return users
}
