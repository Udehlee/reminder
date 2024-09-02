package db

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/Udehlee/reminder/models/user"
)

type Store interface {
	SaveUser(user models.User) error
	UserEmail(email string) (models.User, error)
	UserPhoneNumber() ([]int, error)
}

type PgDB struct {
	Db  *PgConn
	Ctx context.Context
}

func NewPgDB(db *PgConn) *PgDB {
	ctx := context.Background()

	return &PgDB{
		Db:  db,
		Ctx: ctx,
	}

}

func (p *PgDB) SaveUser(user models.User) error {
	query := "INSERT INTO users (firstname,lastname,email,password) VALUES($1,$2,$3,$4) RETURNING userId,firstname,lastname,email"

	row := p.Db.Conn.QueryRow(p.Ctx, query, user.FirstName, user.LastName, user.Email)
	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		fmt.Println("error scanning row %w ", err)
	}
	return nil
}

func (p *PgDB) UserEmail(email string) (models.User, error) {
	user := models.User{}
	query := "SELECT user_id, firstname, lastname, email, pass_word FROM users WHERE email = $1"

	err := p.Db.Conn.QueryRow(p.Ctx, query, email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, fmt.Errorf("error querying user by email: %w", err)
	}

	return user, nil
}

func (p *PgDB) UserPhoneNumber() ([]int, error) {
	var phoneNumbers []int
	query := "SELECT phone_number FROM users"

	rows, err := p.Db.Conn.Query(p.Ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var phoneNumber int
		if err := rows.Scan(&phoneNumber); err != nil {
			return nil, err
		}
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return phoneNumbers, nil
}
