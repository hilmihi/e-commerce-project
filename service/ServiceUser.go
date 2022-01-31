package service

import (
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/repository"
)

type ServiceUser interface {
	ServiceUserLogin(input helper.RequestUserLogin) (entities.User, error)
	ServiceUsersGet() ([]entities.User, error)
	ServiceUserGet(id int) (entities.User, error)
	ServiceUserCreate(input entities.User) (entities.User, error)
	ServiceUserUpdate(id int, input entities.User) (entities.User, error)
	ServiceUserDelete(id int) (entities.User, error)
}

type serviceUser struct {
	repository1 repository.RepositoryUser
}

func NewUserService(repository1 repository.RepositoryUser) *serviceUser {
	return &serviceUser{repository1}
}

func (su *serviceUser) ServiceUserLogin(input helper.RequestUserLogin) (entities.User, error) {
	email := input.Email
	password := input.Password

	var user entities.User
	user, err := su.repository1.FindByEmail(email)
	if err != nil {
		return user, err
	}

	match, err := helper.CheckPasswordHash(password, user.Password)
	if err != nil {
		return user, err
	}

	if !match {
		return user, fmt.Errorf("Email atau Password Anda Salah!")
	}

	return user, nil
}

func (su *serviceUser) ServiceUsersGet() ([]entities.User, error) {
	users, err := su.repository1.GetUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *serviceUser) ServiceUserGet(id int) (entities.User, error) {
	user, err := s.repository1.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *serviceUser) ServiceUserCreate(input entities.User) (entities.User, error) {
	var err error
	input.Password, err = helper.HashPassword(input.Password)

	if err != nil {
		fmt.Println(err)
		return input, err
	}

	createUser, err := s.repository1.CreateUser(input)
	if err != nil {
		fmt.Println(err)
		return createUser, err
	}

	return createUser, nil
}

func (s *serviceUser) ServiceUserUpdate(id int, input entities.User) (entities.User, error) {
	user, err := s.repository1.GetUser(id)
	if err != nil {
		return user, err
	}

	input.Id = id

	updateUser, err := s.repository1.UpdateUser(input)
	if err != nil {
		return updateUser, err
	}
	return updateUser, nil
}

func (s *serviceUser) ServiceUserDelete(id int) (entities.User, error) {
	userID, err := s.ServiceUserGet(id)
	if err != nil {
		return userID, err
	}

	deleteUser, err := s.repository1.DeleteUser(userID)
	if err != nil {
		return deleteUser, err
	} else {
		return deleteUser, nil
	}
}
