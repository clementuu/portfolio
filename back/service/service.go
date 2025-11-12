package service

import "back/model"

func GetCompetences() []model.Competence {
	return storage.GetCompetences()
}
