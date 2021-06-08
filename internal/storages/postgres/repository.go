package postgres

import (
	"skeleton/internal/storages"
	"skeleton/libs"

	"github.com/jinzhu/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func InitRepository(conn *gorm.DB) storages.EntityRepository {
	return &repository{
		Conn: conn,
	}
}

func (repo *repository) Login(userID string, pwd string) (err error) {
	user := storages.User{}
	err = repo.Conn.Where(`id = ?`, userID).First(&user).Error
	if err != nil {
		return err
	}
	err = libs.ComparePassword(user.Password, pwd)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetUserByID(userID string) (user storages.User, err error) {
	err = repo.Conn.Where(`id = ?`, userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repository) InsertUser(user *storages.User) (err error) {
	err = repo.Conn.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
