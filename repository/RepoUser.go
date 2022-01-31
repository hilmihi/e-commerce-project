package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
)

type RepositoryUser interface {
	FindByEmail(email string) (entities.User, error)
	GetUsers() ([]entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
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
	row := r.db.QueryRow("select id, email, password from users where email=? and deleted_date is null", email)
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
	results, err := r.db.Query("select id, name, email, birth_date, phone_number, photo, gender, address from users where deleted_date is null order by id asc")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var user entities.User

		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Birth_date, &user.Phone_number, &user.Photo, &user.Gender, &user.Address)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

//get user
func (r *Repository_User) GetUser(id int) (entities.User, error) {
	var user entities.User

	row := r.db.QueryRow(`SELECT id, name, email, birth_date, phone_number, photo, gender, address FROM users WHERE id = ? AND deleted_date IS NULL`, id)
	fmt.Println(row)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Birth_date, &user.Phone_number, &user.Photo, &user.Gender, &user.Address)
	if err != nil {
		return user, err
	}

	return user, nil
}

//create user
func (r *Repository_User) CreateUser(user entities.User) (entities.User, error) {
	query := `INSERT INTO users (name, email, password, birth_date, phone_number, photo, gender, address, created_date, updated_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, now(), now())`

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
	query := `UPDATE users SET name = ?, email = ?, birth_date = ?, phone_number = ?, photo = ?, gender = ?, address = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Birth_date, user.Phone_number, user.Photo, user.Gender, user.Address, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

//delete user
func (r *Repository_User) DeleteUser(user entities.User) (entities.User, error) {
	query := `UPDATE users SET deleted_date = now() WHERE id = ?`

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
