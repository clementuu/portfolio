package service

import "back/model"

func GetFormations() []model.Formation {
	return storage.GetFormations()
}
