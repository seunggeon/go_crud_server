package repository

import (
	"awesomeProject/types"
	"awesomeProject/types/errors"
)

type UserRepository struct {
	UserMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		UserMap: []*types.User{},
	}
}

func (u *UserRepository) Create(user *types.User) error {
	u.UserMap = append(u.UserMap, user)
	return nil
}

func (u *UserRepository) Update(name string, newAge int64) error {
	isExisted := false

	for _, user := range u.UserMap {
		if user.Name == name {
			user.Age = newAge
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(userName string) error {
	isExisted := false

	for index, user := range u.UserMap {
		if user.Name == userName {
			// slice UserMap
			u.UserMap = append(u.UserMap[:index], u.UserMap[index+1:]...)
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Get() []*types.User {
	return u.UserMap
}
