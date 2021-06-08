package test

import (
	"log"
	"skeleton/driver"
	"skeleton/internal/storages/postgres"
	"testing"
)

func Test_Storages_Login(t *testing.T) {

	var (
		uname = "notail"
		pwd   = "1234567"
	)

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	repository := postgres.InitRepository(db.Conn)

	err = repository.Login(uname, pwd)

	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_Storages_GetUserByID(t *testing.T) {

	var (
		uname = "notail"
	)

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	repository := postgres.InitRepository(db.Conn)

	_, err = repository.GetUserByID(uname)

	if err != nil {
		t.Errorf(err.Error())
	}
}
