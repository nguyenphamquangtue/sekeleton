package usecase

import (
	"errors"
	"skeleton/internal/storages"
	"skeleton/libs"
	"skeleton/validate"
)

type useCase struct {
	repository storages.EntityRepository
}

func InitUseCase(repo storages.EntityRepository) storages.EntityUseCase {
	return &useCase{
		repository: repo,
	}
}

func (u *useCase) GetAuthToken(args map[string]string) (map[string]interface{}, error) {

	var (
		res = map[string]interface{}{}
		err error
	)

	err = u.repository.Login(args["username"], args["password"])

	if err != nil {
		return res, errors.New("incorrect user_name/pwd")
	}

	res["token"], err = validate.CreateToken(args["username"])

	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *useCase) RegisterUser(user storages.User) (map[string]interface{}, error) {
	var (
		res          = map[string]interface{}{}
		errorGetUser = make(chan error)
	)
	go func() {
		_, err := u.repository.GetUserByID(user.ID)
		errorGetUser <- err
	}()

	err := <-errorGetUser
	if err == nil {
		res["status"] = 400
		return res, err
	}

	defer close(errorGetUser)
	password, err := libs.HashPassword(user.Password)
	if err != nil {
		res["status"] = 400
		return res, err
	}
	user.Password = *password
	err = u.repository.InsertUser(&user)
	if err != nil {
		return res, err
	}

	res["token"], err = validate.CreateToken(user.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}
