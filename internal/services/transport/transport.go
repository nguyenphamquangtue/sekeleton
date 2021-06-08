package transport

import (
	"net/http"
	"skeleton/internal/services/usecase"
	"skeleton/internal/storages"
	"skeleton/internal/storages/postgres"
	"skeleton/libs"

	"github.com/jinzhu/gorm"
)

type transport struct {
	useCase storages.EntityUseCase
}

func Init(db *gorm.DB, method string, req *http.Request) (res map[string]interface{}, err error) {
	repository := postgres.InitRepository(db)
	useCase := usecase.InitUseCase(repository)
	transport := transport{
		useCase: useCase,
	}

	switch method {
	case "get-auth":
		res, err = transport.GetAuthToken(req)
	case "register":
		res, err = transport.RegisterUser(req)
	}

	return res, err
}

func (trans *transport) GetAuthToken(req *http.Request) (res map[string]interface{}, err error) {

	var (
		args = map[string]string{}
		user = libs.Value(req, "username")
		pwd  = libs.Value(req, "password")
	)

	args["username"] = user.String
	args["password"] = pwd.String

	res, err = trans.useCase.GetAuthToken(args)
	return res, err
}

func (trans *transport) RegisterUser(req *http.Request) (res map[string]interface{}, err error) {
	var (
		user     = storages.User{}
		username = libs.Value(req, "username")
		pwd      = libs.Value(req, "password")
	)
	user.ID = username.String
	user.Password = pwd.String
	res, err = trans.useCase.RegisterUser(user)
	return res, err
}
