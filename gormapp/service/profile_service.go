package service

import (
	"gormapp/dao"
	"gormapp/model"
)

func GetProfiles() ([]model.Profile, error) {
	return dao.GetAllProfiles()
}

func GetProfile(id int) (model.Profile, error) {
	return dao.GetProfileByID(id)
}

func CreateProfile(profile model.Profile) error {
	return dao.CreateProfile(profile)
}

func UpdateProfile(profile model.Profile) error {
	return dao.UpdateProfile(profile)
}

func DeleteProfile(id int) error {
	return dao.DeleteProfile(id)
}
