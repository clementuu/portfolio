package cv

import "back/model"

func GetFormations() []model.Formation {
	return storage.GetFormations()
}

func GetExperiences() []model.Experience {
	return storage.GetExperiences()
}
