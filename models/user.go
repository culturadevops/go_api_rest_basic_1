package models

import (
	"errors"
	"log"

	"github.com/culturadevops/api/libs"
	"github.com/jinzhu/gorm"
)

var User *user

type user struct {
	gorm.Model
	Account  string
	Password string
}

func (this *user) List() []user {
	var data = []user{}
	err := libs.DB.Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (this *user) Info(id uint) (user, error) {
	var data user

	if libs.DB.Where("id  = ? ", id).Find(&data).RecordNotFound() {
		return user{}, errors.New("No existe registro con es id")
	}

	return data, nil
}

func (this *user) Add(account string, password string) (user, error) {
	var data user
	data.Account = account
	data.Password = password
	if err := libs.DB.Create(&data).Error; err != nil {
		return user{}, err
	} else {

		return data, nil
	}

}

func (this *user) Update(id uint, account string, password string) (user, error) {
	var data user

	if libs.DB.Where("id = ? ", id).First(&data).RecordNotFound() {
		return user{}, errors.New("No hay registro con ese id")
	}
	data.Account = account
	data.Password = password
	if err := libs.DB.Save(&data).Error; err != nil {
		return user{}, errors.New("no se pudo actualizar")
	}
	return data, nil

}

func (this *user) Del(id uint) error {
	var data user
	if err := libs.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
