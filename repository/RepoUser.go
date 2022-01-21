package repository

import (
	"database/sql"
	"log"
	"sirclo/api/entities"
	"sirclo/api/helper"
)

type RepositoryUser interface {
	FindByEmail(email string) (entities.User, error)
	GetUsers() ([]entities.User, error)
	CreateUser(user helper.RequestUserCreate) (helper.RequestUserCreate, error)
	GetUser(id int) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	DeleteUser(user entities.User) (entities.User, error)
}

type Repository_User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *Repository_User {
	return &Repository_User{db}
}

func (r *Repository_User) FindByEmail(email string) (entities.User, error) {
	row := r.db.QueryRow("select id, email, password from users where email=?", email)
	var user entities.User

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil

}

//get users
func (r *Repository_User) GetUsers() ([]entities.User, error) {
	var users []entities.User
	results, err := r.db.Query("select id, name, email from users order by id asc")
	if err != nil {
		log.Fatalf("Error")
	}

	defer results.Close()

	for results.Next() {
		var user entities.User

		err = results.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Fatalf("Error")
		}

		users = append(users, user)
	}
	return users, nil
}

//get user
func (r *Repository_User) GetUser(id int) (entities.User, error) {
	var user entities.User

	row := r.db.QueryRow(`SELECT id, name, email FROM users WHERE id = ?`, id)

	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

//create user
func (r *Repository_User) CreateUser(user helper.RequestUserCreate) (helper.RequestUserCreate, error) {
	query := `INSERT INTO users (name, email, password, birth_date, phone_number, photo, gender_char, address, created_date, updated_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, now(), now())`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.Birth_date, user.Phone_number, user.Photo, user.Gender, user.Address)
	if err != nil {
		return user, err
	}

	return user, nil
}

//update user
func (r *Repository_User) UpdateUser(user entities.User) (entities.User, error) {
	query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

//delete user
func (r *Repository_User) DeleteUser(user entities.User) (entities.User, error) {
	query := `DELETE FROM users WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}
