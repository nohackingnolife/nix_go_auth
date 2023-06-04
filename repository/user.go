package repository

import (
	"authGo/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	p1, _ = bcrypt.GenerateFromPassword([]byte("111111111111"), bcrypt.DefaultCost)
	p2, _ = bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)

	users = []*model.User{
		{
			ID:       1,
			Name:     "user1",
			Email:    "first@gmail.com",
			Password: string(p1),
		},
		{
			ID:       2,
			Name:     "user2",
			Email:    "second@email.com",
			Password: string(p2),
		},
	}
)

type UserRepository struct {
	users *[]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: &users}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	for _, user := range *r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) GetUserById(id int) (*model.User, error) {
	for _, user := range *r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) RegisterUser(name string, email string, password string) (*model.User, error) {
	_, err := r.GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("user already exist")
	}

	id := (*r.users)[len(*r.users)-1].ID + 1
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	u := &model.User{ID: id, Name: name, Email: email, Password: string(p)}
	*r.users = append(*r.users, u)

	return u, nil
}
